package ctags

import (
	"path"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/nfs"
)

type inner struct {
	ice.Code
	list string `name:"inner path=src/@key file=main.go@key line=1 auto ctags" help:"编辑器"`
}

func (s inner) Ctags(m *ice.Message, arg ...string) {
	if !nfs.ExistsFile(m, path.Join(arg[0], nfs.TAGS)) {
		s.System(m.Spawn(), arg[0], "ctags", "-R")
	}
}
func (s inner) Man(m *ice.Message, arg ...string) {
	m.Cmdy(cli.SYSTEM, "man", arg)
}
func (s inner) List(m *ice.Message, arg ...string) {
	m.Cmdy("web.code.inner", arg)
}

func init() { ice.CodeCtxCmd(inner{}) }
