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

type chain struct {
	ice.Code
	list string `name:"list path file line auto" help:"编译器"`
}

func (s chain) Find(m *ice.Message, arg ...string) {
	if !nfs.ExistsFile(m, path.Join(m.Option(nfs.PATH), nfs.TAGS)) {
		m.Cmd(cli.SYSTEM, "ctags", "-R", kit.Dict(cli.CMD_DIR, m.Option(nfs.PATH)))
	}

	if msg := m.Cmd(code.INNER, nfs.TAGS, arg[0]); msg.Append(nfs.FILE) != "" {
		m.ProcessStory(m.PrefixKey(), msg.Append(nfs.PATH), msg.Append(nfs.FILE), msg.Append(nfs.LINE))
		return
	}

	if msg := m.Cmd(cli.SYSTEM, "sh", "-c", kit.Format("man 3 %s|col -b", arg[0])); cli.IsSuccess(msg) && !strings.HasPrefix(msg.Result(), "No manual entry for") {
		m.ProcessStory(m.PrefixKey(), "man", arg[0], 3)
		return
	}
	if msg := m.Cmd(cli.SYSTEM, "sh", "-c", kit.Format("man %s|col -b", arg[0])); cli.IsSuccess(msg) && !strings.HasPrefix(msg.Result(), "No manual entry for") {
		m.ProcessStory(m.PrefixKey(), "man", arg[0])
		return
	}

	if nfs.ExistsFile(m, path.Join(m.Option(nfs.PATH), arg[0])) {
		m.ProcessStory(m.PrefixKey(), m.Option(nfs.PATH), arg[0], "1")
		return
	}
}
func (s chain) Tags(m *ice.Message, arg ...string) {
	if msg := m.Cmd(code.INNER, nfs.TAGS, arg[0]); msg.Append(nfs.FILE) != "" {
		m.Copy(msg.Message)
		return
	}
	if msg := m.Cmd(cli.SYSTEM, "sh", "-c", kit.Format("man 3 %s|col -b", arg[0])); cli.IsSuccess(msg) && !strings.HasPrefix(msg.Result(), "No manual entry for") {
		m.Push("", kit.Dict(nfs.PATH, "man", nfs.FILE, arg[0], nfs.LINE, 3))
		return
	}
}
func (s chain) Man(m *ice.Message, arg ...string) {
	m.Option(cli.CMD_ENV, "COLUMNS", kit.Int(kit.Select("1920", m.Option("width")))/12)
	if len(arg) > 1 && arg[1] == "1" {
		arg[1] = ""
	}
	m.Cmdy(cli.SYSTEM, "sh", "-c", kit.Format("man %s %s|col -b", kit.Select("", arg, 1), arg[0]))
	m.Display("/plugin/local/code/inner.js")
}
func (s chain) List(m *ice.Message, arg ...string) {
	if strings.HasSuffix(arg[0], ice.PS) && !strings.Contains(arg[0], ice.NL) {
		m.Cmdy(code.INNER, arg)
		m.Display("/plugin/local/code/inner.js")
		return
	}
	m.Cmdy(wiki.CHART, wiki.CHAIN, arg, kit.Dict(ctx.INDEX, m.PrefixKey()))
}

func init() { ice.CodeCtxCmd(chain{}) }
