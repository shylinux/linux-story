package system

import (
	"shylinux.com/x/ice"
)

type tools struct {
	ice.Hash
	export string `data:"true"`
	short  string `data:"index,args"`
	field  string `data:"time,enable,order,index,args,style,nick,icon,hash"`
	create string `name:"create index* args icon nick"`
	list   string `name:"list list" icon:"bi bi-grid"`
}

func (s tools) List(m *ice.Message, arg ...string) {
	s.Hash.List(m, arg...).Sort("enable,order,index,args", []string{ice.TRUE, ice.FALSE, ""}, ice.INT)
}

func init() { ice.CodeCtxCmd(tools{}) }
