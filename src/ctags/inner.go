package ctags

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/core/code"
	kit "shylinux.com/x/toolkits"
)

type inner struct {
}

func (s inner) Man(m *ice.Message, arg ...string) {
	m.Option(cli.CMD_ENV, "COLUMNS", kit.Int(kit.Select("1920", m.Option("width")))/12)
	m.Cmdy(cli.SYSTEM, "sh", "-c", kit.Format("man %s %s|col -b", kit.Select("", arg, 1, arg[1] != "1"), arg[0]))
}
func (s inner) List(m *ice.Message, arg ...string) {
	m.Cmdy(code.INNER, arg)
}
func init() { ice.CodeCtxCmd(inner{}) }
