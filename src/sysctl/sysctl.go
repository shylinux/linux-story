package sysctl

import (
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	kit "shylinux.com/x/toolkits"
)

type sysctl struct {
	ice.Code
}

func (s sysctl) List(m *ice.Message, arg ...string) {
	m.OptionFields("detail")
	data := kit.Dict()
	for _, l := range strings.Split(s.System(m.Spawn(), "", "sysctl", "-a").Append(cli.CMD_OUT), ice.NL) {
		if ls := kit.Split(l, " =", " "); len(ls) > 1 {
			kit.Value(data, ls[0], ls[1])
			m.Push(ls[0], ls[1])
		}
	}
	m.Echo(kit.Formats(data))
	// ctx.DisplayStoryJSON(m.Message)
}

func init() { ice.CodeCtxCmd(sysctl{}) }
