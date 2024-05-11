package system

import (
	"runtime"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
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
	if runtime.GOOS == cli.DARWIN {
		kit.For(kit.SplitLine(m.SystemCmdx(cli.SH, "-c", `dscacheutil -q group`)), func(text string) {
			if ls := strings.SplitN(text, ":", 2); len(ls) > 1 {
				m.Push(ls[0], strings.TrimSpace(ls[1]))
			}
		})
	} else {
		kit.For(kit.SplitLine(m.SystemCmdx(GETENT, GROUP)), func(text string) {
			if ls := strings.Split(text, ":"); len(ls) > 2 {
				if m.Push("group,password,gid,username", ls); kit.Int(ls[2]) > 999 {
					m.PushButton(s.Remove)
				} else {
					m.PushButton()
				}
			}
		})
	}
	m.SortIntR(GID)
}

func init() { ice.CodeCtxCmd(group{}) }
