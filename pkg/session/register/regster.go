package register

import "github.com/yubo/golib/proc"

func init() {
	proc.RegisterHooks(hookOps)
}