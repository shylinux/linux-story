package system

import (
	"strings"

	"shylinux.com/x/ice"
	kit "shylinux.com/x/toolkits"
)

const (
	USERADD = "useradd"
	USERDEL = "userdel"
	USERMOD = "usermod"
	PASSWD  = "passwd"
	GETENT  = "getent"
	WHOAMI  = "whoami"
	SUDO    = "sudo"

	USERNAME = "username"
	PASSWORD = "password"
	COMMENT  = "comment"
	UID      = "uid"
	GID      = "gid"
)

type user struct {
	create string `name:"create username* password comment"`
	list   string `name:"list username auto"`
}

func (s user) Create(m *ice.Message, arg ...string) {
	m.SystemCmd(SUDO, USERADD, m.OptionArgs(COMMENT, PASSWORD), m.Option(USERNAME))
}
func (s user) Remove(m *ice.Message, arg ...string) {
	m.SystemCmd(SUDO, USERDEL, "-r", m.Option(USERNAME))
}
func (s user) List(m *ice.Message, arg ...string) {
	who := m.SystemCmdx(WHOAMI)
	kit.For(kit.SplitLine(m.SystemCmdx(GETENT, PASSWD)), func(text string) {
		if ls := strings.Split(text, ":"); len(ls) > 6 {
			if m.Push("username,password,uid,gid,comment,home,shell", ls); kit.Int(ls[2]) > 999 && ls[0] != who {
				m.PushButton(s.Remove)
			} else {
				m.PushButton()
			}
		}
	})
	m.ActionFilter(s.Create).SortIntR(UID)
}

func init() { ice.CodeCtxCmd(user{}) }
