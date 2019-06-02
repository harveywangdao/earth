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
#include <aio.h>

//gcc -o app aio.c -lrt

void do1()
{
  char msg[] = "this is an aio test file, hello aio.";

  int fd = open("file.txt", O_RDWR|O_APPEND|O_CREAT, 0666);
  if (fd == -1)
  {
    perror("open fail");
    return;
  }

  int nb = write(fd, msg, strlen(msg));
  if (nb == -1)
  {
    perror("write fail");
    return;
  }
  close(fd);

  fd = open("file.txt", O_RDONLY);
  if (fd == -1)
  {
    perror("open fail");
    return;
  }

  struct aiocb faio;
  int BUFSIZE = 128;
  memset(&faio, 0, sizeof(struct aiocb));
  char *buf = malloc(BUFSIZE * sizeof(char));
  if (buf == NULL)
  {
    perror("malloc fail");
    return;
  }

  for (int i = 0; i < BUFSIZE; i++)
  {
    buf[i] = '2';
  }
  memset(buf, 0, BUFSIZE * sizeof(char));

  faio.aio_buf = buf;
  faio.aio_fildes = fd;
  faio.aio_nbytes = BUFSIZE;
  faio.aio_offset = 0;

  int ret = aio_read(&faio);
  if (ret == -1)
  {
    perror("aio_read fail");
    return;
  }

  ssize_t err;
  while (1)
  {
    err = aio_error(&faio);
    if (err == 0)
    {
      printf("aio_read done\n");
      break;
    }
    else if (err == -1)
    {
      perror("aio_error fail");
      return;
    }
    else if (err == EINPROGRESS)
    {
      printf("aio_read EINPROGRESS\n");
      continue;
    }
    else
    {
      perror("aio_error fail");
      return;
    }
  }

  ret = aio_return(&faio);
  if (ret == -1)
  {
    perror("aio_return fail");
    return;
  }
  else
  {
    printf("aio_return:%d\n", ret);
    printf("faio.aio_buf:%s\n", (char*)faio.aio_buf);
  }

  free(buf);
  close(fd);
  remove("file.txt");
}

int main(int argc, char const *argv[])
{
  do1();

  return 0;
}