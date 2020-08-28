package kernel

import (
	"os"
	"path"

	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
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
		KERNEL: {Name: "kernel port=auto path=auto auto 启动:button 构建:button 下载:button", Help: "内核", Action: map[string]*ice.Action{
			"download": {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, "download", m.Conf(KERNEL, kit.META_SOURCE))
			}},
			"build": {Name: "build", Help: "构建", Hand: func(m *ice.Message, arg ...string) {
				m.Optionv("prepare", func(p string) {
					m.Option(cli.CMD_DIR, p)
					m.Cmdy(cli.SYSTEM, "make", "defconfig")
				})
				// m.Cmdy(code.INSTALL, "build", m.Conf(KERNEL, kit.META_SOURCE))

				p := path.Join(m.Conf(code.INSTALL, kit.META_PATH), kit.TrimExt(m.Conf(KERNEL, kit.META_SOURCE)))
				os.MkdirAll(path.Join(p, "install"), ice.MOD_DIR)
				m.Option(cli.CMD_DIR, p)
				m.Cmdy(cli.SYSTEM, "cp", "arch/x86_64/boot/bzImage", "install/linux")
				m.Cmdy(cli.SYSTEM, "dd", "if=/dev/zero", "of=install/rootfs.img", "bs=1M", "count=100")
				m.Cmdy(cli.SYSTEM, "mkfs.ext4", "install/rootfs.img")
				// m.Cmdy(cli.SYSTEM, "yum", "install", "-y", "e4fsprogs")
			}},
			"start": {Name: "start", Help: "启动", Hand: func(m *ice.Message, arg ...string) {
				m.Optionv("prepare", func(p string) []string { return []string{} })
				m.Cmdy(code.INSTALL, "start", m.Conf(KERNEL, kit.META_SOURCE),
					"qemu-system-x86_64", "-kernel", "linux", "-nographic")
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(code.INSTALL, m.Conf(KERNEL, kit.META_SOURCE), arg)
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
