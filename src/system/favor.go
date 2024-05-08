package system

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web/html"
	kit "shylinux.com/x/toolkits"
)

type favor struct {
	ice.Hash
	export string `data:"true"`
	short  string `data:"path"`
	field  string `data:"time,path,zone,type,name,text"`
	zone   string `name:"zone zone"`
	list   string `name:"list path auto"`
}

func (s favor) Create(m *ice.Message, arg ...string) {
	ls := kit.Split(m.Option(nfs.PATH), nfs.PS)
	m.OptionDefault(mdb.TYPE, kit.Select("", ls, -2), mdb.NAME, kit.Select("", ls, -1))
	s.Hash.Create(m, m.OptionSimple("path,zone,type,name,text")...)
}
func (s favor) List(m *ice.Message, arg ...string) {
	if s.Hash.List(m, arg...); len(arg) == 0 || arg[0] == "" {
		m.PushAction(s.Zone, s.Remove).Action(s.Create, html.FILTER).Sort("zone,path")
		m.StatusTimeCountStats(mdb.ZONE, mdb.TYPE)
	} else {
		s.show(m, m.Append(mdb.TYPE), m.Append(nfs.PATH))
	}
}
func (s favor) Zone(m *ice.Message, arg ...string) {
	s.Hash.Modify(m, nfs.PATH, m.Option(nfs.PATH), mdb.ZONE, m.Option(mdb.ZONE))
	m.ProcessRefresh()
}

func init() { ice.CodeCtxCmd(favor{}) }

func (s favor) show(m *ice.Message, t, p string) *ice.Message {
	switch t {
	case "bin", "sbin":
		m.Echo(m.SystemCmdx(p, "--help"))
	case "etc":
		m.Cmdy(nfs.CAT, p)
	}
	return m
}
