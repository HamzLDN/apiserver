package abac

import (
	"github.com/yubo/apiserver/pkg/authorization"
	"github.com/yubo/apiserver/pkg/authorization/abac"
	"github.com/yubo/apiserver/pkg/authorization/authorizer"
	"github.com/yubo/apiserver/pkg/options"
	"github.com/yubo/golib/proc"
)

const (
	moduleName       = "authorization"
	submoduleName    = "ABAC"
	noUsernamePrefix = "-"
)

var (
	_auth   = &authModule{name: moduleName + "." + submoduleName}
	hookOps = []proc.HookOps{{
		Hook:        _auth.init,
		Owner:       moduleName,
		HookNum:     proc.ACTION_START,
		Priority:    proc.PRI_SYS_INIT,
		SubPriority: options.PRI_M_AUTHZ - 1,
	}}
	_config *config
)

type config struct {
	PolicyFile string `json:"policyFile" flag:"authorization-policy-file" description:"File with authorization policy in json line by line format, used with --authorization-mode=ABAC, on the secure port."`
}

func (o *config) Validate() error {
	return nil
}

type authModule struct {
	name   string
	config *config
}

func newConfig() *config {
	return &config{}
}

func (p *authModule) init(ops *proc.HookOps) error {
	c := ops.Configer()

	cf := newConfig()
	if err := c.ReadYaml(moduleName, cf); err != nil {
		return err
	}
	p.config = cf

	return nil
}

func init() {
	proc.RegisterHooks(hookOps)
	proc.RegisterFlags(moduleName, "authorization", newConfig())

	factory := func() (authorizer.Authorizer, error) {
		return abac.NewFromFile(_auth.config.PolicyFile)
	}

	authorization.RegisterAuthz(submoduleName, factory)
}