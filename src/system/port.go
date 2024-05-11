package system

import (
	"runtime"

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
	args := []string{"-antp"}
	list := map[string]ice.Maps{}
	if runtime.GOOS == cli.DARWIN {
		list = m.CmdMap(proc{}, "PID")
		args = []string{"-an", "-p", "tcp", "-v"}
	}
	kit.For(kit.SplitLine(m.SystemCmdx(SUDO, "netstat", args)), func(text string, index int) {
		if index < 2 {
			return
		}
		ls := kit.SplitWord(text)
		m.Push(mdb.TYPE, ls[0]).Push(mdb.STATUS, ls[5])
		m.Push(tcp.PORT, kit.Select("", kit.Split(ls[3], ".:"), -1))
		m.Push("local", ls[3]).Push("remote", ls[4])
		if ls[0] == "tcp6" || ls[4] == "*.*" {
			m.Push(aaa.IP, "")
		} else {
			if runtime.GOOS == cli.DARWIN {
				m.Push(aaa.IP, kit.Join(kit.Slice(kit.Split(ls[4], "."), 0, 4), "."))
			} else {
				m.Push(aaa.IP, kit.Select("::", kit.Split(ls[4], ":*"), 0))
			}
		}

		if runtime.GOOS == cli.DARWIN {
			m.Push(cli.PID, ls[8])
			m.Push(cli.CMD, list[ls[8]][ice.CMD])
		} else if len(ls) > 6 {
			_ls := kit.Split(ls[6], "/:")
			m.Push(cli.PID, kit.Select("", _ls, 0))
			m.Push(cli.CMD, kit.Select("", _ls, 1))
		}
	})
	s.getwhois(m).StatusTimeCountStats(mdb.TYPE, mdb.STATUS, cli.CMD, aaa.LOCATION)
	m.Sort("status,location,port,local", []string{"LISTEN", "ESTABLISHED", "TIME_WAIT"}, ice.STR_R, ice.INT)
}

func init() { ice.CodeCtxCmd(port{}) }

func (s port) getwhois(m *ice.Message) *ice.Message {
	list := map[string]string{}
	m.Table(func(value ice.Maps) {
		p, ok := list[value[aaa.IP]]
		if !ok && value[aaa.IP] != "" && !kit.HasPrefix(value[aaa.IP], "0.0.", "127.0.", "192.168.") && value[mdb.TYPE] != "tcp6" {
			if p = m.Cmd(s.whois, value[aaa.IP]).Append(aaa.LOCATION); p == "" {
				p = m.PublicIP(value[aaa.IP])
				m.Cmd(s.whois, s.whois.Create, aaa.IP, value[aaa.IP], aaa.LOCATION, p, ice.CMD, value[ice.CMD])
			}
			list[value[aaa.IP]] = p
		}
		m.Push(aaa.LOCATION, p)
	})
	return m
}
