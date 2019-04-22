#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <string.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <errno.h>

#define FIFO_PATH "/home/thomas/golang/src/earth/paper/fifo_pipe"

void write1()
{
  int retn;
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

    int rfd = open(FIFO_PATH, O_RDONLY);
    if (rfd == -1)
    {
      printf("open fail\n");
      exit(0);
    }

    char rbuff[1024];

    retn = read(rfd, rbuff, sizeof(rbuff)-1);
    if (retn == -1)
    {
      printf("read fail\n");
      exit(0);
    }

    rbuff[retn] = '\0';

    printf("son read %d:%s\n", retn, rbuff);

    close(rfd);

    printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
    exit(0);
  }
  else
  {
    int wfd = open(FIFO_PATH, O_WRONLY);
    if (wfd == -1)
    {
      printf("open fail\n");
      exit(0);
    }

    char wbuff[1024] = "fifo, hello";

    retn = write(wfd, wbuff, strlen(wbuff));
    if (retn == -1)
    {
      printf("write fail\n");
      exit(0);
    }

    close(wfd);

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

void write2()
{
  int retn;
  pid_t pid;
  pid = fork();
  if (pid == -1)
  {
    printf("fork fail\n");
    return;
  }
  else if (pid == 0)
  {
    /*
    int flag;
    flag = fcntl(fd, F_GETFL, 0);
     
    flag |= O_NONBLOCK;
    fcntl(fd, F_SETFL, flag);
    */

    printf("son start, pid = %d, ppid = %d\n", getpid(), getppid());
    
    //sleep(10);

    printf("son fifo opening\n");
    int rfd = open(FIFO_PATH, O_RDONLY | O_NONBLOCK);
    //int rfd = open(FIFO_PATH, O_RDONLY);
    if (rfd == -1)
    {
      printf("open fail\n");
      exit(0);
    }
    printf("son fifo open done\n");

    char rbuff[1024];

    //sleep(10);

    printf("son reading\n");

    retn = read(rfd, rbuff, sizeof(rbuff)-1);
    if (retn == -1)
    {
      printf("read fail:%s\n", strerror(errno));
      perror("1111");
      exit(0);
    }

    rbuff[retn] = '\0';

    printf("son read %d:%s\n", retn, rbuff);

    close(rfd);
    
    printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
    exit(0);
  }
  else
  {
    sleep(10);

    printf("fifo opening\n");
    int wfd = open(FIFO_PATH, O_WRONLY);
    if (wfd == -1)
    {
      printf("open fail\n");
      exit(0);
    }
    printf("fifo open done\n");

    char wbuff[1024] = "fifo, hello";

    //sleep(10);

    printf("fifo writing\n");
    retn = write(wfd, wbuff, strlen(wbuff));
    if (retn == -1)
    {
      printf("write fail\n");
      exit(0);
    }
    printf("fifo write done\n");

    close(wfd);

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
  int ret = mkfifo(FIFO_PATH, S_IRUSR | S_IWUSR);
  if (ret == -1 && errno != EEXIST)
  {
    printf("mkfifo fail\n");
    return -1;
  }

  //write1();
  write2();

  unlink(FIFO_PATH);

  sleep(2);
  return 0;
}