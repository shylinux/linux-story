package system

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/web/html"
	kit "shylinux.com/x/toolkits"
)

type proc struct {
	list string `name:"list PID auto"`
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
		m.Push(head[len(head)-1], kit.JoinWord(cmds...))
	})
	m.RewriteAppend(func(value string, key string, index int) string {
		switch key {
		case "VSZ", "RSS":
			value = kit.FmtSize(kit.Int(value) * 1024)
		}
		return value
	})
	m.Action(html.FILTER).StatusTimeCountStats("USER").SortIntR("RSS")
}

func init() { ice.CodeCtxCmd(proc{}) }
