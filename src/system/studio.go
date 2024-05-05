package system

import "shylinux.com/x/ice"

type studio struct {
	list string `name:"list refresh"`
}

func (s studio) List(m *ice.Message, arg ...string) {
	m.Cmdy(dir{}, arg).Display("")
}

func init() { ice.CodeCtxCmd(studio{}) }
