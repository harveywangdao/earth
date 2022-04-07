#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <string.h>

void write0(int wfd, int rfd)
{
	int retn;
	char rbuff[1024];

	close(wfd);

	printf("reading pipe\n");
	retn = read(rfd, rbuff, sizeof(rbuff)-1);
	if (retn == -1)
	{
		printf("read fail\n");
		return;
	}

	rbuff[retn] = '\0';
	printf("read pipe %d:%s\n", retn, rbuff);
}

void write1(int wfd, int rfd)
{
	char wbuff[1024];
	int retn;

	sprintf(wbuff, "%s", "hello, pipe");
	printf("writing pipe:%s\n", wbuff);
	retn = write(wfd, wbuff, strlen(wbuff)+1);
	if (retn == -1)
	{
		printf("write fail\n");
		return;
	}
	printf("write pipe done:%d\n", retn);

	char rbuff[1024];
	printf("reading pipe\n");
	retn = read(rfd, rbuff, sizeof(rbuff)-1);
	if (retn == -1)
	{
		printf("read fail\n");
		return;
	}

	rbuff[retn] = '\0';
	printf("read pipe %d:%s\n", retn, rbuff);
}

void write2(int wfd, int rfd)
{
	int retn, sum = 0;
	char wbuff[1024*64]; //pipe buffer len is 64K

	//close(rfd);

	for (int i = 0; i < 10; ++i)
	{
		//sprintf(wbuff, "%d", i);
		printf("writing pipe:%ld\n", sizeof(wbuff));
		retn = write(wfd, wbuff, sizeof(wbuff));
		if (retn == -1)
		{
			printf("write fail\n");
			return;
		}
		else if (retn == 0)
		{
			printf("write pipe done retn:%d\n", retn);
			break;
		}

		sum+=retn;
		printf("write pipe done retn:%d\n", retn);
		printf("write pipe done sum:%d\n", sum);
	}
}

void write3(int wfd, int rfd)
{
	int retn = 0;
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
		close(rfd);

		char wbuff[1024];

		sprintf(wbuff, "%s", "hello, fork");
		printf("son writing pipe:%s\n", wbuff);
		retn = write(wfd, wbuff, strlen(wbuff)+1);
		//retn = write(wfd, wbuff, 0);
		if (retn == -1)
		{
			printf("write fail\n");
			return;
		}
		printf("son write pipe done:%d\n", retn);
		//sleep(5);
		printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
		close(wfd);
		//sleep(5);

		char rbuff[1024];
		printf("main process reading pipe\n");
		retn = read(rfd, rbuff, sizeof(rbuff)-1);
		if (retn == -1)
		{
			printf("read fail\n");
			return;
		}

		rbuff[retn] = '\0';
		printf("main process read pipe %d:%s\n", retn, rbuff);

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

void write4()
{
	int fd[2];
	int *rfd = &fd[0];
	int *wfd = &fd[1];

	int ret = pipe(fd);
	if (ret == -1)
	{
		printf("pipe fail\n");
		return;
	}

	int retn;
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
		close(*rfd);

		char wbuff[1024*64*10] = "pipe-------";
		int writenum = 1024*64*10;

		while(writenum >= 0)
		{
			printf("son writing pipe:%d\n", writenum);
			retn = write(*wfd, wbuff, writenum);
			if (retn == -1)
			{
				printf("write fail\n");
				return;
			}
			else if (retn == 0)
			{
				sleep(1);
				break;
			}

			writenum -= retn;
			printf("son write pipe done:%d:%d\n", retn, writenum);
		}

		printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
		close(*wfd);
		char rbuff[1024*32];

		while(1)
		{
			printf("main process reading pipe\n");
			retn = read(*rfd, rbuff, sizeof(rbuff));
			if (retn == -1)
			{
				printf("read fail\n");
				return;
			}
			else if (retn == 0)
			{
				break;
			}

			printf("main process read pipe %d\n", retn);
		}

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
	int fd[2];
	int *rfd = &fd[0];
	int *wfd = &fd[1];

	int ret = pipe(fd);
	if (ret == -1)
	{
		printf("pipe fail\n");
		return -1;
	}

	//write0(*wfd, *rfd);
	//write1(*wfd, *rfd);
	//write2(*wfd, *rfd);
	//write3(*wfd, *rfd);
	write4();
	
	sleep(2);
	return 0;
}