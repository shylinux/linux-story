package busybox

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/linux-story/src/gcc"
)

type Source struct {
	gcc.Source
	source string `data:"http://mirrors.tencent.com/tinycorelinux/13.x/aarch64/releases/RPi/src/busybox/busybox-1.33.0.tar.bz2"`
	list   string `name:"list path auto build download" help:"命令行"`
}

func (s Source) Build(m *ice.Message, arg ...string) {
	s.Code.Prepare(m, func(p string) { s.Code.System(m, p, "make", "defconfig") })
	s.Code.Build(m, m.Config(nfs.SOURCE), arg...)
}

func init() { ice.CodeCmd(Source{}) }
