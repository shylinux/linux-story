package system

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	kit "shylinux.com/x/toolkits"
)

type disk struct {
	ice.Hash

	list string `name:"list hash auto" help:"磁盘"`
}

func (s disk) List(m *ice.Message, arg ...string) {
	m.Split(kit.ReplaceAll(m.Cmdx(cli.SYSTEM, "df"), "Mounted on", "Mounted_on"))
}

func init() { ice.CodeCtxCmd(disk{}) }
