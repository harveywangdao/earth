#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>

//pstree
void do1()
{
	pid_t pid;
	pid = fork();
	if (pid == -1)
	{
		printf("fork fail\n");
		return;
	}
	else if (pid == 0)
	{
		printf("son1 start, pid = %d, ppid = %d\n", getpid(), getppid());
		printf("son1 end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
		int status = 0;
		int ret = waitpid(pid, &status, 0);
		if (ret == -1)
		{
			printf("son1 ret = %d, status = %d\n", ret, status);
			return;
		}

		printf("son1 ret = %d, status = %d\n", ret, status);
	}
}

void do2()
{
	pid_t pid;
	pid = fork();
	if (pid == -1)
	{
		printf("fork fail\n");
		return;
	}
	else if (pid == 0)
	{
		printf("son2 start, pid = %d, ppid = %d\n", getpid(), getppid());

		char *args[] = {"ls", "-l", NULL};
		char *envp[] = {"PATH=/bin", NULL};
		int ret = execve("/bin/ls", args, envp);
		if (ret == -1)
		{
			printf("execve fail\n");
		}

		printf("son2 end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
		int status = 0;
		int ret = waitpid(pid, &status, 0);
		if (ret == -1)
		{
			printf("son2 ret = %d, status = %d\n", ret, status);
			return;
		}

		printf("son2 ret = %d, status = %d\n", ret, status);
	}
}

void do3()
{
	printf("system start\n");

	int ret = system("ping www.baidu.com -c 2");

	printf("system end ret:%d\n", ret);
}

void orphan()
{
	pid_t pid;
	pid = fork();
	if (pid == -1)
	{
		printf("fork fail\n");
		return;
	}
	else if (pid == 0)
	{
		printf("orphan father start, pid = %d, ppid = %d\n", getpid(), getppid());

		pid_t pid2;
		pid2 = fork();
		if (pid2 == -1)
		{
			printf("fork fail\n");
			return;
		}
		else if (pid2 == 0)
		{
			printf("orphan start, pid = %d, ppid = %d\n", getpid(), getppid());
			sleep(2);
			printf("orphan end, pid = %d, ppid = %d\n", getpid(), getppid());
			exit(0);
		}
		else
		{
			printf("orphan father end, pid = %d, ppid = %d\n", getpid(), getppid());
			exit(0);
		}
	}
	else
	{
		int status = 0;
		int ret = waitpid(pid, &status, 0);
		if (ret == -1)
		{
			printf("orphan father ret = %d, status = %d\n", ret, status);
			return;
		}

		printf("orphan father ret = %d, status = %d\n", ret, status);
	}
}


void zombie()
{
	pid_t pid;
	pid = fork();
	if (pid == -1)
	{
		printf("fork fail\n");
		return;
	}
	else if (pid == 0)
	{
		printf("zombie start, pid = %d, ppid = %d\n", getpid(), getppid());
		printf("zombie end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
		sleep(2);
		system("ps -o pid,ppid,state,tty,command");
	}
}

int main(int argc, char const *argv[])
{
	pid_t pid, ppid;

	pid = getpid();
	ppid = getppid();

	printf("main process pid = %d, ppid = %d\n", pid, ppid);

	do1();
	do2();
	do3();
	orphan();
	sleep(3);
	zombie();

	printf("Main process Goodbye\n");
	return 0;
}