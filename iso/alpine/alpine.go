package alpine

import (
	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"
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
