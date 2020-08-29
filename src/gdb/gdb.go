package gdb

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/core/code"
	"github.com/shylinux/toolkits"

	"path"
)

const GDB = "gdb"

var Index = &ice.Context{Name: GDB, Help: "调试器",
	Configs: map[string]*ice.Config{
		GDB: {Name: GDB, Help: "调试器", Value: kit.Data(
			"source", "http://mirrors.aliyun.com/gnu/gdb/gdb-7.6.1.tar.bz2",
			"build", []interface{}{},
		)},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		GDB: {Name: "gcc port=auto path=auto auto 启动:button 构建:button 下载:button", Help: "编译器", Action: map[string]*ice.Action{
			"download": {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, "download", m.Conf(GDB, kit.META_SOURCE))
			}},
			"build": {Name: "build", Help: "构建", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, "build", m.Conf(GDB, kit.META_SOURCE), m.Confv(GDB, "meta.build"))
			}},
			"start": {Name: "start", Help: "启动", Hand: func(m *ice.Message, arg ...string) {
				m.Optionv("prepare", func(p string) []string {
					m.Option(cli.CMD_DIR, p)
					return []string{}
				})
				m.Cmdy(code.INSTALL, "start", m.Conf(GDB, kit.META_SOURCE), "bin/gdb")
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(code.INSTALL, path.Base(m.Conf(GDB, kit.META_SOURCE)), arg)
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
