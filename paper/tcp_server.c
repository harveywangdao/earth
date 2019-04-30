#include <stdio.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <ctype.h>
#include <string.h>

int port =8000;

int main(int argc, char** argv)
{
  struct sockaddr_in sin;
  struct sockaddr_in pin;
  int sock_descriptor;
  int temp_sock_descriptor;
  int address_size;
  char buf[16384];

  int i, len;

  /*
  *int socket(int domain, int type, int protocol);
  * PF_INET, AF_INET： Ipv4网络协议
  * PF_INET6, AF_INET6： Ipv6网络协议。
  * type参数的作用是设置通信的协议类型，可能的取值如下所示：
  　　      SOCK_STREAM： 提供面向连接的稳定数据传输，即TCP协议。
  　　      OOB： 在所有数据传送前必须使用connect()来建立连接状态。
  　      　SOCK_DGRAM： 使用不连续不可靠的数据包连接。
  　　      SOCK_SEQPACKET： 提供连续可靠的数据包连接。
  　      　SOCK_RAW： 提供原始网络协议存取。
  　      　SOCK_RDM： 提供可靠的数据包连接。
  　　      SOCK_PACKET： 与网络驱动程序直接通信。
  */
  sock_descriptor = socket(AF_INET, SOCK_STREAM, 0);//IPV4 TCP协议
  if (sock_descriptor == -1)
  {
    perror("call to socket");
    exit(1);
  }

  memset(&sin, 0, sizeof(sin));
  sin.sin_family = AF_INET;
  sin.sin_addr.s_addr = INADDR_ANY;
  printf("pin.sin_addr.s_addr = %d.\n", pin.sin_addr.s_addr);
  sin.sin_port = htons(port);

  if (bind(sock_descriptor, (struct sockaddr *)&sin, sizeof(sin)) == -1)
  {
    perror("call to bind");
    exit(1);
  }

  if (listen(sock_descriptor, 20) == -1) //在端口sock_descriptor监听
  {
    perror("call to listen");
    exit(1);
  }

  while(1)
  {
    printf("waiting connection.\n");
    temp_sock_descriptor = accept(sock_descriptor, (struct sockaddr *)&pin, &address_size);
    if (temp_sock_descriptor == -1)
    {
      perror("call to accept");
      exit(1);
    }

    memset(buf, 0, sizeof(buf));

    if (recv(temp_sock_descriptor, buf, 16384, 0) == -1)
    {
      perror("call to recv");
      exit(1);
    }

    printf("received from client:[%s]\n",buf);

    len = strlen(buf);
    for (i = 0; i < len; i++)
    {
      buf[i]= toupper(buf[i]);
    }

    if (send(temp_sock_descriptor, buf, len, 0) == -1)
    {
      perror("call to send");
      exit(1);
    }

    close(temp_sock_descriptor);
  }

  return (EXIT_SUCCESS);
}