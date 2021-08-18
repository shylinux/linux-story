package context

import (
	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/core/code"
	kit "shylinux.com/x/toolkits"
)

const CONTEXT = "context"

var Index = &ice.Context{Name: CONTEXT, Help: "context",
	Configs: map[string]*ice.Config{
		CONTEXT: {Name: CONTEXT, Help: "context", Value: kit.Data()},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		CONTEXT: {Name: "context", Help: "context", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Echo("hello context world")
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
