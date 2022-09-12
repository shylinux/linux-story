package ctags

import (
	"path"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/nfs"
)

type inner struct {
	ice.Code
	ctags string `name:"ctags" help:"索引"`
	list  string `name:"inner path=src/@key file=main.go@key line=1 auto" help:"编辑器"`
}

func (s inner) Tags(m *ice.Message, arg ...string) {
	if !nfs.ExistsFile(m, path.Join(m.Option(nfs.PATH), "tags")) {
		s.System(m.Spawn(), m.Option(nfs.PATH), "ctags", "-a", "-R")
	}
	if m.Cmdy("web.code.inner", "tags", arg); m.Length() > 0 {
		return // 索引
	}
	for _, n := range []string{"3", "2", ""} {
		if msg := s.System(m.Spawn(), "", cli.MAN, n, arg[0]); cli.IsSuccess(msg) && !strings.HasPrefix(msg.Result(), "No manual entry for") {
			m.Push(nfs.PATH, cli.MAN)
			m.Push(nfs.FILE, arg[0])
			m.Push(nfs.LINE, n)
			return // 手册
		}
	}
}
func (s inner) Man(m *ice.Message, arg ...string) {
	m.Cmdy(cli.SYSTEM, cli.MAN, arg)
}
func (s inner) List(m *ice.Message, arg ...string) {
	m.Cmdy("web.code.inner", arg)
	m.Option("plug", m.Config("show.plug"))
	m.Option("exts", m.Config("show.exts"))
	m.Option("tabs", m.Config("show.tabs"))
}

func init() { ice.CodeCtxCmd(inner{}) }
