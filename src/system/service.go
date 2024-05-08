package system

import (
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/mdb"
	kit "shylinux.com/x/toolkits"
)

type service struct {
	list string `name:"list name auto"`
}

func (s service) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 || arg[0] == "" {
		ls := kit.SplitLine(m.SystemCmdx("systemctl", "list-units", "--type=service"))
		head := kit.Split(ls[0])
		for _, text := range ls[1:] {
			ls := kit.Split(strings.TrimPrefix(text, "●"))
			if len(ls) < 1 {
				break
			}
			m.Push(mdb.NAME, ls[0])
			m.Push(head[1], ls[1])
			m.Push(head[2], ls[2])
			m.Push(head[3], ls[3])
			m.Push(head[4], kit.JoinWord(ls[4:]...))
		}
		m.ActionFilter()
		m.StatusTimeCountStats("LOAD", "ACTIVE", "SUB")
		m.SortStrR("SUB")
	} else {
		kit.For(kit.SplitLine(m.SystemCmdx("systemctl", "show", arg[0])), func(text string) {
			m.Push("name,value", strings.SplitN(text, "=", 2))
		})
	}
}

func init() { ice.CodeCtxCmd(service{}) }
