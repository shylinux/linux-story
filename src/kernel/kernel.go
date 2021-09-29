package kernel

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/linux-story/src/gcc"
)

type Source struct {
	gcc.Source
	source string `data:"http://mirrors.tencent.com/osmc/osmc/download/buildroot/2014.05/linux-3.12.20.tar.xz"`
	list   string `name:"list path auto build download" help:"源代码"`
}

func (s Source) Build(m *ice.Message, arg ...string) {
	s.Code.Prepare(m, func(p string) { s.Code.System(m, p, "make", "defconfig") })
	s.Code.Build(m, m.Config(cli.SOURCE))
}

func init() { ice.CodeCmd(Source{}) }
