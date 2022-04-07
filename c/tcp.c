#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <netinet/in.h>
#include <stdint.h>
#include <netdb.h>
#include <time.h>

/*
void socket_test()
{
  int ret;
  int fd = socket(AF_INET, SOCK_STREAM, 0);//1.AF_INET6 AF_UNIX; 2.SOCK_DGRAM SOCK_RAW SOCK_SEQPACKET
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
    sa_family_t sa_family;    //Common data: address family and length.
    char sa_data[14];       //Address data.
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
    in_port_t sin6_port;    //Transport layer port #
    uint32_t sin6_flowinfo; //IPv6 flow information
    struct in6_addr sin6_addr;  //IPv6 address
    uint32_t sin6_scope_id; //IPv6 scope-id
  };

  struct in_addr
  {
    in_addr_t s_addr;
  };

  struct sockaddr_in
  {
    sa_family_t sin_family;
    in_port_t sin_port;         //Port number.
    struct in_addr sin_addr;        //Internet address.

    //Pad to size of `struct sockaddr'.
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

  int val = 1024;
  len = sizeof(val);
  ret = setsockopt(fd, SOL_SOCKET, SO_RCVBUF, &val, len);
  if (ret == -1)
  {
    printf("setsockopt fail\n");
    return;
  }

  ret = getsockopt(fd, SOL_SOCKET, SO_RCVBUF, &val, &len);
  if (ret == -1)
  {
    printf("getsockopt fail\n");
    return;
  }

  ret = shutdown(fd, SHUT_WR);//SHUT_RDWR SHUT_RD
  if (ret == -1)
  {
    printf("shutdown fail\n");
    return;
  }

  in_addr_t inet_addr (const char *cp)
  char *inet_ntoa (struct in_addr in)
  int inet_aton (const char *cp, struct in_addr *inp)
  #define INADDR_ANY ((unsigned long int) 0x00000000)
  int gethostname (char *name, size_t len)
  struct hostent *gethostbyname (const char *name)

  struct hostent
  {
    char *h_name;         //Official name of host.  
    char **h_aliases;     //Alias list.  
    int h_addrtype;       //Host address type.  
    int h_length;         //Length of address.  
    char **h_addr_list;   //List of addresses from name server.  
  };

  close(fd);
}
*/

void test1()
{
  struct in_addr in;
  int ret = inet_aton("192.168.1.17", &in);
  if (ret == 0)
  {
    printf("inet_aton fail\n");
    return;
  }
  printf("inet_ntoa(in) = %s\n", inet_ntoa(in));

  char hostname[128];
  ret = gethostname(hostname, sizeof(hostname));
  if (ret != 0)
  {
    printf("gethostname fail\n");
    return;
  }
  printf("hostname=%s\n", hostname);

  struct hostent *host;
  host = gethostbyname(hostname);
  if (host == NULL)
  {
    printf("gethostbyname fail\n");
    return;
  }
  printf("%s, %d, %d\n", host->h_name, host->h_addrtype, host->h_length);

  int i = 0;
  while(host->h_aliases[i] != NULL)
  {
    printf("h_aliases:%s\n", host->h_aliases[i]);
    i++;
  }

  i = 0;
  while(host->h_addr_list[i] != NULL)
  {
    printf("h_addr_list:%s\n", inet_ntoa(*(struct in_addr*)host->h_addr_list[i]));
    i++;
  }

  char addrstr[] = "192.168.1.25";
  char netaddr[128];

  memset(netaddr, 0, sizeof(netaddr));
  ret = inet_pton(AF_INET, addrstr, netaddr);
  if (ret != 1)
  {
    printf("inet_pton fail\n");
    return;
  }

  memset(addrstr, 0, sizeof(addrstr));
  const char *p = inet_ntop(AF_INET, netaddr, addrstr, sizeof(addrstr));
  if (p == NULL)
  {
    printf("inet_ntop fail\n");
    return;
  }
  printf("addrstr:%s\n", addrstr);
}

void printsockopt(int fd)
{
  int rcvbufsz = 0;
  int sndbufsz = 0;
  socklen_t socklen;

  socklen = sizeof(rcvbufsz);
  int ret = getsockopt(fd, SOL_SOCKET, SO_RCVBUF, &rcvbufsz, &socklen);
  if (ret == -1)
  {
    printf("getsockopt fail\n");
    return;
  }
  printf("rcvbufsz:%d, socklen:%d\n", rcvbufsz, socklen);

  socklen = sizeof(sndbufsz);
  ret = getsockopt(fd, SOL_SOCKET, SO_SNDBUF, &sndbufsz, &socklen);
  if (ret == -1)
  {
    printf("getsockopt fail\n");
    return;
  }
  printf("sndbufsz:%d, socklen:%d\n", sndbufsz, socklen);
}

void modifysockopt(int fd, int bufsz)
{
  //对于客户， SO_RCVBUF选项必须在调用connect之前设置
  //对于服务器，该选项必须在调用listen之前给监听套接字设置
  //实际设置的值是X2
  //结果不能大于416k
  int rcvbufsz = bufsz;
  int sndbufsz = bufsz;
  socklen_t socklen;

  socklen = sizeof(rcvbufsz);
  int ret = setsockopt(fd, SOL_SOCKET, SO_RCVBUF, &rcvbufsz, socklen);
  if (ret == -1)
  {
    printf("setsockopt fail\n");
    return;
  }
  printf("modifysockopt rcvbufsz:%d, socklen:%d\n", rcvbufsz, socklen);

  socklen = sizeof(sndbufsz);
  ret = setsockopt(fd, SOL_SOCKET, SO_SNDBUF, &sndbufsz, socklen);
  if (ret == -1)
  {
    printf("setsockopt fail\n");
    return;
  }
  printf("modifysockopt sndbufsz:%d, socklen:%d\n", sndbufsz, socklen);
}

int port = 9057;

void server()
{
  int sockfd, connfd;
  struct sockaddr_in server_addr;
  struct sockaddr_in client_addr;
  socklen_t sinlen;
  ssize_t nbytes;
  const char hello1[] = "I am server xiao hong1.";
  const char hello2[] = "I am server xiao hong2.";
  char buffer[128];

  if ((sockfd = socket(AF_INET, SOCK_STREAM, 0)) == -1)
  {
    perror("socket fail");
    return;
  }

  memset(&server_addr, 0, sizeof(struct sockaddr_in));
  server_addr.sin_family = AF_INET;
  server_addr.sin_addr.s_addr = htonl(INADDR_ANY);  //inet_addr("192.168.1.0")
  server_addr.sin_port = htons(port);

  printf("server listen address = %s\n", inet_ntoa(server_addr.sin_addr));

  if (bind(sockfd, (struct sockaddr *)(&server_addr), sizeof(struct sockaddr)) == -1)
  {
    perror("bind fail");
    return;
  }

  if (listen(sockfd, 5) == -1)
  {
    perror("listen fail");
    return;
  }

  while(1)
  {
    sinlen = sizeof(struct sockaddr_in);
    memset(&client_addr, 0, sizeof(struct sockaddr_in));
    if ((connfd = accept(sockfd, (struct sockaddr *)(&client_addr), &sinlen)) == -1)
    {
      perror("accept fail");
      return;
    }
    printf("server sinlen:%d, server get connection from ip:%s\n", sinlen, inet_ntoa(client_addr.sin_addr));
    
    memset(buffer, 0, sizeof(buffer));
    nbytes = recv(connfd, buffer, sizeof(buffer), MSG_DONTWAIT);
    if (nbytes < 0)
    { 
      perror("recv MSG_DONTWAIT fail");
    }
    buffer[nbytes] = '\0';
    printf("server recv data:%s\n", buffer);

    memset(buffer, 0, sizeof(buffer));
    nbytes = recv(connfd, buffer, sizeof(buffer), 0);
    if (nbytes < 0)
    { 
      perror("recv fail");
      return;
    }
    buffer[nbytes] = '\0';
    printf("server recv data:%s\n", buffer);

    if (send(connfd, hello1, strlen(hello1), 0) == -1)
    { 
      perror("send fail");
      return;
    }
    
    sleep(1);

    if(write(connfd, hello2, strlen(hello2)) == -1)
    {
      perror("write fail");
      return;
    }

    close(connfd);
    break;
  }

  close(sockfd);
}

void client()
{
  int sockfd;
  const char sendbuf[] = "I am client xiao ming.";
  char recvbuf[128];
  struct sockaddr_in server_addr;
  ssize_t nbytes;

  if ((sockfd = socket(AF_INET, SOCK_STREAM, 0)) == -1)
  {
    perror("socket fail");
    return;
  }

  memset(&server_addr, 0, sizeof(server_addr));
  server_addr.sin_family = AF_INET;
  server_addr.sin_port = htons(port);
  server_addr.sin_addr.s_addr = inet_addr("127.0.0.1");

  printsockopt(sockfd);
  modifysockopt(sockfd, 1024*128);
  printsockopt(sockfd);
  
  if (connect(sockfd, (struct sockaddr *)(&server_addr), sizeof(struct sockaddr)) == -1)
  {
    perror("connect fail");
    return;
  }

  sleep(1);

  if (send(sockfd, sendbuf, strlen(sendbuf), 0) == -1)
  { 
    perror("send fail");
    return;
  }

  memset(recvbuf, 0, sizeof(recvbuf));
  nbytes = recv(sockfd, recvbuf, sizeof(recvbuf), 0);
  if (nbytes < 0)
  { 
    perror("recv fail");
    return;
  }
  recvbuf[nbytes] = '\0';
  printf("client recv data:%s\n", recvbuf);
  
  memset(recvbuf, 0, sizeof(recvbuf));
  nbytes = read(sockfd, recvbuf, sizeof(recvbuf));
  if(nbytes == -1)
  {
    perror("read fail");
    return;
  }
  recvbuf[nbytes] = '\0';
  printf("client recv data:%s\n", recvbuf);

  close(sockfd);
}

void do1()
{
  srand(time(NULL));
  port = rand()%100 + 4000;
  printf("port:%d\n", port);

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
  //test1();

  return 0;
}