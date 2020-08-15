package gdb

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/core/code"
	"github.com/shylinux/toolkits"
)

var Index = &ice.Context{Name: "gdb", Help: "gdb",
	Configs: map[string]*ice.Config{
		"gdb": {Name: "gdb", Help: "gdb", Value: kit.Data()},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"gdb": {Name: "gdb", Help: "gdb", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
            m.Echo("hello gdb world")
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
