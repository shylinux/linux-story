package system

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
)

type port struct {
	list string `name:"list hash auto"`
}

func (s port) List(m *ice.Message, arg ...string) {
	m.Cmdy(cli.SYSTEM, "netstat", "-an")
}

func init() { ice.CodeCtxCmd(port{}) }
