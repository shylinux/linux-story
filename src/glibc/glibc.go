package glibc

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/core/code"
	"github.com/shylinux/toolkits"
)

const GLIBC = "glibc"

var Index = &ice.Context{Name: GLIBC, Help: "标准库",
	Configs: map[string]*ice.Config{
		GLIBC: {Name: GLIBC, Help: "标准库", Value: kit.Data()},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		GLIBC: {Name: GLIBC, Help: "标准库", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Echo("hello glibc world")
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
