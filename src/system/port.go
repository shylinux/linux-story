package system

import (
	"strconv"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

type port struct {
	list string `name:"list port auto"`
}

func (s port) List(m *ice.Message, arg ...string) {
	s.tcp4(m)
	s.tcp6(m)
	m.StatusTimeCountStats(mdb.TYPE, mdb.STATUS).Sort("status,port,local", []string{"LISTEN", "ESTABLISHED", "TIME_WAIT"}, ice.INT)
}

func init() { ice.CodeCtxCmd(port{}) }

func (s port) tcp4(m *ice.Message, arg ...string) {
	m.Spawn().Split(m.Cmdx(nfs.CAT, "/proc/net/tcp")).Table(func(value ice.Maps) {
		ls := kit.Split(value["local_address"], ":")
		m.Push(mdb.TYPE, "tcp4").Push(mdb.STATUS, s.trans(value["st"]))
		m.Push("port", s.parse(ls[1]))
		m.Push("local", kit.Format("%d.%d.%d.%d:%d", s.parse(ls[0][6:8]), s.parse(ls[0][4:6]), s.parse(ls[0][2:4]), s.parse(ls[0][:2]), s.parse(ls[1])))
		ls = kit.Split(value["rem_address"], ":")
		m.Push("remote", kit.Format("%d.%d.%d.%d:%d", s.parse(ls[0][6:8]), s.parse(ls[0][4:6]), s.parse(ls[0][2:4]), s.parse(ls[0][:2]), s.parse(ls[1])))
		m.Push("ip", kit.Format("%d.%d.%d.%d", s.parse(ls[0][6:8]), s.parse(ls[0][4:6]), s.parse(ls[0][2:4]), s.parse(ls[0][:2])))
	})
}
func (s port) tcp6(m *ice.Message, arg ...string) {
	m.Spawn().Split(m.Cmdx(nfs.CAT, "/proc/net/tcp6")).Table(func(value ice.Maps) {
		ls := kit.Split(value["local_address"], ":")
		m.Push(mdb.TYPE, "tcp6").Push(mdb.STATUS, s.trans(value["st"]))
		m.Push("port", s.parse(ls[1]))
		m.Push("local", kit.Format("%d.%d.%d.%d:%d", s.parse(ls[0][30:32]), s.parse(ls[0][28:30]), s.parse(ls[0][26:28]), s.parse(ls[0][24:26]), s.parse(ls[1])))
		ls = kit.Split(value["remote_address"], ":")
		m.Push("remote", kit.Format("%d.%d.%d.%d:%d", s.parse(ls[0][30:32]), s.parse(ls[0][28:30]), s.parse(ls[0][26:28]), s.parse(ls[0][24:26]), s.parse(ls[1])))
		m.Push("ip", kit.Format("%d.%d.%d.%d", s.parse(ls[0][30:32]), s.parse(ls[0][28:30]), s.parse(ls[0][26:28]), s.parse(ls[0][24:26])))
	})
}
func (s port) parse(str string) int64 {
	port, _ := strconv.ParseInt(str, 16, 32)
	return port
}
func (s port) trans(str string) string {
	return kit.Select(str, map[string]string{
		"01": "ESTABLISHED",
		"02": "TCP_SYNC_SEND",
		"03": "TCP_SYNC_RECV",
		"04": "TCP_FIN_WAIT1",
		"05": "TCP_FIN_WAIT2",
		"06": "TIME_WAIT",
		"07": "TCP_CLOSE",
		"08": "TCP_CLOSE_WAIT",
		"0A": "LISTEN",
	}[str])
}
