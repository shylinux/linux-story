package kernel

import (
	"path"

	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/base/tcp"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"
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

		KERNEL: {Name: "kernel auto 启动:button 编译:button 下载:button", Help: "内核", Action: map[string]*ice.Action{
			"download": {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, "download", m.Conf(KERNEL, "meta.source"))
			}},
			"prepare": {Name: "prepare", Help: "安装", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(cli.SYSTEM, "yum", "install", "-y", "e4fsprogs")
			}},
			"compile": {Name: "compile", Help: "编译", Hand: func(m *ice.Message, arg ...string) {
				name := kit.TrimExt(m.Conf(KERNEL, kit.Keys(kit.MDB_META, "source")))
				m.Option(cli.CMD_DIR, path.Join(m.Conf(code.INSTALL, kit.META_PATH), name))
				m.Cmdy(cli.SYSTEM, "make", "defconfig")
				m.Cmdy(cli.SYSTEM, "make", "-j8")
			}},
			"start": {Name: "start", Help: "启动", Hand: func(m *ice.Message, arg ...string) {
				port := m.Cmdx(tcp.PORT, "select")
				p := path.Join(m.Conf(cli.DAEMON, kit.META_PATH), port)
				m.Option(cli.CMD_DIR, p)
				m.Cmdy(cli.SYSTEM, "dd", "if=/dev/zero", "of=rootfs.img", "bs=1M", "count=100")
				m.Cmdy(cli.SYSTEM, "mkfs.ext4", "rootfs.img")
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Echo("hello kernel world")
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
