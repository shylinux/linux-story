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
	list string `name:"list text auto" help:"框架图"`
}

func (s chain) processInner(m *ice.Message, arg ...ice.Any) {
	m.ProcessStory(append(kit.List(m.Prefix(code.INNER)), arg...)...)
}
func (s chain) Find(m *ice.Message, arg ...string) {
	if nfs.ExistsFile(m, path.Join(m.Option(nfs.PATH), arg[0])) {
		s.processInner(m, m.Option(nfs.PATH), arg[0], "1")
		return
	}

	if !nfs.ExistsFile(m, path.Join(m.Option(nfs.PATH), nfs.TAGS)) {
		s.System(m.Spawn(), m.Option(nfs.PATH), "ctags", "-R")
	}

	ls := kit.Split(arg[0], "", ":=/()")
	var meta = struct{ name, kind, value, sub, arg string }{name: ls[0]}
	for i := 1; i < len(ls); i += 2 {
		switch ls[i] {
		case ":":
			meta.kind = ls[i+1]
		case "=":
			meta.value = ls[i+1]
		case "/":
			meta.sub = ls[i+1]
		case "(":
			meta.arg = ls[i+1]
		}
	}
	if msg := m.Cmd(code.INNER, nfs.TAGS, kit.Select(meta.name, meta.kind, meta.value, meta.sub)); msg.Append(nfs.FILE) != "" {
		s.processInner(m, msg.Append(nfs.PATH), msg.Append(nfs.FILE), msg.Append(nfs.LINE))
		return
	}

	if msg := s.System(m.Spawn(), "", cli.MAN, "3", arg[0]); cli.IsSuccess(msg) && !strings.HasPrefix(msg.Result(), "No manual entry for") {
		s.processInner(m, cli.MAN, arg[0], 3)
		return
	}
	if msg := s.System(m.Spawn(), "", cli.MAN, arg[0]); cli.IsSuccess(msg) && !strings.HasPrefix(msg.Result(), "No manual entry for") {
		s.processInner(m, cli.MAN, arg[0])
		return
	}
}
func (s chain) List(m *ice.Message, arg ...string) {
	args := []string{nfs.PATH, m.Option("ctags.path")}
	for _, v := range arg {
		if v == nfs.PATH {
			args = nil
		}
	}
	m.Cmdy(wiki.CHART, wiki.CHAIN, arg, args, kit.Dict(ctx.INDEX, m.PrefixKey()))
}

func init() { ice.CodeCtxCmd(chain{}) }
