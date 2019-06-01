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

int port = 9057;

void server0()
{
  int sockfd, connfd;
  struct sockaddr_in server_addr;
  struct sockaddr_in client_addr;
  socklen_t sinlen;
  ssize_t nbytes;
  const char hello[] = "I am server xiao hong";
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

  printf("sizeof(fd_set)=%ld\n", sizeof(fd_set));

  int selectnum;
  int maxfd;

  fd_set rdset, wrset, exceptset;
  fd_set originrdset, originwrset, originexceptset;

  FD_ZERO(&originrdset);
  FD_ZERO(&originwrset);
  FD_ZERO(&originexceptset);

  FD_SET(sockfd, &originrdset);
  maxfd = sockfd;

  while(1)
  {
    rdset = originrdset;

    selectnum = select(maxfd+1, &rdset, NULL, NULL, NULL);
    if (selectnum == -1)
    {
      perror("select fail");
      return;
    }
    else if (selectnum == 0)
    {
      printf("selectnum=0\n");
      continue;
    }
    else
    {
      for (int fd = 0; fd < maxfd+1; fd++)
      {
        if (FD_ISSET(fd, &rdset))
        {
          if (fd == sockfd)
          {
            sinlen = sizeof(struct sockaddr_in);
            memset(&client_addr, 0, sizeof(struct sockaddr_in));
            if ((connfd = accept(fd, (struct sockaddr *)(&client_addr), &sinlen)) == -1)
            {
              perror("accept fail");
              return;
            }
            printf("server sinlen:%d, server get connection from ip:%s\n", sinlen, inet_ntoa(client_addr.sin_addr));
            FD_SET(connfd, &originrdset);

            if (connfd > maxfd)
            {
              maxfd = connfd;
            }
          }
          else
          {
            memset(buffer, 0, sizeof(buffer));
            nbytes = recv(fd, buffer, sizeof(buffer), 0);
            if (nbytes < 0)
            { 
              perror("recv fail");
              return;
            }
            else if (nbytes == 0)
            {
              close(fd);
              FD_CLR(fd, &originrdset);
              printf("conn close\n\n\n");
              continue;
            }

            buffer[nbytes] = '\0';
            printf("server recv data:%s\n", buffer);

            snprintf(buffer, sizeof(buffer), "%s%d", hello, fd);

            if (send(fd, buffer, strlen(buffer), 0) == -1)
            { 
              perror("send fail");
              return;
            }
          }
        }
      }
    }
  }

  close(sockfd);
}

void server()
{
  int sockfd, connfd;
  struct sockaddr_in server_addr;
  struct sockaddr_in client_addr;
  socklen_t sinlen;
  ssize_t nbytes;
  const char hello[] = "I am server xiao hong";
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

  printf("sizeof(fd_set)=%ld\n", sizeof(fd_set));

  int selectnum;
  fd_set rdset, wrset, exceptset;
  fd_set originrdset, originwrset, originexceptset;

  FD_ZERO(&originrdset);
  FD_ZERO(&originwrset);
  FD_ZERO(&originexceptset);

  FD_SET(sockfd, &originrdset);

  while(1)
  {
    rdset = originrdset;

    selectnum = select(FD_SETSIZE, &rdset, NULL, NULL, NULL);
    if (selectnum == -1)
    {
      perror("select fail");
      return;
    }
    else if (selectnum == 0)
    {
      printf("selectnum=0\n");
      continue;
    }
    else
    {
      for (int fd = 0; fd < FD_SETSIZE; fd++)
      {
        if (FD_ISSET(fd, &rdset))
        {
          if (fd == sockfd)
          {
            sinlen = sizeof(struct sockaddr_in);
            memset(&client_addr, 0, sizeof(struct sockaddr_in));
            if ((connfd = accept(fd, (struct sockaddr *)(&client_addr), &sinlen)) == -1)
            {
              perror("accept fail");
              return;
            }
            printf("server sinlen:%d, server get connection from ip:%s\n", sinlen, inet_ntoa(client_addr.sin_addr));
            FD_SET(connfd, &originrdset);
          }
          else
          {
            memset(buffer, 0, sizeof(buffer));
            nbytes = recv(fd, buffer, sizeof(buffer), 0);
            if (nbytes < 0)
            { 
              perror("recv fail");
              return;
            }
            else if (nbytes == 0)
            {
              close(fd);
              FD_CLR(fd, &originrdset);
              printf("conn close\n\n\n");
              continue;
            }

            buffer[nbytes] = '\0';
            printf("server recv data:%s\n", buffer);

            snprintf(buffer, sizeof(buffer), "%s%d", hello, fd);

            if (send(fd, buffer, strlen(buffer), 0) == -1)
            { 
              perror("send fail");
              return;
            }
          }
        }
      }
    }
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
  
  if (connect(sockfd, (struct sockaddr *)(&server_addr), sizeof(struct sockaddr)) == -1)
  {
    perror("connect fail");
    return;
  }

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

    for (int i = 0; i < 10; ++i)
    {
      client();
      sleep(1);
    }

    printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
    exit(0);
  }
  else
  {
    server0();
    //server();

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