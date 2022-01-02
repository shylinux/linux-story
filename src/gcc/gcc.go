package gcc

import (
	"path"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/nfs"
)

type Source struct {
	ice.Code
	source string `data:"http://mirrors.tencent.com/macports/distfiles/gcc48/gcc-4.8.4.tar.bz2"`
	list   string `name:"list path auto build download" help:"编译器"`
}

func (s Source) Download(m *ice.Message, arg ...string) {
	s.Code.Download(m, m.Config(nfs.SOURCE), arg...)
}
func (s Source) Build(m *ice.Message, arg ...string) {
	s.Code.Build(m, m.Config(nfs.SOURCE), "--enable-languages=c,c++", "--disable-multilib", "--disable-checking")
}
func (s Source) List(m *ice.Message, arg ...string) {
	m.Option(nfs.DIR_ROOT, path.Join(s.Code.Path(m, m.Config(nfs.SOURCE)), "_install"))
	m.Cmdy(nfs.CAT, arg)
}

func init() { ice.CodeCmd(Source{}) }
