package user

import (
	"shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/core/wiki"
	"shylinux.com/x/toolkits"
)

var Index = &ice.Context{Name: "user", Help: "user",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"user": {Name: "user", Help: "user", Value: kit.Data(kit.MDB_SHORT, "name")},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"user": {Name: "user", Help: "user", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
            m.Echo("hello world")
		}},
	},
}

func init() { wiki.Index.Register(Index, nil) }

