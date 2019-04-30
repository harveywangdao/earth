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
#include <mqueue.h>
#include <signal.h>

//gcc -o app mq_posix.c -lrt

#define MQ_NAME "/bigspoon"

/*union sigval
{
  int sival_int;
  void *sival_ptr;
}*/

typedef void (*sighandler)(int);

static void handler(int no)
{
  printf("no = %d\n", no);
  printf("SIGABRT = %d\n", SIGABRT);
  printf("SIGALRM = %d\n", SIGALRM);
  printf("SIGINT = %d\n", SIGINT);
  printf("SIGUSR1 = %d\n", SIGUSR1);
}

void do1()
{
  int ret = 0;

  struct mq_attr attr;
  attr.mq_maxmsg = 10;
  attr.mq_msgsize = 32;
  mqd_t fd = mq_open(MQ_NAME, O_RDWR|O_CREAT|O_EXCL, 0666, &attr);
  if (fd == -1)
  {
    printf("mq_open fail\n");
    return;
  }

  ret = mq_getattr(fd, &attr);
  if (ret == -1)
  {
    printf("mq_getattr fail\n");
    return;
  }
  printf("%ld %ld %ld %ld\n", attr.mq_flags, attr.mq_maxmsg, attr.mq_msgsize, attr.mq_curmsgs);

  struct mq_attr oattr;
  attr.mq_flags = 0;
  ret = mq_setattr(fd, &attr, &oattr);
  if (ret == -1)
  {
    printf("mq_setattr fail\n");
    return;
  }
  printf("%ld %ld %ld %ld\n", oattr.mq_flags, oattr.mq_maxmsg, oattr.mq_msgsize, oattr.mq_curmsgs);

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

    sighandler sig;
    sig = signal(SIGUSR1, handler);
    if (sig == SIG_ERR)
    {
      perror("signal fail");
      exit(0);
    }

    struct sigevent sigeve;
    sigeve.sigev_notify = SIGEV_SIGNAL;//SIGEV_NONE SIGEV_THREAD
    sigeve.sigev_signo = SIGUSR1;
    sigeve.sigev_value.sival_ptr = NULL;
    sigeve.sigev_notify_function = NULL;
    sigeve.sigev_notify_attributes = NULL;

    ret = mq_notify(fd, &sigeve);
    if (ret == -1)
    {
      printf("mq_notify fail\n");
      exit(0);
    }

    char *str = "hello mq";
    ret = mq_send(fd, str, strlen(str)+1, 1);//MQ_PRIO_MAX
    if (ret == -1)
    {
      printf("mq_send fail\n");
      exit(0);
    }
    
    ret = mq_close(fd);
    if (ret == -1)
    {
      printf("mq_close fail\n");
      return;
    }

    printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
    exit(0);
  }
  else
  {
    sleep(2);

    char buf[32];
    unsigned int prio = 1;
    ret = mq_receive(fd, buf, sizeof(buf), &prio);//MQ_PRIO_MAX
    if (ret == -1)
    {
      printf("mq_receive fail\n");
      return;
    }

    printf("buf:%s\n", buf);

    ret = mq_close(fd);
    if (ret == -1)
    {
      printf("mq_close fail\n");
      return;
    }

    ret = mq_unlink(MQ_NAME);
    if (ret == -1)
    {
      printf("mq_unlink fail\n");
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