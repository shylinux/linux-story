package system

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/aaa"
)

type whois struct {
	ice.Hash
	short string `data:"ip"`
	field string `data:"time,ip,location"`
	list  string `name:"list ip auto"`
}

func (s whois) Create(m *ice.Message, arg ...string) {
	m.OptionDefault(aaa.LOCATION, func() string { return m.PublicIP(m.Option(aaa.IP)) })
	s.Hash.Create(m, m.OptionSimple("ip,location")...)
}
func (s whois) List(m *ice.Message, arg ...string) {
	s.Hash.List(m, arg...)
}

func init() { ice.CodeCtxCmd(whois{}) }
