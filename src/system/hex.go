package system

import (
	"os"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

type hex struct {
	list string `name:"list path auto"`
}

func (s hex) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 || strings.HasSuffix(arg[0], nfs.PS) {
		m.Cmdy(nfs.DIR, arg)
	} else if nfs.IsSourceFile(m.Message, kit.Ext(arg[0])) {
		m.Cmdy(nfs.CAT, arg[0])
	} else if f, e := os.Open(arg[0]); !m.Warn(e) {
		buf := make([]byte, 128)
		n, _ := f.Read(buf)
		for i := 0; i < n; i++ {
			kit.If(i%8 == 0, func() { m.Push("byte", i) })
			m.Push(kit.Format(i%8), kit.Format("%02X", buf[i]))
		}
	}
}

func init() { ice.CodeCtxCmd(hex{}) }
