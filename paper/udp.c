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

#define port 9057

void server()
{
  int sockfd;
  struct sockaddr_in server_addr;
  struct sockaddr_in client_addr;
  socklen_t sinlen;
  ssize_t nbytes;
  const char hello[] = "I am server xiao hong";
  char buffer[128];

  if ((sockfd = socket(AF_INET, SOCK_DGRAM, 0)) == -1)
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

  int cnt = 0;
  char sendbuf[64];
  while(cnt < 2)
  {
    cnt++;

    sinlen = sizeof(struct sockaddr_in);
    memset(&client_addr, 0, sizeof(struct sockaddr_in));
    memset(buffer, 0, sizeof(buffer));
    nbytes = recvfrom(sockfd, buffer, sizeof(buffer), 0, (struct sockaddr *)&client_addr, &sinlen);
    if (nbytes < 0)
    { 
      perror("recvfrom fail");
      return;
    }
    buffer[nbytes] = '\0';
    printf("server sinlen:%d, server get connection from ip:%s\n", sinlen, inet_ntoa(client_addr.sin_addr));
    printf("server recv data:%s\n", buffer);

    snprintf(sendbuf, sizeof(sendbuf),"%s%d", hello, cnt);

    nbytes = sendto(sockfd, sendbuf, strlen(sendbuf), 0, (struct sockaddr *)&client_addr, sizeof(struct sockaddr_in));
    if (nbytes == -1)
    { 
      perror("sendto fail");
      return;
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
  socklen_t sinlen;
  ssize_t nbytes;

  if ((sockfd = socket(AF_INET, SOCK_DGRAM, 0)) == -1)
  {
    perror("socket fail");
    return;
  }

  memset(&server_addr, 0, sizeof(server_addr));
  server_addr.sin_family = AF_INET;
  server_addr.sin_port = htons(port);
  server_addr.sin_addr.s_addr = inet_addr("127.0.0.1");

  nbytes = sendto(sockfd, sendbuf, strlen(sendbuf), 0, (struct sockaddr *)&server_addr, sizeof(struct sockaddr_in));
  if (nbytes == -1)
  { 
    perror("sendto fail");
    return;
  }

  sinlen = sizeof(struct sockaddr_in);
  memset(&server_addr, 0, sizeof(struct sockaddr_in));
  memset(recvbuf, 0, sizeof(recvbuf));
  nbytes = recvfrom(sockfd, recvbuf, sizeof(recvbuf), 0, (struct sockaddr *)&server_addr, &sinlen);
  if (nbytes < 0)
  { 
    perror("recvfrom fail");
    return;
  }
  recvbuf[nbytes] = '\0';
  printf("client sinlen:%d, server get connection from ip:%s\n", sinlen, inet_ntoa(server_addr.sin_addr));
  printf("client recv data:%s\n", recvbuf);

  printf("\n");

  close(sockfd);
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