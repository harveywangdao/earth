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
#include <sys/mman.h>

void do1()
{
  int ret;
  int fd = open("file.txt", O_RDWR|O_CREAT|O_TRUNC, 0666);
  if (fd == -1)
  {
    perror("open fail");
    return;
  }

  char buf[512];
  char *str = "11112222";
  memcpy(buf, str, strlen(str)+1);

  int nb = write(fd, buf, strlen(buf)+1);
  if (nb == -1)
  {
    perror("write fail");
    return;
  }

  lseek(fd, 0, SEEK_SET);
  memset(buf, 0, sizeof(buf));
  nb = read(fd, buf, 9);
  if (nb == -1)
  {
    perror("read fail");
    return;
  }
  printf("1 read %d:%s\n", nb, buf);

  struct stat st;
  fstat(fd, &st);

  int mapsize = st.st_size;

  char *addr = mmap(0, mapsize, PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
  if (addr == MAP_FAILED)
  {
  	perror("mmap fail");
  	return;
  }

  printf("addr:%s\n", addr);
  addr[0] = 'a';
  addr[1] = 'b';
  addr[2] = 'c';
  addr[3] = 'd';

  ret = msync(addr, mapsize, MS_SYNC);
  if (ret == -1)
  {
  	perror("msync fail");
  	return;
  }

  ret = munmap(addr, mapsize);
  if (ret == -1)
  {
  	perror("munmap fail");
  	return;
  }

  lseek(fd, 0, SEEK_SET);
  memset(buf, 0, sizeof(buf));
  nb = read(fd, buf, mapsize);
  if (nb == -1)
  {
    perror("read fail");
    return;
  }
  printf("2 read %d:%s\n", nb, buf);

  close(fd);
  remove("file.txt");
}

void do2()
{
  int ret;
  int mapsize = 32;

  char *addr = mmap(0, mapsize, PROT_READ | PROT_WRITE, MAP_SHARED|MAP_ANON, -1, 0);
  if (addr == MAP_FAILED)
  {
    perror("mmap fail");
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

    addr[0] = 'a';
    addr[1] = 'b';
    addr[2] = 'c';
    addr[3] = 'd';

    printf("son addr:%s\n", addr);

    ret = msync(addr, mapsize, MS_SYNC);
    if (ret == -1)
    {
      perror("msync fail");
      return;
    }
    
    ret = munmap(addr, mapsize);
    if (ret == -1)
    {
      perror("munmap fail");
      return;
    }

    printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
    exit(0);
  }
  else
  {
    sleep(2);
    printf("addr:%s\n", addr);

    ret = msync(addr, mapsize, MS_SYNC);
    if (ret == -1)
    {
      perror("msync fail");
      return;
    }
    
    ret = munmap(addr, mapsize);
    if (ret == -1)
    {
      perror("munmap fail");
      return;
    }

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
  //do1();
	do2();

	return 0;
}