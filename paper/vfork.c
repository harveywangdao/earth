#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>

void do1()
{
	pid_t pid;
	pid = vfork();
	if (pid == -1)
	{
		printf("vfork fail\n");
		return;
	}
	else if (pid == 0)
	{
		printf("son start, pid = %d, ppid = %d\n", getpid(), getppid());
		printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
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

void do2()
{
	pid_t pid;
	pid = vfork();
	if (pid == -1)
	{
		printf("vfork fail\n");
		return;
	}
	else if (pid == 0)
	{
		printf("son start, pid = %d, ppid = %d\n", getpid(), getppid());

		char *args[] = {"ls", "-l", NULL};
		char *envp[] = {"PATH=/bin", NULL};
		int ret = execve("/bin/ls", args, envp);
		if (ret == -1)
		{
			printf("execve fail\n");
		}

		printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
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
	pid_t pid;
	int n = 0;
	pid = vfork();
	if (pid == -1)
	{
		printf("vfork fail\n");
		return;
	}
	else if (pid == 0)
	{
		printf("son start, pid = %d, ppid = %d\n", getpid(), getppid());

		sleep(2);

		printf("n = %d\n", n);

		n = 14;

		printf("n = %d\n", n);

		printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
		printf("father running\n");
		printf("n = %d\n", n);

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

void do4()
{
	int i;
	for(i = 0; i < 2; i++)
	{
		fork();
		//printf("-");
		//printf("-\n");
		printf("ppid=%d, pid=%d, i=%d   ", getppid(), getpid(), i);
		//printf("ppid=%d, pid=%d, i=%d\n", getppid(), getpid(), i);
	}

	wait(NULL);
	wait(NULL);

	return;
}

int main(int argc, char const *argv[])
{
	pid_t pid, ppid;

	pid = getpid();
	ppid = getppid();

	printf("main process pid = %d, ppid = %d\n", pid, ppid);

	//do1();
	//do2();
	do3();
	//do4();

	//printf("Main process Goodbye\n");
	return 0;
}