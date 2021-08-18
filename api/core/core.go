package core

import (
	"shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/core/wiki"
	"shylinux.com/x/toolkits"
)

var Index = &ice.Context{Name: "core", Help: "core",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"core": {Name: "core", Help: "core", Value: kit.Data(kit.MDB_SHORT, "name")},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"core": {Name: "core", Help: "core", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
            m.Echo("hello world")
		}},
	},
}

func init() { wiki.Index.Register(Index, nil) }

