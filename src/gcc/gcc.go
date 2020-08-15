package gcc

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/core/code"
	"github.com/shylinux/toolkits"
)

const GCC = "gcc"

var Index = &ice.Context{Name: GCC, Help: "编译器",
	Configs: map[string]*ice.Config{
		GCC: {Name: GCC, Help: "编译器", Value: kit.Data()},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		GCC: {Name: GCC, Help: "编译器", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Echo("hello gcc world")
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
