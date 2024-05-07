package system

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

type studio struct {
	dir  dir
	list string `name:"list refresh" icon:"studio.png"`
}

func (s studio) List(m *ice.Message, arg ...string) {
	if m.Cmdy(dir{}, arg).PushAction(s.dir.Upload, s.dir.Trash).Action(); len(arg) == 0 {
		m.Display("").DisplayCSS("")
		kit.If(m.Config(ctx.TOOLS) == "", func() { m.Toolkit(port{}, proc{}, user{}, disk{}) })
		m.StatusTimeCount(nfs.VERSION, m.SystemCmdx("uname", "-sr"))
	}
}

func init() { ice.CodeCtxCmd(studio{}) }
