package android

import (
	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/core/code"
	kit "shylinux.com/x/toolkits"
)

var Index = &ice.Context{Name: "android", Help: "android",
	Configs: map[string]*ice.Config{
		"android": {Name: "android", Help: "android", Value: kit.Data()},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"android": {Name: "android", Help: "android", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
            m.Echo("hello android world")
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
