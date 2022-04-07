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
#include <sys/select.h>
#include <poll.h>
#include <sys/epoll.h>
#include <sys/un.h>
#include <stddef.h>

void do1()
{
  int ret;
  char buf[128];
  int fds[2];

  ret = socketpair(AF_UNIX, SOCK_STREAM, 0, fds);
  if (ret == -1)
  {
    perror("socketpair fail");
    return;
  }

  char *msg = "socket domain";
  int nb = write(fds[0], msg, strlen(msg));
  if (nb == -1)
  {
    perror("write fail");
    return;
  }

  memset(buf, 0, sizeof(buf));
  nb = read(fds[1], buf, sizeof(buf));
  if (nb == -1)
  {
    perror("read fail");
    return;
  }

  printf("buf:%s\n", buf);
}

void do2()
{
  int ret;
  int fds[2];
  ret = socketpair(AF_UNIX, SOCK_STREAM, 0, fds);
  if (ret == -1)
  {
    perror("socketpair fail");
    return;
  }

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

    char *msg = "client socket domain";
    int nb = write(fds[0], msg, strlen(msg));
    if (nb == -1)
    {
      perror("write fail");
    }

    printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
    exit(0);
  }
  else
  {
    char buf[128];

    memset(buf, 0, sizeof(buf));
    int nb = read(fds[1], buf, sizeof(buf));
    if (nb == -1)
    {
      perror("read fail");
    }

    printf("buf:%s\n", buf);    

    int status = 0;
    ret = waitpid(pid, &status, 0);
    if (ret == -1)
    {
      printf("son ret = %d, status = %d\n", ret, status);
      return;
    }

    printf("son ret = %d, status = %d\n", ret, status);
  }
}

/*struct sockaddr_un {
  __kernel_sa_family_t sun_family; //AF_UNIX
  char sun_path[UNIX_PATH_MAX];   //pathname
};*/

void do3()
{
  int ret;
  int fds[2];
  ret = socketpair(AF_UNIX, SOCK_STREAM, 0, fds);
  if (ret == -1)
  {
    perror("socketpair fail");
    return;
  }

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

    char recvbuf[128];

    int fd = socket(AF_UNIX, SOCK_STREAM, 0);
    if (fd == -1)
    {
      perror("socket fail");
    }

    struct sockaddr_un un, sun;
    memset(&un, 0, sizeof(un));
    un.sun_family = AF_UNIX;
    snprintf(un.sun_path, sizeof(un.sun_path), "%s%05d", "unixpath01_client", getpid());

    int len = offsetof(struct sockaddr_un, sun_path) + strlen(un.sun_path);
    if (bind(fd, (struct sockaddr *)(&un), len) == -1)
    {
      perror("bind fail");
    }

    sun.sun_family = AF_UNIX;
    strcpy(sun.sun_path, "unixpath01");
    len = offsetof(struct sockaddr_un, sun_path) + strlen(sun.sun_path);
    if (connect(fd, (struct sockaddr *)(&sun), len) == -1)
    {
      perror("connect fail");
    }

    const char sendbuf[] = "I am client xiao ming.";
    if (send(fd, sendbuf, strlen(sendbuf), 0) == -1)
    { 
      perror("send fail");
    }

    memset(recvbuf, 0, sizeof(recvbuf));
    int nbytes = recv(fd, recvbuf, sizeof(recvbuf), 0);
    if (nbytes < 0)
    { 
      perror("recv fail");
    }
    recvbuf[nbytes] = '\0';
    printf("client recv data:%s\n", recvbuf);

    close(fd);
    unlink(un.sun_path);

    printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
    exit(0);
  }
  else
  {
    char buffer[128];
    struct sockaddr_un un;
    un.sun_family = AF_UNIX;
    strcpy(un.sun_path, "unixpath01");

    int fd = socket(AF_UNIX, SOCK_STREAM, 0);
    if (fd == -1)
    {
      perror("socket fail");
    }

    int len = offsetof(struct sockaddr_un, sun_path) + strlen(un.sun_path);
    if (bind(fd, (struct sockaddr *)(&un), len) == -1)
    {
      perror("bind fail");
    }

    if (listen(fd, 5) == -1)
    {
      perror("listen fail");
    }

    struct sockaddr_un clientun;
    socklen_t unlen;
    unlen = sizeof(struct sockaddr_un);
    memset(&clientun, 0, sizeof(struct sockaddr_un));
    int connfd;
    if ((connfd = accept(fd, (struct sockaddr *)(&clientun), &unlen)) == -1)
    {
      perror("accept fail");
    }
    printf("server unlen:%d, server get connection from ip:%s\n", unlen, clientun.sun_path);
    
    memset(buffer, 0, sizeof(buffer));
    int nbytes = recv(connfd, buffer, sizeof(buffer), 0);
    if (nbytes < 0)
    { 
      perror("recv fail");
    }
    buffer[nbytes] = '\0';
    printf("server recv data:%s\n", buffer);

    const char hello[] = "I am unix domain server xiao hong.";

    if (send(connfd, hello, strlen(hello), 0) == -1)
    { 
      perror("send fail");
    }

    sleep(1);
    close(fd);
    close(connfd);

    unlink(un.sun_path);

    int status = 0;
    ret = waitpid(pid, &status, 0);
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
  //do1();
  //do2();
  do3();

  return 0;
}