package alpine

import (
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

type alpine struct {
	mirrors string `name:"mirrors from=dl-cdn.alpinelinux.org to=mirrors.tencent.com" help:"镜像"`
	tzone   string `name:"mirrors tzone=Asia/Shanghai" help:"镜像"`
	list    string `name:"list auto mirrors" help:"alpine"`
}

func (s alpine) Mirrors(m *ice.Message, arg ...string) {
	if !cli.IsAlpine(m.Message) {
		return
	}
	kit.Rewrite("/etc/apk/repositories", func(text string) string {
		return strings.Replace(text, m.Option("from"), m.Option("to"), -1)
	})
	m.Cmdy(nfs.CAT, "/etc/apk/repositories")
}
func (s alpine) Tzone(m *ice.Message, arg ...string) {
	if !cli.IsAlpine(m.Message) {
		return
	}
	m.Cmd(cli.SYSTEM, "apk", "add", "tzdata")
	m.Cmd(cli.SYSTEM, "cp", "/usr/share/zoneinfo/"+m.Option("tzone"), "/etc/localtime")
	m.Cmd(nfs.SAVE, "/etc/timezone", m.Option("tzone"))
	m.Cmdy(nfs.CAT, "/etc/timezone")
}
func (s alpine) List(m *ice.Message, arg ...string) {
}

func init() { ice.CodeCmd(alpine{}) }
