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

int main(int argc, char const *argv[])
{
	//do1();
	//do2();
	//do3();
	do4();
	return 0;
}