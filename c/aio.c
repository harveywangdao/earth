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

void genfile(const char *filename)
{
  char content[] = "this is an aio test file, hello aio.";
  char buf[128];

  int fd = open(filename, O_RDWR|O_APPEND|O_CREAT, 0666);
  if (fd == -1)
  {
    perror("open fail");
    return;
  }

  snprintf(buf, sizeof(buf), "%s %s", content, filename);

  int nb = write(fd, buf, strlen(buf));
  if (nb == -1)
  {
    perror("write fail");
    return;
  }

  close(fd);
}

void do2()
{
  const char *files[] = {"file1.txt", "file2.txt", "file3.txt"};
  int filenum = sizeof(files)/sizeof(files[0]);
  int rdfds[filenum];
  printf("filenum:%d\n", filenum);

  for (int i = 0; i < filenum; ++i)
  {
    genfile(files[i]);
    
    rdfds[i] = open(files[i], O_RDONLY);
    if (rdfds[i] == -1)
    {
      perror("open fail");
      return;
    }
  }

  struct aiocb faios[filenum];
  const struct aiocb *faiolist[filenum];
  int BUFSIZE = 128;
  int ret;
  
  memset(faios, 0, sizeof(faios));

  for (int i = 0; i < filenum; ++i)
  {
    char *buf = malloc(BUFSIZE * sizeof(char));
    if (buf == NULL)
    {
      perror("malloc fail");
      return;
    }
    memset(buf, 0, BUFSIZE * sizeof(char));

    faios[i].aio_buf = buf;
    faios[i].aio_fildes = rdfds[i];
    faios[i].aio_nbytes = BUFSIZE;
    faios[i].aio_offset = 0;

    ret = aio_read(&faios[i]);
    if (ret == -1)
    {
      perror("aio_read fail");
      return;
    }

    faiolist[i] = &faios[i];
  }

  ssize_t err;
  int done = 0;
  while (1)
  {
    for (int i = 0; i < filenum; ++i)
    {
      if (faiolist[i] == NULL)
      {
        continue;
      }

      err = aio_error(&faios[i]);
      if (err == 0)
      {
        printf("aio_read done\n");

        ret = aio_return(&faios[i]);
        if (ret == -1)
        {
          printf("aio_return fail\n");
          return;
        }
        else
        {
          printf("aio_return:%d\n", ret);
          printf("faios[%d].aio_buf:%s\n", i, (char*)faios[i].aio_buf);
        }

        faiolist[i] = NULL;

        done++;
      }
      else if (err == -1)
      {
        printf("aio_error fail\n");
        return;
      }
      else if (err == EINPROGRESS)
      {
        printf("aio_read EINPROGRESS faios[%d]\n", i);
        continue;
      }
      else
      {
        printf("aio_error fail\n");
        return;
      }
    }

    if (done >= filenum)
    {
      break;
    }
    else
    {
      printf("aio_suspend starting\n");
      ret = aio_suspend(faiolist, filenum, NULL);
      if (ret == -1)
      {
        printf("aio_suspend fail");
        return;
      }
      printf("aio_suspend end\n");

      /*for (int i = 0; i < filenum; ++i)
      {
        if (faiolist[i] != NULL)
        {
          printf("faiolist[%d]->aio_buf:%s\n", i, (char*)faiolist[i]->aio_buf);
        }
      }*/
    }
  }

  for (int i = 0; i < filenum; ++i)
  {
    free((void *)faios[i].aio_buf);
    close(rdfds[i]);
    remove(files[i]);
  }
}

void readfile(const char *filename)
{
  char buf[128];
  int fd = open(filename, O_RDONLY);
  if (fd == -1)
  {
    perror("open fail");
    return;
  }

  memset(buf, 0, sizeof(buf));
  int nb = read(fd, buf, sizeof(buf));
  if (nb == -1)
  {
    perror("read fail");
    return;
  }

  if (nb >= sizeof(buf))
  {
    buf[sizeof(buf)-1] = '\0';
  }

  printf("%s:%s\n", filename, buf);
  close(fd);
}

void do3()
{
  int ret;
  int fd = open("file.txt", O_RDWR|O_APPEND|O_CREAT, 0666);
  if (fd == -1)
  {
    perror("open fail");
    return;
  }

  struct aiocb wraiocb;
  memset(&wraiocb, 0, sizeof(struct aiocb));

  char buf[] = "this is an aio test file, hello aio.";

  wraiocb.aio_buf = buf;
  wraiocb.aio_fildes = fd;
  wraiocb.aio_nbytes = strlen(buf);
  wraiocb.aio_offset = 0;
  
  ret = aio_write(&wraiocb);
  if (ret == -1)
  {
    printf("aio_write fail\n");
    return;
  }

  ssize_t err;
  while (1)
  {
    err = aio_error(&wraiocb);
    if (err == 0)
    {
      printf("aio_write done\n");
      break;
    }
    else if (err == -1)
    {
      perror("aio_error fail");
      return;
    }
    else if (err == EINPROGRESS)
    {
      printf("aio_write EINPROGRESS\n");
      continue;
    }
    else
    {
      perror("aio_error fail");
      return;
    }
  }

  ret = aio_return(&wraiocb);
  if (ret == -1)
  {
    perror("aio_return fail");
    return;
  }
  else
  {
    printf("wraiocb.aio_nbytes:%ld, aio_return:%d\n", wraiocb.aio_nbytes, ret);
  }

  ret = aio_fsync(O_SYNC, &wraiocb);
  if (ret == -1)
  {
    printf("aio_fsync fail\n");
    return;
  }
  close(fd);

  readfile("file.txt");
  remove("file.txt");
}

int main(int argc, char const *argv[])
{
  //do1();
  //do2();
  do3();

  return 0;
}