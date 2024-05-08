package system

import (
	"encoding/base64"
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

func (s dir) Upload(m *ice.Message, arg ...string) {
	if strings.HasSuffix(m.OptionDefault(nfs.PATH, nfs.PS), nfs.PS) {
		m.UploadSave(m.Option(nfs.PATH))
	} else {
		m.UploadSave(path.Dir(m.Option(nfs.PATH)))
	}
}
func (s dir) Trash(m *ice.Message, arg ...string) {
	m.Trash(m.Option(nfs.PATH))
}
func (s dir) Show(m *ice.Message, arg ...string) {
	m.ProcessFloat(hex{}, m.Option(nfs.PATH), arg...)
}
func (s dir) List(m *ice.Message, arg ...string) {
	if len(arg) > 0 && arg[0] == "/dev/" {
		return
	} else if len(arg) == 0 || strings.HasSuffix(arg[0], nfs.PS) {
		m.Cmdy(nfs.DIR, kit.Select(nfs.PS, arg, 0), "time,path,type,size").PushAction(s.Show, s.Trash).Action(s.Upload).Sort(nfs.PATH)
	} else {
		if html.IsImage(arg[0], "") {
			m.Echo(`<img src="data:image/%s;base64,%s" title='%s' />`, kit.Ext(arg[0]), base64.StdEncoding.EncodeToString([]byte(m.Cmdx(nfs.CAT, arg[0]))), arg[0])
		} else if nfs.IsSourceFile(m.Message, kit.Ext(arg[0])) || kit.HasPrefix(arg[0], "/etc/", "/proc/") {
			s.cmds(m, web.CODE_VIMER, path.Dir(arg[0]), path.Base(arg[0]))
		} else {
			s.cmds(m, hex{}, arg...)
			if kit.HasPrefix(arg[0],
				"/bin/", "/sbin/",
				"/usr/bin/", "/usr/sbin/",
				"/usr/local/bin", "/usr/local/sbin",
			) {
				s.cmds(m, web.CODE_XTERM, mdb.TYPE, cli.SH, mdb.TEXT, kit.JoinWord(arg[0], "--help"), nfs.PATH, arg[0])
			} else {
				s.cmds(m, web.CODE_XTERM, mdb.TYPE, cli.SH, nfs.PATH, arg[0])
			}
		}
	}
}

func init() { ice.CodeCtxCmd(dir{}) }

func (s dir) cmds(m *ice.Message, cmd ice.Any, arg ...string) {
	m.Cmdy(ctx.COMMAND, ice.GetTypeKey(cmd)).Push(ctx.ARGS, kit.Format(arg)).Push(ctx.STYLE, html.OUTPUT)
}
