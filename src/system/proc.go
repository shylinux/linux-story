package system

import (
	"path"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	kit "shylinux.com/x/toolkits"
)

type proc struct {
	kill string `name:"kill signal"`
	list string `name:"list PID auto"`
}

func (s proc) Kill(m *ice.Message, arg ...string) {
	m.SystemCmd(cli.KILL, "-"+m.OptionDefault("signal", "9"), m.Option("PID"))
}
func (s proc) List(m *ice.Message, arg ...string) {
	head := []string{}
	kit.For(kit.SplitLine(m.SystemCmdx("ps", "aux")), func(text string, index int) {
		if index == 0 {
			head = kit.SplitWord(text)
			return
		}
		cmds := []string{}
		kit.For(kit.SplitWord(text), func(text string, index int) {
			if index < len(head)-1 {
				m.Push(head[index], text)
			} else {
				cmds = append(cmds, text)
			}
		})
		if strings.HasPrefix(cmds[0], "[") {
			m.Push(cli.CMD, cmds[0])
		} else {
			m.Push(cli.CMD, strings.TrimSuffix(path.Base(cmds[0]), ":"))
		}
		m.Push(head[len(head)-1], kit.JoinWord(cmds...))
		m.PushButton(s.Kill)
	})
	m.RewriteAppend(func(value string, key string, index int) string {
		switch key {
		case "VSZ", "RSS":
			value = kit.FmtSize(kit.Int(value) * 1024)
		}
		return value
	})
	m.StatusTimeCountStats("USER").SortIntR("RSS")
}

func init() { ice.CodeCtxCmd(proc{}) }
