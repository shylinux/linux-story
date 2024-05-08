package system

import (
	"shylinux.com/x/ice"
	kit "shylinux.com/x/toolkits"
)

type yum struct {
	list string `name:"list name auto"`
}

func (s yum) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 || arg[0] == "" {
		list := []string{}
		kit.For(kit.SplitLine(m.SystemCmdx("yum", "list", "installed"), ""), func(text string, index int) {
			if kit.HasPrefix(text, "Repodata is over", "Installed Packages", "Loaded plugins:") {
				return
			}
			if list = append(list, kit.Split(text)...); len(list) < 3 {
				return
			}
			m.Push("name", list[0]).Push("version", list[1]).Push("info", list[2])
			list = list[:0]
		})
		m.ActionFilter()
	} else {
		m.Echo(m.SystemCmdx("yum", "info", arg[0]))
	}
}

func init() { ice.CodeCtxCmd(yum{}) }
