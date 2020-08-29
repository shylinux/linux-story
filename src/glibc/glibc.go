package glibc

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/core/code"
	"github.com/shylinux/toolkits"

	"path"
)

const GLIBC = "glibc"

var Index = &ice.Context{Name: GLIBC, Help: "标准库",
	Configs: map[string]*ice.Config{
		GLIBC: {Name: GLIBC, Help: "标准库", Value: kit.Data(
			"source", "http://mirrors.aliyun.com/gnu/glibc/glibc-2.17.tar.gz",
			"build", []interface{}{
				"--enable-languages=c,c++",
				"--disable-multilib",
				"--disable-checking",
			},
		)},
	},
	Commands: map[string]*ice.Command{
		GLIBC: {Name: "gcc port=auto path=auto auto 启动:button 构建:button 下载:button", Help: "标准库", Action: map[string]*ice.Action{
			"download": {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, "download", m.Conf(GLIBC, kit.META_SOURCE))
			}},
			"build": {Name: "build", Help: "构建", Hand: func(m *ice.Message, arg ...string) {
				p := m.Option(cli.CMD_DIR, path.Join(m.Conf(code.INSTALL, kit.META_PATH), kit.TrimExt(m.Conf(GLIBC, kit.META_SOURCE)), "install"))
				m.Cmdy(cli.SYSTEM, "../configure", "--prefix="+kit.Path(path.Dir(p)))
				m.Cmdy(cli.SYSTEM, "make", "-j8")
				m.Cmdy(cli.SYSTEM, "make", "install")
			}},
			"start": {Name: "start", Help: "启动", Hand: func(m *ice.Message, arg ...string) {
				m.Optionv("prepare", func(p string) []string {
					m.Option(cli.CMD_DIR, p)
					return []string{}
				})
				m.Cmdy(code.INSTALL, "start", m.Conf(GLIBC, kit.META_SOURCE), "./testrun.sh", "glibc")
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(code.INSTALL, path.Base(m.Conf(GLIBC, kit.META_SOURCE)), arg)
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
