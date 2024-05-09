package system

import (
	"strconv"

	"shylinux.com/x/ice"
	kit "shylinux.com/x/toolkits"
)

type unicode struct {
	ice.Hash
	short  string `data:"begin"`
	field  string `data:"time,begin,end"`
	vendor string `data:"https://symbl.cc/cn/unicode-table/"`
	list   string `name:"list begin auto"`
}

func (s unicode) List(m *ice.Message, arg ...string) {
	s.Hash.List(m.Spawn(), arg...).Sort("begin").Table(func(value ice.Maps) {
		begin, _ := strconv.ParseInt(value["begin"], 16, 64)
		end, _ := strconv.ParseInt(value["end"], 16, 64)
		for i := begin; i <= end; i++ {
			m.Push(kit.Format("%X", i%16), kit.Format("<span style='font-size:32px;'>%s<span>", string(i)))
			m.Push(kit.Format("%X", i%16), kit.Format("%X", i))
		}
	})
	m.Action(s.Create, s.Vendor)
}

func init() { ice.CodeCtxCmd(unicode{}) }
