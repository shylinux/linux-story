package system

import (
	"path"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/icebergs/base/web/html"
	kit "shylinux.com/x/toolkits"
)

type dir struct {
	list string `name:"list path auto"`
}

func (s dir) Trash(m *ice.Message, arg ...string) {
	m.Trash(m.Option(nfs.PATH))
}
func (s dir) Upload(m *ice.Message, arg ...string) {
	if strings.HasSuffix(m.Option(nfs.PATH), nfs.PS) {
		m.UploadSave(m.Option(nfs.PATH))
	} else {
		m.UploadSave(path.Dir(m.Option(nfs.PATH)))
	}
}
func (s dir) List(m *ice.Message, arg ...string) {
	if len(arg) > 0 && !strings.HasSuffix(arg[0], nfs.PS) {
		if nfs.IsSourceFile(m.Message, kit.Ext(arg[0])) {
			s.cmds(m, web.CODE_VIMER, path.Dir(arg[0]), path.Base(arg[0]))
		} else {
			s.cmds(m, ice.GetTypeKey(hex{}), arg...)
			s.cmds(m, web.CODE_XTERM, mdb.TYPE, cli.SH, nfs.PATH, arg[0])
		}
	} else {
		m.Cmdy(nfs.DIR, kit.Select(nfs.PS, arg, 0))
		m.PushAction(s.Upload, s.Trash)
		m.SortStr(nfs.PATH)
	}
}

func init() { ice.CodeCtxCmd(dir{}) }

func (s dir) cmds(m *ice.Message, cmd string, arg ...string) {
	m.Cmdy(ctx.COMMAND, cmd)
	m.Push(ctx.ARGS, kit.Format(arg))
	m.Push(ctx.STYLE, html.OUTPUT)
}
