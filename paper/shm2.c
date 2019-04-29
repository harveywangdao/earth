#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <string.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <errno.h>
#include <sys/ipc.h>
#include <sys/msg.h>
#include <sys/sem.h>
#include <sys/shm.h>
#include <sys/mman.h>

//gcc -o app shm2.c -lrt

#define SHM_NAME "/bigcup"

void do1()
{
  int ret = 0;
  //int fd = shm_open(SHM_NAME, O_RDWR|O_CREAT|O_EXCL, 0666);
  int fd = shm_open(SHM_NAME, O_RDWR|O_CREAT, 0666);
  if (fd == -1)
  {
    printf("shm_open fail\n");
    return;
  }

  int mapsize = 128;
  ret = ftruncate(fd, mapsize);
  if (ret == -1)
  {
    printf("ftruncate fail\n");
    return;
  }

  struct stat st;
  ret = fstat(fd, &st);
  if (ret == -1)
  {
    perror("fstat fail");
    return;
  }
  printf("mode = %d, size = %ld, mtime = %ld\n", st.st_mode, st.st_size, st.st_mtime);

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

    char *addr = mmap(0, mapsize, PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
    if (addr == MAP_FAILED)
    {
      perror("mmap fail");
      exit(0);
    }
    close(fd);

    printf("1 son addr:%s\n", addr);
    addr[0] = 'a';
    addr[1] = '1';
    addr[2] = 'c';
    addr[3] = 'd';
    printf("2 son addr:%s\n", addr);

    /*ret = msync(addr, mapsize, MS_SYNC);
    if (ret == -1)
    {
      perror("msync fail");
      exit(0);
    }*/

    ret = munmap(addr, mapsize);
    if (ret == -1)
    {
      perror("munmap fail");
      exit(0);
    }

    printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
    exit(0);
  }
  else
  {
    sleep(2);

    char *addr = mmap(0, mapsize, PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
    if (addr == MAP_FAILED)
    {
      perror("mmap fail");
      return;
    }
    close(fd);

    printf("addr:%s\n", addr);

    /*ret = msync(addr, mapsize, MS_SYNC);
    if (ret == -1)
    {
      perror("msync fail");
      return;
    }*/

    ret = munmap(addr, mapsize);
    if (ret == -1)
    {
      perror("munmap fail");
      return;
    }

    ret = shm_unlink(SHM_NAME);
    if (ret == -1)
    {
      printf("shm_unlink fail\n");
      return;
    }

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
  do1();

  return 0;
}