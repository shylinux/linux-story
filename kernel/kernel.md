# {{title "kernel"}}
{{brief "简介" `linux是一种简单高效的操作系统`}}
{{refer "官方网站" `
官网 https://kernel.org
`}}

## {{chapter "下载源码"}}
{{shell "下载源码" "usr" "install" `wget https://mirror.bjtu.edu.cn/kernel/linux/kernel/v3.x/linux-3.10.tar.gz`}}
{{shell "解压源码" "usr" "install" `tar xvf linux-3.10.tar.gz`}}

## {{chapter "项目结构"}}

{{shell "项目结构" "usr" `dir linux-3.10`}}

{{shell "生成索引" "usr" "install" `cd linux-3.10 && make tags`}}

## {{chapter "启动流程"}}
## {{chapter "设备驱动"}}
## {{chapter "文件系统"}}
## {{chapter "启动流程"}}
## {{chapter "启动流程"}}
## {{chapter "启动流程"}}
{{stack "网络连接" `
inet_init()
`}}

## {{chapter "系统调用"}}

## {{chapter "epoll"}}

{{shell "接口列表" "usr/linux-3.10" `sed -n "/epoll/p" arch/x86/syscalls/syscall_64.tbl`}}
{{shell "接口定义" "usr/linux-3.10" `sed -n "/epoll/p" include/uapi/asm-generic/unistd.h`}}
{{shell "系统调用" "usr/linux-3.10" `sed -n "/epoll/p" include/linux/syscalls.h`}}

{{shell "结构定义" "usr/linux-3.10" `sed -n "/struct eventpoll {/,/^};/p" fs/eventpoll.c`}}
{{shell "创建轮询" "usr/linux-3.10" `sed -n "/SYSCALL_DEFINE1(epoll_create1/,/^}/p" fs/eventpoll.c`}}
{{shell "添加文件" "usr/linux-3.10" `sed -n "/SYSCALL_DEFINE4(epoll_ctl/,/^}/p" fs/eventpoll.c`}}
{{shell "阻塞等待" "usr/linux-3.10" `sed -n "/SYSCALL_DEFINE4(epoll_wait/,/^}/p" fs/eventpoll.c`}}

