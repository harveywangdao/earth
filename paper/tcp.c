#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <netinet/in.h>
#include <stdint.h>
#include <netdb.h>

void server()
{
  int ret;
  int fd = socket(AF_INET, SOCK_STREAM, 0);//AF_INET6 AF_UNIX SOCK_DGRAM SOCK_RAW SOCK_SEQPACKET
  if (fd == -1)
  {
    printf("socket fail\n");
    return;
  }

  uint32_t host32 = 1, net32 = 1;
  uint16_t host16 = 1, net16 = 1;
  
  printf("htonl:%d\n", htonl(host32));
  printf("htons:%d\n", htons(host16));
  printf("ntohl:%d\n", ntohl(net32));
  printf("ntohs:%d\n", ntohs(net16));

  char addrstr[INET_ADDRSTRLEN];//INET6_ADDRSTRLEN
  unsigned char netaddr[4];//16
  char *p = inet_ntop(AF_INET, netaddr, addrstr, sizeof(addrstr));
  if (p == NULL)
  {
    printf("inet_ntop fail\n");
    return;
  }

  ret = inet_pton(AF_INET, addrstr, netaddr);
  if (ret != 1)
  {
    printf("inet_pton fail\n");
    return;
  }

  struct hostent *host = gethostent();
  if (host == NULL)
  {
    printf("gethostent fail\n");
    return;
  }

  sethostent(1);
  endhostent();

  struct netent *net = getnetbyaddr();
  if (net == NULL)
  {
    printf("getnetbyaddr fail\n");
    return;
  }

  net = getnetbyname(name);
  if (net == NULL)
  {
    printf("getnetbyname fail\n");
    return;
  }

  net = getnetent();
  if (net == NULL)
  {
    printf("getnetent fail\n");
    return;
  }

  setnetent(1);
  endnetent();

  struct protoent *proto = getprotobyname(name);
  if (proto == NULL)
  {
    printf("getprotobyname fail\n");
    return;
  }

  proto = getprotobunumber();
  if (proto == NULL)
  {
    printf("getprotobunumber fail\n");
    return;
  }

  proto = getprotoent();
  if (proto == NULL)
  {
    printf("getprotoent fail\n");
    return;
  }
  setprotoent(1);
  endprotoent();

  struct servent *serv = getservbyname();
  if (serv == NULL)
  {
    printf("getservbyname fail\n");
    return;
  }

  serv = getserbyport();
  if (serv == NULL)
  {
    printf("getserbyport fail\n");
    return;
  }

  serv = getservent();
  if (serv == NULL)
  {
    printf("getservent fail\n");
    return;
  }

  setservent(1);
  endservent();

  ret = getaddrinfo();
  if (ret != 0)
  {
    printf("%s\n", gai_strerror(ret));
    printf("getaddrinfo fail\n");
    return;
  }

  freeaddrinfo();

  ret = getnameinfo();
  if (ret != 0)
  {
    printf("%s\n", gai_strerror(ret));
    printf("getnameinfo fail\n");
    return;
  }

  typedef unsigned short int sa_family_t;
  typedef uint16_t in_port_t;
  typedef uint32_t in_addr_t;

  struct sockaddr
  {
    sa_family_t sa_family;    /* Common data: address family and length.  */
    char sa_data[14];       /* Address data.  */
  };

  struct in6_addr
  {
    union
    {
      uint8_t __u6_addr8[16];
      #ifdef __USE_MISC
      uint16_t __u6_addr16[8];
      uint32_t __u6_addr32[4];
      #endif
    } __in6_u;
    #define s6_addr         __in6_u.__u6_addr8
    #ifdef __USE_MISC
    # define s6_addr16      __in6_u.__u6_addr16
    # define s6_addr32      __in6_u.__u6_addr32
    #endif
  };

  struct sockaddr_in6
  {
    sa_family_t sin6_family;
    in_port_t sin6_port;    /* Transport layer port # */
    uint32_t sin6_flowinfo; /* IPv6 flow information */
    struct in6_addr sin6_addr;  /* IPv6 address */
    uint32_t sin6_scope_id; /* IPv6 scope-id */
  };

  struct in_addr
  {
    in_addr_t s_addr;
  };

  struct sockaddr_in
  {
    sa_family_t sin_family;
    in_port_t sin_port;         /* Port number.  */
    struct in_addr sin_addr;        /* Internet address.  */

    /* Pad to size of `struct sockaddr'.  */
    unsigned char sin_zero[sizeof (struct sockaddr) -
    __SOCKADDR_COMMON_SIZE -
    sizeof (in_port_t) -
    sizeof (struct in_addr)];
  };

  struct sockaddr_in addrin;
  addrin.sin_family = AF_INET;//sa_family_t
  addrin.sin_port = htons(3253);//in_port_t
  addrin.sin_addr.s_addr = inet_addr("192.168.1.7");//struct in_addr -> in_addr_t
  memset(addrin.sin_zero, 0, sizeof(addrin.sin_zero));

  ret = bind(fd, (struct sockaddr*)&addrin, sizeof(addrin));
  if (ret == -1)
  {
    printf("bind fail\n");
    return;
  }

  socklen_t len;
  ret = getsockname(fd, (struct sockaddr*)&addrin, &len);
  if (ret == -1)
  {
    printf("getsockname fail\n");
    return;
  }

  ret = getpeername(fd, (struct sockaddr*)&addrin, &len);
  if (ret == -1)
  {
    printf("getpeername fail\n");
    return;
  }

  ret = listen(fd, 1);
  if (ret == -1)
  {
    printf("listen fail\n");
    return;
  }

  ret = accept(fd, (struct sockaddr*)&addrin, &len)
  if (ret == -1)
  {
    printf("accept fail\n");
    return;
  }

  ret = connect(fd, (struct sockaddr*)&addrin, sizeof(addrin));
  if (ret == -1)
  {
    printf("connect fail\n");
    return;
  }

  ssize_t nsendbytes;
  unsigned char buf[128];
  nsendbytes = send(fd, buf, sizeof(buf), 0);
  if (nsendbytes == -1)
  {
    printf("send fail\n");
    return;
  }

  nsendbytes = sendto(fd, buf, sizeof(buf), 0, (struct sockaddr*)addrin, sizeof(addrin));
  if (nsendbytes == -1)
  {
    printf("sendto fail\n");
    return;
  }

  ssize_t nrecvbytes;
  nrecvbytes = recv(fd, buf, sizeof(buf), 0);
  if (nrecvbytes == -1)
  {
    printf("recv fail\n");
    return;
  }

  nrecvbytes = recvfrom(fd, buf, sizeof(buf), 0, (struct sockaddr*)&addrin, &len);
  if (nrecvbytes == -1)
  {
    printf("recvfrom fail\n");
    return;
  }

  ret = shutdown(fd, SHUT_WR);//SHUT_RDWR SHUT_RD
  if (ret == -1)
  {
    printf("shutdown fail\n");
    return;
  }

  close(fd);
}

void client()
{
  
}

void do1()
{
  pid_t pid;
  pid = fork();
  if (pid == -1)
  {
    printf("fork fail\n");
    return;
  }
  else if (pid == 0)
  {
    printf("son start, pid = %d, ppid = %d\n", getpid(), getppid());

    client();

    printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
    exit(0);
  }
  else
  {
    server();
    int status = 0;
    int ret = waitpid(pid, &status, 0);
    if (ret == -1)
    {
      printf("son ret = %d, status = %d\n", ret, status);
      return;
    }

    printf("son ret = %d, status = %d\n", ret, status);
  }
}

int main(int argc, char const *argv[])
{
  do1();

  return 0;
}