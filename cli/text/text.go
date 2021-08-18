package cli

import (
	"shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/toolkits"

	"strings"
)

var Index = &ice.Context{Name: "text", Help: "文本命令",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"text": {Name: "text", Help: "文本命令", Value: kit.Data(kit.MDB_SHORT, "name")},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"stat": {Name: "stat", Help: "stat", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
		}},
		"find": {Name: "find path args...", Help: "find", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(ice.CLI_SYSTEM, cmd, arg)
			m.Set(ice.MSG_APPEND)
		}},
		"grep": {Name: "grep text file...", Help: "grep", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			for _, v := range kit.Split(m.Cmdx(ice.CLI_SYSTEM, cmd, "-rn", arg), "\n", "\n", "\n") {
				if list := strings.SplitN(v, ":", 3); len(list) > 2 {
					m.Push("file", list[0])
					m.Push("line", list[1])
					m.Push("text", list[2])
				}
			}
		}},
		"head": {Name: "head", Help: "head", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(ice.CLI_SYSTEM, cmd, arg)
			m.Set(ice.MSG_APPEND)
		}},
		"tail": {Name: "tail", Help: "tail", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(ice.CLI_SYSTEM, cmd, arg)
			m.Set(ice.MSG_APPEND)
		}},
		"sed": {Name: "sed", Help: "sed", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(ice.CLI_SYSTEM, cmd, arg)
			m.Set(ice.MSG_APPEND)
		}},
		"awk": {Name: "awk", Help: "awk", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(ice.CLI_SYSTEM, cmd, arg)
			m.Set(ice.MSG_APPEND)
		}},
	},
}

func init() {
	cli.Index.Register(Index, nil)
}
