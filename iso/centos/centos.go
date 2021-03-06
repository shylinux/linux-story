package centos

import (
	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"
)

const CENTOS = "centos"

var Index = &ice.Context{Name: CENTOS, Help: "centos",
	Configs: map[string]*ice.Config{
		CENTOS: {Name: CENTOS, Help: "centos", Value: kit.Data()},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		CENTOS: {Name: "centos", Help: "centos", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Echo("hello centos world")
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
