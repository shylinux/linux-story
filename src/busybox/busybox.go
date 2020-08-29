package busybox

import (
	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"

	"path"
)

const BUSYBOX = "busybox"

var Index = &ice.Context{Name: BUSYBOX, Help: "busybox",
	Configs: map[string]*ice.Config{
		BUSYBOX: {Name: BUSYBOX, Help: "busybox", Value: kit.Data(
			"source", "https://busybox.net/downloads/busybox-1.32.0.tar.bz2",
		)},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmd("web.spide_rewrite", "create", "from", "https://busybox.net/downloads/busybox-1.32.0.tar.bz2", "to", "http://localhost:9020/publish/busybox-1.32.0.tar.bz2")
		}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		BUSYBOX: {Name: "busybox port=auto path=auto auto 启动 构建 下载", Help: "busybox", Action: map[string]*ice.Action{
			"download": {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, "download", m.Conf(BUSYBOX, kit.META_SOURCE))
			}},
			"build": {Name: "build", Help: "构建", Hand: func(m *ice.Message, arg ...string) {
				m.Option("install", "_install")
				m.Option("prepare", func(p string) {
					m.Option(cli.CMD_DIR, p)
					m.Cmdy(cli.SYSTEM, "make", "defconfig")
				})
				m.Cmdy(code.INSTALL, "build", m.Conf(BUSYBOX, kit.META_SOURCE))
			}},
			"start": {Name: "start", Help: "启动", Hand: func(m *ice.Message, arg ...string) {
				m.Optionv("prepare", func(p string) []string {
					m.Option(cli.CMD_DIR, p)
					m.Cmd(cli.SYSTEM, "ln", "-s", "_install", "install")
					return []string{}
				})
				m.Cmdy(code.INSTALL, "start", m.Conf(BUSYBOX, kit.META_SOURCE), "bin/busybox")
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(code.INSTALL, path.Base(m.Conf(BUSYBOX, kit.META_SOURCE)), arg)
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
