package busybox

import (
	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"

	"path"
	"runtime"
	"strings"
)

const BUSYBOX = "busybox"

var Index = &ice.Context{Name: BUSYBOX, Help: "busybox",
	Configs: map[string]*ice.Config{
		BUSYBOX: {Name: BUSYBOX, Help: "busybox", Value: kit.Data(
			"windows", "https://busybox.net/downloads/busybox-1.32.0.tar.bz2",
			"darwin", "https://busybox.net/downloads/busybox-1.32.0.tar.bz2",
			"linux", "https://busybox.net/downloads/busybox-1.32.0.tar.bz2",
		)},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		BUSYBOX: {Name: "busybox port=auto auto 启动:button 编译:button 下载:button cmd:textarea", Help: "busybox", Action: map[string]*ice.Action{
			"download": {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, "download", m.Conf(BUSYBOX, kit.Keys(kit.MDB_META, runtime.GOOS)))
			}},
			"compile": {Name: "compile", Help: "编译", Hand: func(m *ice.Message, arg ...string) {
				name := path.Base(strings.TrimSuffix(strings.TrimSuffix(m.Conf(BUSYBOX, kit.Keys(kit.MDB_META, runtime.GOOS)), ".tar.bz2"), "zip"))
				m.Option(cli.CMD_DIR, path.Join(m.Conf(code.INSTALL, kit.META_PATH), name))
				m.Cmdy(cli.SYSTEM, "make", "defconfig")
				m.Cmdy(cli.SYSTEM, "make", "-j8")
				m.Cmdy(cli.SYSTEM, "make", "install")
			}},
			"start": {Name: "start", Help: "启动", Hand: func(m *ice.Message, arg ...string) {
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
