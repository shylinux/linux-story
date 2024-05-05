package system

import (
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/web/html"
	kit "shylinux.com/x/toolkits"
)

type proc struct {
	list string `name:"list PID auto"`
}

func (s proc) List(m *ice.Message, arg ...string) {
	head := []string{}
	kit.For(kit.SplitLine(m.Cmdx(cli.SYSTEM, "ps", "aux")), func(text string, index int) {
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
		m.Push(head[len(head)-1], strings.Join(cmds, " "))
	})
	m.RewriteAppend(func(value string, key string, index int) string {
		switch key {
		case "VSZ", "RSS":
			value = kit.FmtSize(kit.Int(value) * 1024)
		}
		return value
	})
	m.Action(html.FILTER)
	m.StatusTimeCountStats("USER")
	m.SortIntR("RSS")
	m.Option("table.checkbox", "true")
}

func init() { ice.CodeCtxCmd(proc{}) }
