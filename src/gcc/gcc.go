package gcc

import (
	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"

	"path"
)

const GCC = "gcc"

var Index = &ice.Context{Name: GCC, Help: "编译器",
	Configs: map[string]*ice.Config{
		GCC: {Name: GCC, Help: "编译器", Value: kit.Data(
			"source", "http://mirrors.aliyun.com/gnu/gcc/gcc-4.8.5/gcc-4.8.5.tar.gz",
			"build", []interface{}{
				"--enable-languages=c,c++",
				"--disable-multilib",
				"--disable-checking",
			},
		)},
	},
	Commands: map[string]*ice.Command{
		GCC: {Name: "gcc path auto build install", Help: "编译器", Action: map[string]*ice.Action{
			code.INSTALL: {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, code.INSTALL, m.Conf(GCC, kit.META_SOURCE))
			}},
			cli.BUILD: {Name: "build", Help: "构建", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, cli.BUILD, m.Conf(GCC, kit.META_SOURCE), m.Confv(GCC, "meta.build"))
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(code.INSTALL, path.Base(m.Conf(GCC, kit.META_SOURCE)), arg)
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
