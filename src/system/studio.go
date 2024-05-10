package system

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

type studio struct {
	dir    dir
	online string `data:"true"`
	plugs  string `name:"plugs path* index args"`
	favor  string `name:"favor path* zone type name text"`
	list   string `name:"list refresh" icon:"studio.png"`
}

func (s studio) Inputs(m *ice.Message, arg ...string) {
	switch m.Option(ctx.ACTION) {
	case kit.FuncName(s.Plugs):
		m.Cmdy(plugs{}, s.Inputs, arg)
	case kit.FuncName(s.Favor):
		m.Cmdy(favor{}, s.Inputs, arg)
	}
}
func (s studio) Plugs(m *ice.Message, arg ...string) {
	m.Cmdy(plugs{}, plugs{}.Create, m.OptionSimple("path,index,args")).ProcessHold()
}
func (s studio) Favor(m *ice.Message, arg ...string) {
	m.Cmdy(favor{}, favor{}.Create, m.OptionSimple("path,type,name,text")).ProcessHold()
}
func (s studio) List(m *ice.Message, arg ...string) {
	if m.Cmdy(dir{}, arg).PushAction(s.Plugs, s.Favor, s.dir.Upload, s.dir.Trash).Action(); len(arg) == 0 {
		m.Display("").DisplayCSS("")
		kit.If(m.Config(ctx.TOOLS) == "", func() { m.Toolkit(favor{}, plugs{}, web.XTERM) })
		m.StatusTimeCount(nfs.VERSION, m.SystemCmdx("uname", "-sr"))
	}
}

func init() { ice.CodeCtxCmd(studio{}) }
