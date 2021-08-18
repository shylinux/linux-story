package qemu

import (
	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/core/code"
	kit "shylinux.com/x/toolkits"

	"path"
)

const QEMU = "qemu"

var Index = &ice.Context{Name: QEMU, Help: "虚拟机",
	Configs: map[string]*ice.Config{
		QEMU: {Name: QEMU, Help: "虚拟机", Value: kit.Data(
			"source", "https://mirrors.tuna.tsinghua.edu.cn/git/qemu.git",
			"build", []interface{}{},
		)},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		QEMU: {Name: "qemu port=auto path=auto auto 启动 构建 下载", Help: "虚拟机", Action: map[string]*ice.Action{
			"download": {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Option(cli.CMD_DIR, m.Conf(code.INSTALL, kit.META_PATH))
				m.Cmdy(cli.SYSTEM, "git", "clone", m.Conf(QEMU, kit.META_SOURCE))
			}},
			"build": {Name: "build", Help: "构建", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, "build", m.Conf(QEMU, kit.META_SOURCE), m.Confv(QEMU, "meta.build"))
			}},
			"start": {Name: "start", Help: "启动", Hand: func(m *ice.Message, arg ...string) {
				m.Optionv("prepare", func(p string) []string {
					m.Option(cli.CMD_DIR, p)
					return []string{}
				})
				m.Cmdy(code.INSTALL, "start", m.Conf(QEMU, kit.META_SOURCE), "bin/qemu")
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(code.INSTALL, path.Base(m.Conf(QEMU, kit.META_SOURCE)), arg)
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
