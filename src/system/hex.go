package system

import (
	"debug/macho"
	"encoding/base64"
	"os"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web/html"
	kit "shylinux.com/x/toolkits"
)

type hex struct {
	list string `name:"list path auto"`
}

func (s hex) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 || strings.HasSuffix(arg[0], nfs.PS) {
		m.Cmdy(nfs.DIR, arg)
	} else if html.IsImage(arg[0], "") {
		m.Echo(`<img src="data:image/%s;base64,%s" title='%s' />`, kit.Ext(arg[0]), base64.StdEncoding.EncodeToString([]byte(m.Cmdx(nfs.CAT, arg[0]))), arg[0])
	} else if nfs.IsSourceFile(m.Message, kit.Ext(arg[0])) {
		m.Cmdy(nfs.CAT, arg[0])
	} else if f, e := os.Open(arg[0]); !m.Warn(e) {
		defer f.Close()
		if o, e := macho.NewFile(f); e == nil {
			for _, v := range o.Sections {
				m.Push("type", v.Seg)
				m.Push("name", v.Name)
				m.Push("addr", kit.Format("%0#X", v.Addr))
				m.Push("offset", kit.Format("%0#X", v.Offset))
				m.Push("size", kit.FmtSize(v.Size))
			}
			return
		}
		buf := make([]byte, 128)
		n, _ := f.Read(buf)
		for i := 0; i < n; i++ {
			kit.If(i%8 == 0, func() { m.Push("byte", i) })
			m.Push(kit.Format(i%8), kit.Format("%02X", buf[i]))
		}
	}
}

func init() { ice.CodeCtxCmd(hex{}) }
