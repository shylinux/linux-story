package alpine

import (
	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/core/code"
	kit "shylinux.com/x/toolkits"
)

var Index = &ice.Context{Name: "alpine", Help: "alpine",
	Configs: map[string]*ice.Config{
		"alpine": {Name: "alpine", Help: "alpine", Value: kit.Data()},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"alpine": {Name: "alpine", Help: "alpine", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
            m.Echo("hello alpine world")
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
