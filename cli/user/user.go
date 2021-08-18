package user

import (
	"shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/toolkits"
)

var Index = &ice.Context{Name: "user", Help: "用户命令",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"user": {Name: "user", Help: "用户命令", Value: kit.Data(kit.MDB_SHORT, "name")},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"who": {Name: "who", Help: "who", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(ice.CLI_SYSTEM, cmd, arg)
			m.Set(ice.MSG_APPEND)
		}},
	},
}

func init() { cli.Index.Register(Index, nil) }
