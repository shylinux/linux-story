package system

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
)

type proc struct {
	ice.Hash
	list string `name:"list hash auto" help:"示例"`
}

func (s proc) List(m *ice.Message, arg ...string) {
	m.Cmdy(cli.SYSTEM, "ps")
}

func init() { ice.CodeCtxCmd(proc{}) }
