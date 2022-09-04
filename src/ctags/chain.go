package ctags

import (
	"path"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/core/code"
	"shylinux.com/x/icebergs/core/wiki"
	kit "shylinux.com/x/toolkits"
)

type chain struct{}

func (s chain) Find(m *ice.Message, arg ...string) {
	if !nfs.ExistsFile(m, path.Join(m.Option(nfs.PATH), "tags")) {
		m.Cmd(cli.SYSTEM, "ctags", "-R", kit.Dict(cli.CMD_DIR, m.Option(nfs.PATH)))
	}

	if msg := m.Cmd("web.code.inner", "tags", arg[0]); msg.Append(nfs.FILE) != "" {
		m.ProcessStory(code.INNER, msg.Append(nfs.PATH), msg.Append(nfs.FILE), msg.Append(nfs.LINE))
		return
	}

	m.Option(cli.CMD_ENV, "COLUMNS", "100")
	if msg := m.Cmd(cli.SYSTEM, "sh", "-c", kit.Format("man %s|col -b", arg[0])); cli.IsSuccess(msg) && !strings.HasPrefix(msg.Result(), "No manual entry for") {
		m.ProcessStory(code.INNER, "man", arg[0])
		return
	}

	if nfs.ExistsFile(m, path.Join(m.Option(nfs.PATH), arg[0])) {
		m.ProcessStory(code.INNER, m.Option(nfs.PATH), arg[0], "1")
		return
	}
}
func (s chain) List(m *ice.Message, arg ...string) {
	m.Cmdy(wiki.CHART, wiki.CHAIN, arg, kit.Dict(ctx.INDEX, m.PrefixKey()))
}

func init() { ice.CodeCtxCmd(chain{}) }
