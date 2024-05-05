package ctags

import (
	"path"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	kit "shylinux.com/x/toolkits"
)

type nroff struct {
	ice.Code
	list string `name:"inner path=src/@key file=main.go@key line=1 auto" help:"编辑器"`
}

func (s nroff) List(m *ice.Message, arg ...string) {
	m.Cmdy("web.code.inner", arg)
	if len(arg) > 1 && arg[0] != ice.ACTION && m.Result() == "" {
		m.Option(cli.CMD_ENV, "COLUMNS", kit.Int(kit.Select("1920", m.Option("width")))/12)
		if m.Cmdy(cli.SYSTEM, "sh", "-c", kit.Format("man -l %s|col -b", path.Join(arg[0], arg[1]))); m.Append(cli.CMD_OUT) != "" {
			m.SetResult(m.Append(cli.CMD_OUT))
		}
	}
	m.Option("plug", m.Config("show.plug"))
	m.Option("exts", m.Config("show.exts"))
	m.Option("tabs", m.Config("show.tabs"))
}

func init() { ice.CodeCtxCmd(nroff{}) }
