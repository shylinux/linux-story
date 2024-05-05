package system

import "shylinux.com/x/ice"

type studio struct {
	tools string `data:"web.code.system.port,web.code.system.proc,web.code.system.user,web.code.system.disk"`
	list  string `name:"list refresh" icon:"linux.png"`
}

func (s studio) List(m *ice.Message, arg ...string) {
	if m.Cmdy(dir{}, arg); len(arg) == 0 {
		m.Display("")
	}
}

func init() { ice.CodeCtxCmd(studio{}) }
