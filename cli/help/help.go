package help

import (
	"shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/toolkits"
)

var Index = &ice.Context{Name: "help", Help: "帮助命令",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"help": {Name: "help", Help: "帮助命令", Value: kit.Data(kit.MDB_SHORT, "name")},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"man": {Name: "man", Help: "man", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(ice.CLI_SYSTEM, cmd, arg)
			m.Set(ice.MSG_APPEND)
		}},
	},
}

func init() { cli.Index.Register(Index, nil) }
