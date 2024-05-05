package system

import (
	"strconv"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

type port struct {
	list string `name:"list hash auto"`
}

func (s port) List(m *ice.Message, arg ...string) {
	parse := func(str string) int64 { port, _ := strconv.ParseInt(str, 16, 32); return port }
	trans := func(str string) string {
		switch str {
		case "01":
			return "ESTABLISHED"
		case "02":
			return "TCP_SYNC_SEND"
		case "03":
			return "TCP_SYNC_RECV"
		case "04":
			return "TCP_FIN_WAIT1"
		case "05":
			return "TCP_FIN_WAIT2"
		case "06":
			return "TIME_WAIT"
		case "07":
			return "TCP_CLOSE"
		case "08":
			return "TCP_CLOSE_WAIT"
		case "0A":
			return "LISTEN"
		default:
			return str
		}
	}
	stats := map[string]int{}
	m.Spawn().Split(m.Cmdx(nfs.CAT, "/proc/net/tcp")).Table(func(value ice.Maps) {
		stats[trans(value["st"])]++
		m.Push(mdb.STATUS, trans(value["st"]))
		ls := kit.Split(value["local_address"], ":")
		m.Push("local", kit.Format("%d.%d.%d.%d:%d", parse(ls[0][6:8]), parse(ls[0][4:6]), parse(ls[0][2:4]), parse(ls[0][:2]), parse(ls[1])))
		ls = kit.Split(value["rem_address"], ":")
		m.Push("remote", kit.Format("%d.%d.%d.%d:%d", parse(ls[0][6:8]), parse(ls[0][4:6]), parse(ls[0][2:4]), parse(ls[0][:2]), parse(ls[1])))
	})
	m.Spawn().Split(m.Cmdx(nfs.CAT, "/proc/net/tcp6")).Table(func(value ice.Maps) {
		stats[trans(value["st"])]++
		m.Push(mdb.STATUS, trans(value["st"]))
		ls := kit.Split(value["local_address"], ":")
		m.Push("local", kit.Format("%d.%d.%d.%d:%d", parse(ls[0][30:32]), parse(ls[0][28:30]), parse(ls[0][26:28]), parse(ls[0][24:26]), parse(ls[1])))
		ls = kit.Split(value["remote_address"], ":")
		m.Push("remote", kit.Format("%d.%d.%d.%d:%d", parse(ls[0][30:32]), parse(ls[0][28:30]), parse(ls[0][26:28]), parse(ls[0][24:26]), parse(ls[1])))
	})
	m.Sort("status,local", []string{"LISTEN", "ESTABLISHED", "TIME_WAIT"}).StatusTimeCount(stats)
}

func init() { ice.CodeCtxCmd(port{}) }
