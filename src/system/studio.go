package system

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

type studio struct {
	_dir   dir
	_favor favor
	_plugs plugs
	online string `data:"true"`
	plugs  string `name:"plugs path* index args"`
	favor  string `name:"favor path* zone type name text"`
	list   string `name:"list refresh" icon:"studio.png"`
}

func (s studio) Init(m *ice.Message, arg ...string) {
	web.AddPortalProduct(m.Message, "System Studio", `
一款网页版的系统工作台，有文件管理、用户管理、进程管理等功能。
`, 10.0)
}
func (s studio) Inputs(m *ice.Message, arg ...string) {
	switch m.Option(ctx.ACTION) {
	case kit.FuncName(s.Plugs):
		m.Cmdy(s._plugs, s.Inputs, arg)
	case kit.FuncName(s.Favor):
		m.Cmdy(s._favor, s.Inputs, arg)
	}
}
func (s studio) Plugs(m *ice.Message, arg ...string) {
	m.Cmdy(s._plugs, s._plugs.Create, m.OptionSimple("path,index,args")).ProcessHold()
}
func (s studio) Favor(m *ice.Message, arg ...string) {
	m.Cmdy(s._favor, s._favor.Create, m.OptionSimple("path,type,name,text")).ProcessHold()
}
func (s studio) List(m *ice.Message, arg ...string) {
	if m.Cmdy(s._dir, arg).PushAction(s.Plugs, s.Favor, s._dir.Upload, s._dir.Trash).Action(); len(arg) == 0 {
		m.Display("").DisplayCSS("")
		kit.If(m.Config(ctx.TOOLS) == "", func() { m.Toolkit(favor{}, plugs{}, tools{}) })
		m.StatusTimeCount(nfs.VERSION, m.SystemCmdx("uname", "-sr"))
	}
}

func init() { ice.CodeCtxCmd(studio{}) }
