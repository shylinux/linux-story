package file

import (
	"shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/toolkits"

	"os"
)

var Index = &ice.Context{Name: "file", Help: "文件命令",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"file": {Name: "file", Help: "文件命令", Value: kit.Data(kit.MDB_SHORT, "name")},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"dir": {Name: "dir", Help: "dir", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			if len(arg) > 0 {
				if s, e := os.Stat(arg[0]); e == nil && !s.IsDir() {
					m.Cmdy("cat", arg)
					return
				}
			}

			m.Cmdy(ice.CLI_SYSTEM, "ls", arg)
			m.Set(ice.MSG_APPEND)
		}},
		"cat": {Name: "cat", Help: "cat", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			if len(arg) == 0 {
				return
			}

			if s, e := os.Stat(arg[0]); e == nil && s.IsDir() {
				m.Cmdy("dir", arg)
				return
			}

			m.Cmdy(ice.CLI_SYSTEM, cmd, arg)
			m.Set(ice.MSG_APPEND)
		}},
	},
}

func init() { cli.Index.Register(Index, nil) }
