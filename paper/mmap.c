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

int main(int argc, char const *argv[])
{
	do1();
	return 0;
}