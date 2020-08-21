package qemu

import (
	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"

	"path"
	"strings"
)

const QEMU = "qemu"

var Index = &ice.Context{Name: QEMU, Help: "虚拟机",
	Configs: map[string]*ice.Config{
		QEMU: {Name: QEMU, Help: "虚拟机", Value: kit.Data(
			"source", "https://mirrors.tuna.tsinghua.edu.cn/git/qemu.git",
			"windows", "http://download.redis.io/releases/redis-5.0.4.tar.gz",
			"darwin", "http://download.redis.io/releases/redis-5.0.4.tar.gz",
			"linux", "http://download.redis.io/releases/redis-5.0.4.tar.gz",
		)},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		QEMU: {Name: "qemu port=auto auto 启动:button 编译:button 下载:button cmd:textarea", Help: "服务器", Action: map[string]*ice.Action{
			"download": {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Option(cli.CMD_DIR, m.Conf(code.INSTALL, kit.META_PATH))
				m.Cmdy(cli.SYSTEM, "git", "clone", m.Conf(QEMU, kit.Keys(kit.MDB_META, "source")))
			}},
			"compile": {Name: "compile", Help: "编译", Hand: func(m *ice.Message, arg ...string) {
				m.Option(cli.CMD_DIR, path.Join(m.Conf(code.INSTALL, kit.META_PATH), QEMU))
				m.Cmdy(cli.SYSTEM, "./configure", "--prefix=./install")
				m.Cmdy(cli.SYSTEM, "make", "-j8")
			}},
			"start": {Name: "start", Help: "启动", Hand: func(m *ice.Message, arg ...string) {
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			if len(arg) > 0 && arg[0] != "" {
				m.Cmdy(cli.SYSTEM, "bin/redis-cli", "-p", arg[0], kit.Split(kit.Select("info", arg, 1)))
				return
			}

			m.Cmd(cli.DAEMON).Table(func(index int, value map[string]string, head []string) {
				if strings.HasPrefix(value[kit.MDB_NAME], "bin/redis") {
					m.Push(kit.MDB_TIME, value[kit.MDB_TIME])
					m.Push(kit.MDB_PORT, path.Base(value[kit.MDB_DIR]))
					m.Push(kit.MDB_DIR, value[kit.MDB_DIR])
					m.Push(kit.MDB_STATUS, value[kit.MDB_STATUS])
					m.Push(kit.MDB_PID, value[kit.MDB_PID])
					m.Push(kit.MDB_NAME, value[kit.MDB_NAME])
				}
			})
			m.Sort("time", "time_r")
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
