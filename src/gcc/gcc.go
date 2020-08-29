package gcc

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/core/code"
	"github.com/shylinux/toolkits"
)

const GCC = "gcc"

var Index = &ice.Context{Name: GCC, Help: "编译器",
	Configs: map[string]*ice.Config{
		GCC: {Name: GCC, Help: "编译器", Value: kit.Data(
			"source", "http://mirrors.aliyun.com/gnu/gcc/gcc-4.8.5/gcc-4.8.5.tar.gz",
			"build", []interface{}{
				"--enable-languages=c",
			},
		)},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		GCC: {Name: "gcc port=auto path=auto auto 启动:button 构建:button 下载:button", Help: "命令行", Action: map[string]*ice.Action{
			"download": {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, "download", m.Conf(GCC, kit.META_SOURCE))
			}},
			"build": {Name: "build", Help: "构建", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, "build", m.Conf(GCC, kit.META_SOURCE), m.Confv(GCC, "meta.build"))
			}},
			"start": {Name: "start", Help: "启动", Hand: func(m *ice.Message, arg ...string) {
				m.Optionv("prepare", func(p string) []string {
					m.Option(cli.CMD_DIR, p)
					return []string{}
				})
				m.Cmdy(code.INSTALL, "start", m.Conf(GCC, kit.META_SOURCE), "bin/gcc")
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(code.INSTALL, path.Base(m.Conf(GCC, kit.META_SOURCE)), arg)
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
