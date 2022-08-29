package ctags

import (
	"bufio"
	"path"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/core/code"
	"shylinux.com/x/icebergs/core/wiki"
	kit "shylinux.com/x/toolkits"
)

type chain struct {
}

func (s chain) Find(m *ice.Message, arg ...string) {
	if !nfs.ExistsFile(m, path.Join(m.Option(nfs.PATH), "tags")) {
		m.Cmd(cli.SYSTEM, "ctags", "-R", kit.Dict(cli.CMD_DIR, m.Option(nfs.PATH)))
	}
	for _, l := range strings.Split(m.Cmdx(cli.SYSTEM, "grep", "^"+arg[0]+"\\>", "tags", kit.Dict(cli.CMD_DIR, m.Option(nfs.PATH))), ice.NL) {
		ls := strings.SplitN(l, ice.TB, 2)
		if len(ls) < 2 {
			continue
		}

		ls = strings.SplitN(ls[1], ice.TB, 2)
		file := ls[0]
		ls = strings.SplitN(ls[1], ";\"", 2)
		text := strings.TrimSuffix(strings.TrimPrefix(ls[0], "/^"), "$/")
		line := kit.Int(text)

		f, e := nfs.OpenFile(m, path.Join(m.Option(nfs.PATH), file))
		m.Assert(e)
		defer f.Close()

		bio := bufio.NewScanner(f)
		for i := 1; bio.Scan(); i++ {
			if i == line || bio.Text() == text {
				m.ProcessStory(code.INNER, m.Option(nfs.PATH), strings.TrimPrefix(file, nfs.PWD), kit.Format(i))
				return
			}
		}
	}
	m.Option(cli.CMD_ENV, "COLUMNS", "100")
	if msg := m.Cmd(cli.SYSTEM, "sh", "-c", kit.Format("man %s|col -b", arg[0])); !strings.HasPrefix(msg.Result(), "No manual entry for") {
		m.ProcessStory(code.INNER, "man", arg[0])
		m.Option(mdb.TEXT, msg.Result())
		return
	}

	m.ProcessStory(code.INNER, m.Option(nfs.PATH), arg[0], "1")
}
func (s chain) List(m *ice.Message, arg ...string) {
	m.Cmdy(wiki.CHART, wiki.CHAIN, arg, kit.Dict(ctx.INDEX, m.PrefixKey()))
}

func init() { ice.CodeCtxCmd(chain{}) }
