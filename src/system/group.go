package system

import (
	"strings"

	"shylinux.com/x/ice"
	kit "shylinux.com/x/toolkits"
)

const (
	GROUP    = "group"
	GROUPADD = "groupadd"
	GROUPDEL = "groupdel"
	GROUPMOD = "groupmod"
)

type group struct {
	create string `name:"create group* password"`
	list   string `name:"list group auto"`
}

func (s group) Create(m *ice.Message, arg ...string) {
	m.SystemCmd(SUDO, GROUPADD, m.OptionArgs(PASSWORD), m.Option(GROUP))
}
func (s group) Remove(m *ice.Message, arg ...string) {
	m.SystemCmd(SUDO, GROUPDEL, m.Option(GROUP))
}
func (s group) List(m *ice.Message, arg ...string) {
	kit.For(kit.SplitLine(m.SystemCmdx(GETENT, GROUP)), func(text string) {
		if ls := strings.Split(text, ":"); len(ls) > 2 {
			if m.Push("group,password,gid,username", ls); kit.Int(ls[2]) > 999 {
				m.PushButton(s.Remove)
			} else {
				m.PushButton()
			}
		}
	})
	m.ActionFilter(s.Create).SortIntR(GID)
}

func init() { ice.CodeCtxCmd(group{}) }
