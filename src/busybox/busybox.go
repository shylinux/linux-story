package busybox

import (
	"shylinux.com/x/ice"
)

type project struct {
	ice.Code
	source string `data:"https://busybox.net/downloads/busybox-1.33.0.tar.bz2"`
	list   string `name:"list path auto order build download" help:"命令行"`
}

func (s project) Build(m *ice.Message, arg ...string) {
	s.Code.Build(m, "", func(p string) { s.Code.System(m, p, "make", "defconfig") })
}
func (s project) Order(m *ice.Message, arg ...string) {
	s.Code.Order(m.Spawn(), "", "_install/usr/sbin")
	s.Code.Order(m.Spawn(), "", "_install/usr/bin")
	s.Code.Order(m.Spawn(), "", "_install/sbin")
	s.Code.Order(m, "", "_install/bin")
}
func (s project) List(m *ice.Message, arg ...string) {
	s.Code.Source(m, "", arg...)
}

func init() { ice.CodeCtxCmd(project{}) }
