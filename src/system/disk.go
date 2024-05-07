package system

import (
	"shylinux.com/x/ice"
	kit "shylinux.com/x/toolkits"
)

type disk struct {
	list string `name:"list refresh"`
}

func (s disk) List(m *ice.Message, arg ...string) {
	m.Split(kit.ReplaceAll(m.SystemCmdx("df", "-h"), "%iused", "iusedp", "Mounted on", "Mounted_on")).SortIntR("Size")
}

func init() { ice.CodeCtxCmd(disk{}) }
