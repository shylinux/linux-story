package kernel

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/core/code"
	"github.com/shylinux/toolkits"
)

const KERNEL = "kernel"

var Index = &ice.Context{Name: KERNEL, Help: "内核",
	Configs: map[string]*ice.Config{
		KERNEL: {Name: KERNEL, Help: "内核", Value: kit.Data(
			"source", "https://mirrors.tuna.tsinghua.edu.cn/kernel/v3.x/linux-3.10.1.tar.gz",
		)},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		KERNEL: {Name: "kernel auto 下载:button", Help: "内核", Action: map[string]*ice.Action{
			"download": {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, "download", m.Conf(KERNEL, "meta.source"))
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Echo("hello kernel world")
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
