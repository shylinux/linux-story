#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <stdio.h>

int main() {
	int s = socket(AF_INET, SOCK_STREAM, 0);

	struct sockaddr_in addr;
	addr.sin_family = AF_INET;
	addr.sin_addr.s_addr = inet_addr("127.0.0.1");
	addr.sin_port = htons(9374);
	bind(s, (struct sockaddr*)&addr, (socklen_t)sizeof(addr));

	listen(s, 10);
	printf("listen %d\n", s);

	char buf[1024];
	while (1) {
		int c = accept(s, NULL, NULL);
		printf("accept %d\n", c);
		int n = read(c, buf, sizeof(buf));
		printf("read %s\n", buf);
		write(c, buf, n);
		close(c);
		printf("close %d\n", c);
	}
	close(s);
}
