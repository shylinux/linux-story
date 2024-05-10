package system

import "shylinux.com/x/ice"

type tools struct {
	ice.Hash
	short string `data:"index,args"`
	field string `data:"time,hash,index,args,style,icon,nick"`
	list  string `name:"list hash auto"`
}

func (s tools) List(m *ice.Message, arg ...string) {
	s.Hash.List(m, arg...)
}

func init() { ice.CodeCtxCmd(tools{}) }
