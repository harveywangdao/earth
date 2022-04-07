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

int port = 9057;

/*
struct epitem
{
  struct rb_node rbn;            //用于主结构管理的红黑树
  struct list_head rdllink;      //事件就绪队列
  struct epitem *next;           //用于主结构体中的链表
  struct epoll_filefd ffd;       //每个fd生成的一个结构
  int nwait;                 
  struct list_head pwqlist;      //poll等待队列
  struct eventpoll *ep;          //该项属于哪个主结构体
  struct list_head fllink;       //链接fd对应的file链表
  struct epoll_event event;      //注册的感兴趣的事件,也就是用户空间的epoll_event
}

struct eventpoll
{
  spin_lock_t lock;            //对本数据结构的访问
  struct mutex mtx;            //防止使用时被删除
  wait_queue_head_t wq;        //sys_epoll_wait() 使用的等待队列
  wait_queue_head_t poll_wait; //file->poll()使用的等待队列
  struct list_head rdllist;    //事件满足条件的链表
  struct rb_root rbr;          //用于管理所有fd的红黑树
  struct epitem *ovflist;      //将事件到达的fd进行链接起来发送至用户空间
}

typedef union epoll_data
{
  void *ptr;
  int fd;
  uint32_t u32;
  uint64_t u64;
}epoll_data_t;

struct epoll_event
{
  uint32_t events;      //Epoll events
  epoll_data_t data;    //User data variable
}

int epoll_create(int size);
int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event)
int epoll_wait (int epfd, struct epoll_event *events, int maxevents, int timeout)
EPOLL_CTL_ADD EPOLL_CTL_DEL EPOLL_CTL_MOD
EPOLLIN EPOLLOUT EPOLLRDNORM EPOLLRDBAND EPOLLWRNORM EPOLLWRBAND EPOLLET
*/

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

  int epfd = epoll_create(128);
  if (epfd == -1)
  {
    perror("epoll_create fail");
    return;
  }

  struct epoll_event ev;
  ev.events = EPOLLIN|EPOLLET;
  ev.data.fd = sockfd;

  int ret = epoll_ctl(epfd, EPOLL_CTL_ADD, sockfd, &ev);
  if (ret == -1)
  {
    perror("epoll_ctl fail");
    return;
  }

  struct epoll_event events[128];
  int fdnum;

  while(1)
  {
    fdnum = epoll_wait(epfd, events, 128, -1);//timeout 0:直接返回 -1:阻塞
    if (fdnum == -1)
    {
      perror("epoll_wait fail");
      return;
    }

    for (int i = 0; i < fdnum; ++i)
    {
      if (events[i].data.fd == sockfd)
      {
        sinlen = sizeof(struct sockaddr_in);
        memset(&client_addr, 0, sizeof(struct sockaddr_in));
        if ((connfd = accept(sockfd, (struct sockaddr *)(&client_addr), &sinlen)) == -1)
        {
          perror("accept fail");
          return;
        }
        printf("server sinlen:%d, server get connection from ip:%s\n", sinlen, inet_ntoa(client_addr.sin_addr));

        ev.data.fd = connfd;
        ev.events = EPOLLIN|EPOLLET;
        ret = epoll_ctl(epfd, EPOLL_CTL_ADD, connfd, &ev);
        if (ret == -1)
        {
          perror("epoll_ctl fail");
          return;
        }
      }
      else if (events[i].events & EPOLLIN)
      {
        memset(buffer, 0, sizeof(buffer));
        nbytes = recv(events[i].data.fd, buffer, sizeof(buffer), 0);
        if (nbytes < 0)
        {
          perror("recv fail");
          return;
        }
        else if (nbytes == 0)
        {
          ret = epoll_ctl(epfd, EPOLL_CTL_DEL, events[i].data.fd, &events[i]);
          if (ret == -1)
          {
            perror("epoll_ctl fail");
            return;
          }
          
          close(events[i].data.fd);
          printf("conn close\n\n\n");
          continue;
        }

        buffer[nbytes] = '\0';
        printf("server recv data:%s\n", buffer);

        snprintf(buffer, sizeof(buffer), "%s%d", hello, events[i].data.fd);

        if (send(events[i].data.fd, buffer, strlen(buffer), 0) == -1)
        { 
          perror("send fail");
          return;
        }
      }
      else if (events[i].events & EPOLLOUT)
      {
        
      }
    }
  }

  close(epfd);
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