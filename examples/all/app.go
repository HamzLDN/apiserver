package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/yubo/apiserver/examples/all/authn"
	"github.com/yubo/apiserver/examples/all/authz"
	"github.com/yubo/apiserver/examples/all/session"
	"github.com/yubo/apiserver/examples/all/tracing"
	"github.com/yubo/apiserver/examples/all/user"
	"github.com/yubo/apiserver/pkg/options"
	"github.com/yubo/golib/proc"
	"k8s.io/klog/v2"

	// authz's submodule, should be loaded before the authz module
	_ "github.com/yubo/apiserver/plugin/authz/abac/register"
	_ "github.com/yubo/apiserver/plugin/authz/alwaysallow/register"
	_ "github.com/yubo/apiserver/plugin/authz/alwaysdeny/register"
	_ "github.com/yubo/apiserver/plugin/authz/rbac/register"
	_ "github.com/yubo/apiserver/pkg/authorization/register"
	// TODO
	//_ "github.com/yubo/apiserver/plugin/authz/webhook/register"

	// authn
	_ "github.com/yubo/apiserver/pkg/authentication/register"
	_ "github.com/yubo/apiserver/plugin/authn/session/register"
	_ "github.com/yubo/apiserver/plugin/authn/token/bootstrap/register"
	_ "github.com/yubo/apiserver/plugin/authn/token/oidc/register"
	_ "github.com/yubo/apiserver/plugin/authn/token/tokenfile/register"
	// TODO
	//_ "github.com/yubo/apiserver/pkg/authentication/serviceaccount/register"
	//_ "github.com/yubo/apiserver/plugin/authn/webhook/register"

	_ "github.com/yubo/apiserver/pkg/apiserver/register"
	_ "github.com/yubo/apiserver/pkg/audit/register"
	_ "github.com/yubo/apiserver/pkg/db/register"
	_ "github.com/yubo/apiserver/pkg/debug/register"
	_ "github.com/yubo/apiserver/pkg/grpcserver/register"
	_ "github.com/yubo/apiserver/pkg/rest/swagger/register"
	_ "github.com/yubo/apiserver/pkg/session/register"
	_ "github.com/yubo/apiserver/pkg/tracing/register"
	_ "github.com/yubo/golib/logs/register"
	_ "github.com/yubo/golib/orm/sqlite"
)

const (
	AppName    = "example-all"
	moduleName = "example-all.main"
)

var (
	hookOps = []proc.HookOps{{
		Hook:     start,
		Owner:    moduleName,
		HookNum:  proc.ACTION_START,
		Priority: proc.PRI_MODULE,
	}, {
		Hook:     stop,
		Owner:    moduleName,
		HookNum:  proc.ACTION_STOP,
		Priority: proc.PRI_MODULE,
	}}
)

func newServerCmd() *cobra.Command {
	proc.RegisterHooks(hookOps)
	options.InstallReporter()

	ctx := context.Background()
	ctx = proc.WithName(ctx, os.Args[0])

	cmd := proc.NewRootCmd(ctx)
	cmd.AddCommand(options.NewVersionCmd())

	return cmd
}

func start(ctx context.Context) error {
	klog.Info("start")

	if err := session.New(ctx).Start(); err != nil {
		return err
	}
	if err := tracing.New(ctx).Start(); err != nil {
		return err
	}
	if err := user.New(ctx).Start(); err != nil {
		return err
	}
	if err := authn.New(ctx).Start(); err != nil {
		return err
	}
	if err := authz.New(ctx).Start(); err != nil {
		return err
	}

	return nil
}

func stop(ctx context.Context) error {
	klog.Info("stop")
	return nil
}
