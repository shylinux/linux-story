package system

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/aaa"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/tcp"
	kit "shylinux.com/x/toolkits"
)

type port struct {
	whois whois
	list  string `name:"list port auto"`
}

func (s port) List(m *ice.Message, arg ...string) {
	kit.For(kit.SplitLine(m.SystemCmdx(SUDO, "netstat", "-antp")), func(text string, index int) {
		if index < 2 {
			return
		}
		ls := kit.SplitWord(text)
		m.Push(mdb.TYPE, ls[0]).Push(mdb.STATUS, ls[5])
		m.Push(tcp.PORT, kit.Select("", kit.Split(ls[3], ":"), -1))
		m.Push("local", ls[3]).Push("remote", ls[4])
		m.Push(aaa.IP, kit.Select("::", kit.Split(ls[4], ":*"), 0))
		_ls := kit.Split(ls[6], "/:")
		m.Push(cli.PID, kit.Select("", _ls, 0))
		m.Push(cli.CMD, kit.Select("", _ls, 1))
	})
	s.getwhois(m).StatusTimeCountStats(mdb.TYPE, mdb.STATUS, cli.CMD)
	m.Sort("status,location,port,local", []string{"LISTEN", "ESTABLISHED", "TIME_WAIT"}, ice.STR_R, ice.INT)
}

func init() { ice.CodeCtxCmd(port{}) }

func (s port) getwhois(m *ice.Message) *ice.Message {
	list := map[string]string{}
	m.Table(func(value ice.Maps) {
		p, ok := list[value[aaa.IP]]
		if !ok && !kit.HasPrefix(value[aaa.IP], "::", "0.0.", "127.0.") {
			if p = m.Cmd(s.whois, value[aaa.IP]).Append(aaa.LOCATION); p == "" {
				p = m.PublicIP(value[aaa.IP])
				m.Cmd(s.whois, s.whois.Create, aaa.IP, value[aaa.IP], aaa.LOCATION, p)
			}
			list[value[aaa.IP]] = p
		}
		m.Push(aaa.LOCATION, p)
	})
	return m
}
