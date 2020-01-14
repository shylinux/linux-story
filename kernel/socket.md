## {{chapter "socket"}}

{{shell "接口列表" "usr/linux-3.10" `sed -n "/socket/p" arch/x86/syscalls/syscall_64.tbl`}}
{{shell "接口定义" "usr/linux-3.10" `sed -n "/socket/p" include/uapi/asm-generic/unistd.h`}}
{{shell "系统调用" "usr/linux-3.10" `sed -n "/socket/p" include/linux/syscalls.h`}}

{{shell "连接状态" "usr/linux-3.10" `sed -n "/enum {/,/^};/p" include/net/tcp_states.h`}}

{{shell "报文格式" "usr/linux-3.10" `sed -n "/struct iphdr {/,/^};/p" include/uapi/linux/ip.h`}}
{{shell "报文格式" "usr/linux-3.10" `sed -n "/struct tcphdr {/,/^};/p" include/uapi/linux/tcp.h`}}

```
net/ipv4/af_inet.c
static struct inet_protosw inetsw_array[] =

net/ipv4/tcp_ipv4.c
struct proto tcp_prot = {
```

{{stack "网络连接" `
socket
    socket(PF_INET,SOCKET_STREAM):net/socket.c:1360
        sock_create()
            __sock_create()
                sock=sock_alloc()
					sock->sk_prot=
                sock->type=SOCKET_STREAM
                pf=rcu_dereference(net_families[sock->type])/inet_family_ops
                pf->create()/inet_create()
                    sock->ops=inetsw[sock->type]/inet_stream_ops
    connect()
        sock->ops->connect()/inet_stream_connect()
            __inet_stream_connect()
                sk->sk_prot->connect()/tcp_v4_connect()
            tcp_set_state(TCP_SYN_SENT)
            tcp_connect()
    tcp_v4_rcv()

`}}

