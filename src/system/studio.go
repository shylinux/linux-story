package system

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

type studio struct {
	list string `name:"list refresh" icon:"linux.png"`
}

func (s studio) List(m *ice.Message, arg ...string) {
	if m.Cmdy(dir{}, arg); len(arg) == 0 {
		m.StatusTimeCount(nfs.VERSION, m.SystemCmdx("uname", "-r")).Display("")
		kit.If(m.Config(ctx.TOOLS) == "", func() { m.Toolkit(port{}, proc{}, user{}, disk{}) })
	}
}

func init() { ice.CodeCtxCmd(studio{}) }
