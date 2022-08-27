package ffmpeg

import "shylinux.com/x/ice"

type runtime struct {
	ice.Code
	source string `data:"https://ffmpeg.org/releases/ffmpeg-4.2.1.tar.bz2"`
	list   string `name:"list path auto order build download" help:"音视频"`
}

func (s runtime) Build(m *ice.Message, arg ...string) {
	s.Code.Build(m, "")
}
func (s runtime) List(m *ice.Message, arg ...string) {
	s.Code.Source(m, "", arg...)
}

func init() { ice.CodeCtxCmd(runtime{}) }
