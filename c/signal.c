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
#include <signal.h>
#include <setjmp.h>
#include <sys/signalfd.h>

#define handle_error(msg) \
           do { perror(msg); exit(EXIT_FAILURE); } while (0)

typedef void (*sighandler)(int);

static void handler1(int no)
{
	printf("no = %d\n", no);
	printf("SIGABRT = %d\n", SIGABRT);
	printf("SIGALRM = %d\n", SIGALRM);
	printf("SIGINT = %d\n", SIGINT);
	printf("SIGUSR1 = %d\n", SIGUSR1);
}

void do1()
{
	sighandler sig;
	sig = signal(SIGABRT, handler1);//SIGSTOP SIGKILL
	if (sig == SIG_ERR)
	{
		perror("signal fail");
		return;
	}

	sig = signal(SIGALRM, handler1);
	if (sig == SIG_ERR)
	{
		perror("signal fail");
		return;
	}

	sig = signal(SIGINT, handler1);
	if (sig == SIG_ERR)
	{
		perror("signal fail");
		return;
	}

	sig = signal(SIGUSR1, handler1);
	if (sig == SIG_ERR)
	{
		perror("signal fail");
		return;
	}

	//kill(getpid(), SIGUSR1);
	//raise(SIGUSR1);
	//abort();
	//alarm(2);

	sleep(4);
	printf("sleep end\n");
}

void do2()
{
	sighandler sig;
	sig = signal(SIGUSR1, handler1);
	if (sig == SIG_ERR)
	{
		perror("signal fail");
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
		sleep(3);
		kill(getppid(), SIGUSR1);
		printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
		printf("pause start\n");
		pause();
		printf("pause end\n");

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

void do3()
{
	sighandler sig;
	sig = signal(SIGALRM, handler1);
	if (sig == SIG_ERR)
	{
		perror("signal fail");
		return;
	}

	printf("sleep start\n");
	sleep(2);
	printf("sleep end\n");
}

void do4()
{
	int n = 0;
	jmp_buf buf;

	int ret = setjmp(buf);

	n++;
	printf("dddd %d, %d\n",ret, n);

	if (n == 20)
	{
		exit(0);
	}

	longjmp(buf, 1);
}

void do5()
{
	sigset_t mask;
	int sfd;
	struct signalfd_siginfo fdsi;
	ssize_t s;

	sigemptyset(&mask);
	sigaddset(&mask, SIGINT);
	sigaddset(&mask, SIGQUIT);

	/* Block signals so that they aren't handled
	according to their default dispositions */

	if (sigprocmask(SIG_BLOCK, &mask, NULL) == -1)
		handle_error("sigprocmask");

	sfd = signalfd(-1, &mask, 0);
	if (sfd == -1)
		handle_error("signalfd");

	for (;;) {
		s = read(sfd, &fdsi, sizeof(struct signalfd_siginfo));
		if (s != sizeof(struct signalfd_siginfo))
			handle_error("read");

		if (fdsi.ssi_signo == SIGINT) {
			printf("Got SIGINT\n");
		} else if (fdsi.ssi_signo == SIGQUIT) {
			printf("Got SIGQUIT\n");
			exit(EXIT_SUCCESS);
		} else {
			printf("Read unexpected signal\n");
		}
	}
}

void now()
{
	time_t now = time(NULL);
	printf("now is %s\n", ctime(&now));
}

static void sig_alarm(int sig)
{
	printf("alarm: %d\n", sig);
	now();
}

void do6()
{
	signal(SIGALRM, sig_alarm);

	unsigned int n1 = alarm(7);
	printf("%d\n", n1);
	now();

	sleep(3);

	unsigned int n2 = alarm(7);
	printf("%d\n", n2);
	now();

	sleep(100);
	//sleep(100);
	printf("sleep done\n");
	now();
}

void do7()
{
	signal(SIGALRM, sig_alarm);

	printf("sleep start\n");
	now();

	sleep(2);

	printf("sleep done\n");
	now();
}

int main(int argc, char const *argv[])
{
	//do1();
	//do2();
	//do3();
	do7();
	return 0;
}