#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <string.h>
#include <stdio.h>

int main() {
	int s = socket(AF_INET, SOCK_STREAM, 0);

	struct sockaddr_in addr;
	addr.sin_family = AF_INET;
	addr.sin_addr.s_addr = inet_addr("127.0.0.1");
	addr.sin_port = htons(9374);
	connect(s, (struct sockaddr*)&addr, (socklen_t)sizeof(addr));

	char buf[1024] = "hello world";
	char res[1024] = "";
	write(s, buf, strlen(buf)+1);
	read(s, res, sizeof(res));
	puts(res);
	close(s);
}

