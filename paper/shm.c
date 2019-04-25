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

#define IPC_PATH "/home/thomas/golang/src/github.com/harveywangdao/earth/paper/shm.c"

void do1(key_t key)
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
		printf("son start, pid = %d, ppid = %d\n", getpid(), getppid());

		int ret = 0;
		int flag = IPC_CREAT | IPC_EXCL;
		int shmid = shmget(key, 512, flag | 0666);
		if (shmid == -1)
		{
			perror("shmget fail\n");
		}
		else
		{
			struct shmid_ds sds;
			ret = shmctl(shmid, IPC_STAT, &sds);
			if (ret == -1)
			{
				perror("shmctl fail\n");
			}
			printf("%d %ld %ld %ld\n", sds.shm_cpid, sds.shm_segsz, sds.shm_ctime, sds.shm_nattch);

			char *addr = shmat(shmid, 0, 0); //SHM_RDONLY
			if (addr == (char*)(-1))
			{
				perror("shmat fail\n");
			}

			ret = shmctl(shmid, IPC_STAT, &sds);
			if (ret == -1)
			{
				perror("shmctl fail\n");
			}
			printf("%d %ld %ld %ld\n", sds.shm_cpid, sds.shm_segsz, sds.shm_ctime, sds.shm_nattch);

			char str[] = "share mem";
			memcpy(addr, str, sizeof(str));

			ret = shmdt(addr);
			if (ret == -1)
			{
				perror("shmdt fail\n");
			}

			ret = shmctl(shmid, IPC_STAT, &sds);
			if (ret == -1)
			{
				perror("shmctl fail\n");
			}
			printf("%d %ld %ld %ld\n", sds.shm_cpid, sds.shm_segsz, sds.shm_ctime, sds.shm_nattch);

			/*ret = shmctl(shmid, IPC_RMID, NULL);
			if (ret == -1)
			{
				perror("shmctl fail\n");
			}*/
		}

		printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
		sleep(2);

		int ret = 0;
		int flag = IPC_CREAT;
		int shmid = shmget(key, 512, flag | 0666);
		if (shmid == -1)
		{
			perror("shmget fail\n");
		}
		else
		{
			struct shmid_ds sds;
			ret = shmctl(shmid, IPC_STAT, &sds);
			if (ret == -1)
			{
				perror("shmctl fail\n");
			}
			printf("%d %ld %ld %ld\n", sds.shm_cpid, sds.shm_segsz, sds.shm_ctime, sds.shm_nattch);

			char *addr = shmat(shmid, 0, 0); //SHM_RDONLY
			if (addr == (char*)(-1))
			{
				perror("shmat fail\n");
			}

			printf("addr:%s\n", addr);

			ret = shmctl(shmid, IPC_STAT, &sds);
			if (ret == -1)
			{
				perror("shmctl fail\n");
			}
			printf("%d %ld %ld %ld\n", sds.shm_cpid, sds.shm_segsz, sds.shm_ctime, sds.shm_nattch);

			ret = shmdt(addr);
			if (ret == -1)
			{
				perror("shmdt fail\n");
			}
			//printf("after shmdt addr:%s\n", addr);

			ret = shmctl(shmid, IPC_STAT, &sds);
			if (ret == -1)
			{
				perror("shmctl fail\n");
			}
			printf("%d %ld %ld %ld\n", sds.shm_cpid, sds.shm_segsz, sds.shm_ctime, sds.shm_nattch);

			ret = shmctl(shmid, IPC_RMID, NULL);
			if (ret == -1)
			{
				perror("shmctl fail\n");
			}
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
	key_t key;
	key = ftok(IPC_PATH, 'c');
	if (key == -1)
	{
		printf("ftok fail\n");
		return -1;
	}

	printf("key:%d\n", key);

	do1(key);

	return 0;
}