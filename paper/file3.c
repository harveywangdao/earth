#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <string.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <errno.h>
#include <time.h>
#include <sys/time.h>
#include <sys/times.h>
#include <sys/uio.h>

void do1()
{
  char *file = "file.txt";

  int fd = open(file, O_RDWR|O_APPEND|O_CREAT, 0666);
  if (fd == -1)
  {
    perror("open fail");
    return ;
  }

  char *msgbuf[] = {"1111111", "2222222", "3333333"};
  int max = sizeof(msgbuf)/sizeof(msgbuf[0]);
  struct iovec wrios[max];

  for (int i = 0; i < max; ++i)
  {
    wrios[i].iov_base = msgbuf[i];
    wrios[i].iov_len = strlen(msgbuf[i]);
  }

  ssize_t nb = writev(fd, wrios, max);
  if (nb == -1)
  {
    perror("writev fail");
    return ;
  }
  close(fd);

  fd = open(file, O_RDONLY);
  if (fd == -1)
  {
    perror("open fail");
    return;
  }

  struct iovec rdios[max];
  char rdbuf[max][8];
  memset(rdbuf, 0, sizeof(rdbuf));
  printf("sizeof(rdbuf):%ld\n", sizeof(rdbuf));

  for (int i = 0; i < max; ++i)
  {
    rdios[i].iov_base = rdbuf[i];
    rdios[i].iov_len = sizeof(rdbuf[i])-1;
  }

  nb = readv(fd, rdios, max);
  if (nb == -1)
  {
    perror("readv fail");
    return ;
  }

  for (int i = 0; i < max; ++i)
  {
    printf("rdbuf[%d]:%s\n", i, rdbuf[i]);
  }

  close(fd);
  remove(file);
}

int main(int argc, char const *argv[])
{
  do1();

  return 0;
}