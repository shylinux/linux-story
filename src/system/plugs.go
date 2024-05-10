package system

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/ctx"
	kit "shylinux.com/x/toolkits"
)

type plugs struct {
	ice.Hash
	action string `data:"show"`
	export string `data:"true"`
	short  string `data:"path"`
	field  string `data:"time,path,index,args"`
	list   string `name:"list path auto"`
}

func (s plugs) List(m *ice.Message, arg ...string) {
	s.Hash.List(m, arg...)
}
func (s plugs) Show(m *ice.Message, arg ...string) {
	m.ProcessFloat(m.Option(ctx.INDEX), kit.Split(m.Option(ctx.ARGS)), arg...)
	m.Option(ice.FIELD_PREFIX, ctx.RUN, m.Option(ctx.INDEX))
}

func init() { ice.CodeCtxCmd(plugs{}) }
