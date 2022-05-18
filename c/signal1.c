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
#include <sys/timerfd.h>
#include <signal.h>
#include <setjmp.h>
#include <sys/signalfd.h>
#include <sys/eventfd.h>

static sigjmp_buf jmpbuf;
//static jmp_buf jmpbuf;
static volatile sig_atomic_t canjump;

void pr_mask(const char *str)
{
  sigset_t sigset;
  int errno_save;

  errno_save = errno;

  sigprocmask(0, NULL, &sigset);
  printf("%s", str);

  if (sigismember(&sigset, SIGINT)) printf("SIGINT ");
  if (sigismember(&sigset, SIGQUIT)) printf("SIGQUIT ");
  if (sigismember(&sigset, SIGUSR1)) printf("SIGUSR1 ");
  if (sigismember(&sigset, SIGALRM)) printf("SIGALRM ");

  printf("\n");
  errno = errno_save;
}

static void sig_usr1(int signo)
{
  if (canjump == 0)
    return;

  pr_mask("sig_usr1 start: ");

  alarm(3);

  time_t starttime = time(NULL);
  while(1)
  {
    if (time(NULL) > starttime+5)
      break;
  }

  pr_mask("sig_usr1 done: ");
  canjump = 0;

  siglongjmp(jmpbuf, 1);
  //longjmp(jmpbuf, 1);
}

static void sig_alrm(int signo)
{
  pr_mask("sig_alrm start: ");
}

void do1()
{
  signal(SIGUSR1, sig_usr1);
  signal(SIGALRM, sig_alrm);

  if (sigsetjmp(jmpbuf, 1))
  //if (setjmp(jmpbuf))
  {
    pr_mask("exit: ");
    exit(EXIT_SUCCESS);
  }

  canjump = 1;
  while(1);
}

int main(int argc, char const *argv[])
{
  printf("pid: %d\n", getpid());
  do1();
  return 0;
}
