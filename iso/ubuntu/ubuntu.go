package ubuntu

import (
	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"
)

const UBUNTU = "ubuntu"

var Index = &ice.Context{Name: UBUNTU, Help: "ubuntu",
	Configs: map[string]*ice.Config{
		UBUNTU: {Name: UBUNTU, Help: "ubuntu", Value: kit.Data()},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		UBUNTU: {Name: "ubuntu", Help: "ubuntu", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Echo("hello ubuntu world")
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
