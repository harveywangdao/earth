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

int do1()
{
  int fd = open("file.txt", O_RDWR|O_APPEND|O_CREAT, 0666);
  if (fd == -1)
  {
    perror("open fail");
    return -1;
  }

  int offset = lseek(fd, 0, SEEK_SET);//SEEK_CUR SEEK_END
  if (fd == -1)
  {
    perror("lseek fail");
    return -1;
  }

  char buf[512];
  char *word = "file sys";
  memcpy(buf, word, strlen(word)+1);

  lseek(fd, 0, SEEK_SET);
  int nb = write(fd, buf, strlen(buf)+1);
  if (nb == -1)
  {
    perror("write fail");
    return -1;
  }
  
  lseek(fd, 0, SEEK_SET);
  memset(buf, 0, sizeof(buf));
  nb = read(fd, buf, sizeof(buf));
  if (nb == -1)
  {
    perror("read fail");
    return -1;
  }
  printf("1 read %d:%s\n", nb, buf);

  int dfd = dup(fd);
  if (dfd == -1)
  {
    perror("dup fail");
    return -1;
  }

  lseek(fd, 0, SEEK_SET);
  memset(buf, 0, sizeof(buf));
  nb = read(dfd, buf, sizeof(buf));
  if (nb == -1)
  {
    perror("read fail");
    return -1;
  }
  printf("2 read %d:%s\n", nb, buf);

  int d2fd = dup2(fd, 145);
  if (d2fd == -1)
  {
    perror("dup2 fail");
    return -1;
  }
  printf("d2fd = %d\n", d2fd);

  lseek(fd, 0, SEEK_SET);
  memset(buf, 0, sizeof(buf));
  nb = read(d2fd, buf, sizeof(buf));
  if (nb == -1)
  {
    perror("read fail");
    return -1;
  }
  printf("3 read %d:%s\n", nb, buf);

  //int ret = fdatasync(fd);
  //sync();
  int ret = fsync(fd);
  if (ret == -1)
  {
    perror("fsync fail");
    return -1;
  }

  close(d2fd);
  close(dfd);
  close(fd);

  int fd2 = creat("file2.txt", 0666);
  if (fd2 == -1)
  {
    perror("creat fail");
    return -1;
  }

  memcpy(buf, word, strlen(word)+1);

  lseek(fd2, 0, SEEK_SET);
  nb = write(fd2, buf, strlen(buf)+1);
  if (nb == -1)
  {
    perror("write fail");
    return -1;
  }

  lseek(fd2, 0, SEEK_SET);
  memset(buf, 0, sizeof(buf));
  nb = read(fd2, buf, sizeof(buf));
  if (nb == -1)
  {
    perror("read fail");
    //return -1;
  }
  printf("4 read %d:%s\n", nb, buf);

  struct stat st;
  ret = fstat(fd2, &st);
  if (ret == -1)
  {
    perror("fstat fail");
    return -1;
  }
  printf("mode = %d, size = %ld, mtime = %ld\n", st.st_mode, st.st_size, st.st_mtime);

  close(fd2);

  ret = unlink("file.txt");
  if (ret == -1)
  {
    perror("unlink fail");
  }

  ret = remove("file2.txt");
  if (ret == -1)
  {
    perror("remove fail");
  }

  return 0;
}

void do2()
{
  int fd = creat("file3.txt", 0666);
  if (fd == -1)
  {
    perror("creat fail");
    return;
  }

  char buf[512];
  char *word = "file sys";
  memcpy(buf, word, strlen(word)+1);

  int nb = write(fd, buf, strlen(buf)+1);
  if (nb == -1)
  {
    perror("write fail");
    return;
  }

  int ret = fcntl(fd, F_GETFL, 0);
  if (ret == -1)
  {
    perror("fcntl fail");
    return;
  }

  switch (ret & O_ACCMODE)
  {
    case O_RDONLY:
      printf("O_RDONLY\n");
      break;

    case O_WRONLY:
      printf("O_WRONLY\n");
      break;

    case O_RDWR:
      printf("O_RDWR\n");
      break;

    default:
      printf("O_UNKNOWN\n");
      break;
  }

  int val = ret;
  val |= O_SYNC;
  ret = fcntl(fd, F_SETFL, val);
  if (ret == -1)
  {
    perror("fcntl fail");
    return;
  }

  char *word2 = "O_SYNC";
  memcpy(buf, word2, strlen(word2)+1);

  nb = write(fd, buf, strlen(buf)+1);
  if (nb == -1)
  {
    perror("write fail");
    return;
  }

  close(fd);
  remove("file3.txt");
}

void do3()
{
  //ioctl();
}

void do4()
{
  int fd = open("file.txt", O_RDWR|O_CREAT|O_TRUNC, 0666);
  if (fd == -1)
  {
    perror("open fail");
    return;
  }

  char buf[512];
  char *str = "11112222";
  memcpy(buf, str, strlen(str)+1);

  lseek(fd, 0, SEEK_SET);
  int nb = write(fd, buf, strlen(buf)+1);
  if (nb == -1)
  {
    perror("write fail");
    return;
  }

  int fd2 = dup(fd);
  if (fd2 == -1)
  {
    perror("dup fail");
    return;
  }
  
  lseek(fd, 0, SEEK_SET);
  lseek(fd2, 0, SEEK_SET);

  memset(buf, 0, sizeof(buf));
  nb = read(fd, buf, 4);
  if (nb == -1)
  {
    perror("read fail");
    return;
  }
  printf("1 read %d:%s\n", nb, buf);

  memset(buf, 0, sizeof(buf));
  nb = read(fd2, buf, 4);
  if (nb == -1)
  {
    perror("read fail");
    return;
  }
  printf("2 read %d:%s\n", nb, buf);

  close(fd);
  close(fd2);
  remove("file.txt");
}

void do5()
{
  int fd = open("file.txt", O_RDWR|O_CREAT|O_TRUNC, 0666);
  if (fd == -1)
  {
    perror("open fail");
    return;
  }

  char buf[512];
  char *str = "11112222";
  memcpy(buf, str, strlen(str)+1);

  lseek(fd, 0, SEEK_SET);
  int nb = write(fd, buf, strlen(buf)+1);
  if (nb == -1)
  {
    perror("write fail");
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
    
    lseek(fd, 0, SEEK_SET);
    memset(buf, 0, sizeof(buf));
    nb = read(fd, buf, 4);
    if (nb == -1)
    {
      perror("read fail");
    }
    printf("son read %d:%s\n", nb, buf);
    
    close(fd);
    int ret = remove("file.txt");
    if (ret == -1)
    {
      perror("remove fail");
    }
    printf("ret = %d\n", ret);

    printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
    exit(0);
  }
  else
  {
    lseek(fd, 0, SEEK_SET);
    memset(buf, 0, sizeof(buf));
    nb = read(fd, buf, 4);
    if (nb == -1)
    {
      perror("read fail");
      //return;
    }
    printf("read %d:%s\n", nb, buf);

    close(fd);
    int ret = remove("file.txt");
    if (ret == -1)
    {
      perror("remove fail");
    }
    printf("ret = %d\n", ret);

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

void readlock(int fd)
{
  struct flock lock;
  lock.l_type = F_RDLCK;
  lock.l_whence = SEEK_SET;
  lock.l_start = 0;
  lock.l_len = 0;

  //ret = fcntl(fd, F_SETLKW, &lock);
  int ret = fcntl(fd, F_SETLK, &lock);
  if (ret == -1)
  {
    perror("readlock fail");
    return;
  }

  return;
}

void writelock(int fd)
{
  struct flock lock;
  lock.l_type = F_WRLCK;
  lock.l_whence = SEEK_SET;
  lock.l_start = 0;
  lock.l_len = 0;

  //ret = fcntl(fd, F_SETLKW, &lock);
  int ret = fcntl(fd, F_SETLK, &lock);
  if (ret == -1)
  {
    perror("writelock fail");
    return;
  }

  return;
}

void unlock(int fd)
{
  struct flock lock;
  lock.l_type = F_UNLCK;
  lock.l_whence = SEEK_SET;
  lock.l_start = 0;
  lock.l_len = 0;

  //ret = fcntl(fd, F_SETLKW, &lock);
  int ret = fcntl(fd, F_SETLK, &lock);
  if (ret == -1)
  {
    perror("unlock fail");
    return;
  }

  return;
}

void getlock(int fd, int type)
{
  struct flock lock;
  lock.l_type = type;
  lock.l_whence = SEEK_SET;
  lock.l_start = 0;
  lock.l_len = 0;
  lock.l_pid = 0;

  int ret = fcntl(fd, F_GETLK, &lock);
  if (ret == -1)
  {
    perror("getlock fail");
    //return;
  }

  printf("%d %d %ld %ld %d\n", lock.l_type, lock.l_whence, lock.l_start, lock.l_len, lock.l_pid);

  return;
}

void do6()
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

  lseek(fd, 0, SEEK_SET);
  int nb = write(fd, buf, strlen(buf)+1);
  if (nb == -1)
  {
    perror("write fail");
    return;
  }

  //printf("F_RDLCK = %d, F_WRLCK = %d, F_UNLCK = %d\n", F_RDLCK, F_WRLCK, F_UNLCK);

  //readlock(fd);
  //getlock(fd, F_RDLCK);
  //getlock(fd, F_WRLCK);
  getlock(fd, F_UNLCK);
  
  //unlock(fd);
  //getlock(fd, F_UNLCK);

  //getlock(fd);
  //writelock(fd);
  //getlock(fd);

  lseek(fd, 0, SEEK_SET);
  memset(buf, 0, sizeof(buf));
  nb = read(fd, buf, 4);
  if (nb == -1)
  {
    perror("read fail");
  }
  printf("read %d:%s\n", nb, buf);

  close(fd);
  ret = remove("file.txt");
  if (ret == -1)
  {
    perror("remove fail");
  }
}

int main(int argc, char const *argv[])
{
  //do1();
  //do2();
  //do3();
  //do4();
  //do5();
  do6();
}