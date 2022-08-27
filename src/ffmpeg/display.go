package ffmpeg

import "shylinux.com/x/ice"

type display struct {
	ice.Code
	list string `name:"list path auto order build download" help:"音视频"`
}

func (s display) List(m *ice.Message, arg ...string) {
	s.Code.Source(m, "", arg...)
}

func init() { ice.CodeCtxCmd(display{}) }
