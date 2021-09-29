module shylinux.com/x/linux-story

go 1.11

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

require (
	shylinux.com/x/ice v0.1.9
	shylinux.com/x/icebergs v0.4.7
	shylinux.com/x/toolkits v0.3.2
)
