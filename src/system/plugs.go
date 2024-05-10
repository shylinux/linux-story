package system

import "shylinux.com/x/ice"

type plugs struct {
	ice.Hash
	short string `data:"path"`
	field string `data:"time,path,index,args"`
	list  string `name:"list hash auto"`
}

func (s plugs) List(m *ice.Message, arg ...string) {
	s.Hash.List(m, arg...)
}

func init() { ice.CodeCtxCmd(plugs{}) }
