module shylinux.com/x/linux-story

go 1.13

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

require (
	shylinux.com/x/ice v1.4.7
	shylinux.com/x/icebergs v1.8.7
	shylinux.com/x/toolkits v1.0.2
)
