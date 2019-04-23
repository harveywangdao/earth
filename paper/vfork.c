#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <sys/resource.h>
#include <errno.h>

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

void do5()
{
	pid_t pid;
	pid = fork();
	if (pid == -1)
	{
		printf("vfork fail\n");
		return;
	}
	else if (pid == 0)
	{
		printf("son start, pid = %d, ppid = %d\n", getpid(), getppid());
		sleep(2);

		int nicevalue = nice(0);
		if (nicevalue == -1 && errno != 0)
		{
			printf("nice fail\n");
			//return;
		}
		printf("son nicevalue:%d\n", nicevalue);

		int priority = getpriority(PRIO_PROCESS, getppid());
		if (priority == -1 && errno != 0)
		{
			printf("getpriority fail\n");
			//return;
		}
		printf("%d priority:%d\n", getppid(), priority);

		printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
		//getpriority & setpriority
		int priority = getpriority(PRIO_PROCESS, pid);
		if (priority == -1 && errno != 0)
		{
			printf("getpriority fail\n");
			//return;
		}
		printf("%d priority:%d\n", pid, priority);

		int ret = setpriority(PRIO_PROCESS, pid, 4);
		//int ret = setpriority(PRIO_PROCESS, pid, -2);
		if (ret == -1)
		{
			printf("setpriority fail\n");
			//return;
		}

		priority = getpriority(PRIO_PROCESS, pid);
		if (priority == -1 && errno != 0)
		{
			printf("getpriority fail\n");
			//return;
		}
		printf("%d priority:%d\n", pid, priority);

		//nice
		int nicevalue = nice(0);
		if (nicevalue == -1 && errno != 0)
		{
			printf("nice fail\n");
			//return;
		}
		printf("main process nicevalue:%d\n", nicevalue);

		nicevalue = nice(3);
		//nicevalue = nice(-3);
		if (nicevalue == -1 && errno != 0)
		{
			printf("nice fail\n");
			//return;
		}
		printf("main process nicevalue:%d\n", nicevalue);
		
		nicevalue = nice(0);
		if (nicevalue == -1 && errno != 0)
		{
			printf("nice fail\n");
			//return;
		}
		printf("main process nicevalue:%d\n", nicevalue);

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

void do6()
{
	pid_t pid;
	pid = fork();
	if (pid == -1)
	{
		printf("vfork fail\n");
		return;
	}
	else if (pid == 0)
	{
		printf("son start, pid = %d, ppid = %d\n", getpid(), getppid());
		
		char *un = getlogin();
		if (un == NULL)
		{
			printf("getlogin fail\n");
		}
		else
		{
			printf("son getlogin:%s\n", un);
		}

		printf("son getuid:%d\n", getuid());
		printf("son geteuid:%d\n", geteuid());
		printf("son getgid:%d\n", getgid());
		printf("son getegid:%d\n", getegid());

		printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
		char *un = getlogin();
		if (un == NULL)
		{
			printf("getlogin fail\n");
		}
		else
		{
			printf("getlogin:%s\n", un);
		}

		printf("getuid:%d\n", getuid());
		printf("geteuid:%d\n", geteuid());
		printf("getgid:%d\n", getgid());
		printf("getegid:%d\n", getegid());
		
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
	pid_t pid, ppid;

	pid = getpid();
	ppid = getppid();

	printf("main process pid = %d, ppid = %d\n", pid, ppid);

	//do1();
	//do2();
	//do3();
	//do4();
	//do5();
	do6();

	//printf("Main process Goodbye\n");
	return 0;
}