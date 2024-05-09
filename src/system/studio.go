package system

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

type studio struct {
	dir    dir
	online string `data:"true"`
	favor  string `name:"favor path* zone type name text"`
	list   string `name:"list refresh" icon:"studio.png"`
}

func (s studio) Inputs(m *ice.Message, arg ...string) {
	m.Cmdy(favor{}, s.Inputs, arg)
}
func (s studio) Favor(m *ice.Message, arg ...string) {
	m.Cmdy(favor{}, favor{}.Create, m.OptionSimple("path,type,name,text")).ProcessHold()
}
func (s studio) List(m *ice.Message, arg ...string) {
	if m.Cmdy(dir{}, arg).PushAction(s.Favor, s.dir.Upload, s.dir.Trash).Action(); len(arg) == 0 {
		m.Display("").DisplayCSS("")
		kit.If(m.Config(ctx.TOOLS) == "", func() { m.Toolkit(user{}, group{}, favor{}) })
		m.StatusTimeCount(nfs.VERSION, m.SystemCmdx("uname", "-sr"))
	}
}

func init() { ice.CodeCtxCmd(studio{}) }
