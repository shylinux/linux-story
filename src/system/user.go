package system

import (
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/aaa"
	"shylinux.com/x/icebergs/base/cli"
	kit "shylinux.com/x/toolkits"
)

type user struct {
	list string `name:"list username auto" help:"用户"`
}

func (s user) List(m *ice.Message, arg ...string) {
	kit.For(kit.SplitLine(m.Cmdx(cli.SYSTEM, "getent", "passwd")), func(text string) {
		ls := strings.Split(text, ":")
		m.Push(aaa.USERNAME, ls[0])
		m.Push("uid", ls[2])
		m.Push("gid", ls[3])
		m.Push("home", ls[5])
		m.Push("shell", ls[6])
	})
	m.SortInt("uid")
}

func init() { ice.Cmd("web.code.system.user", user{}) }
