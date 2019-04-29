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

#define IPC_PATH "/home"

union semun
{
	int val;
	struct semid_ds *buf;
	unsigned short *array;
};

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
		int semid = semget(key, 1, flag | 0666);
		if (semid == -1)
		{
			perror("semget fail\n");
		}
		else
		{
			union semun sun;
			sun.val = 10;

			ret = semctl(semid, 0, SETVAL, sun);
			if (ret == -1)
			{
				printf("semctl fail\n");
			}

			printf("1 nsem = %d\n", semctl(semid, 0, GETVAL));

			struct sembuf ops[2];
			ops[0].sem_num = 0;
			ops[0].sem_op = -1;
			ops[0].sem_flg = 0;

			ret = semop(semid, ops, 1);
			if (ret == -1)
			{
				printf("semop fail\n");
			}
			printf("2 nsem = %d\n", semctl(semid, 0, GETVAL));

			for (int i = 0; i < 10; ++i)
			{
				ret = semop(semid, ops, 1);
				if (ret == -1)
				{
					printf("semop fail\n");
					break;
				}
				printf("3 nsem = %d\n", semctl(semid, 0, GETVAL));
			}
			
			struct semid_ds sds;
			sun.buf = &sds;
			ret = semctl(semid, 0, IPC_STAT, sun);
			if (ret == -1)
			{
				perror("semctl fail\n");
			}
			printf("%ld %ld %ld\n", sds.sem_nsems, sds.sem_otime, sds.sem_ctime);

			/*ret = semctl(semid, 0, IPC_RMID);
			if (ret == -1)
			{
				printf("semctl fail\n");
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
		int semid = semget(key, 1, flag | 0666);
		if (semid == -1)
		{
			perror("semget fail\n");
		}
		else
		{
			printf("4 nsem = %d\n", semctl(semid, 0, GETVAL));

			struct sembuf ops[2];
			ops[0].sem_num = 0;
			ops[0].sem_op = 1;
			ops[0].sem_flg = 0;

			ret = semop(semid, ops, 1);
			if (ret == -1)
			{
				printf("semop fail\n");
			}
			printf("5 nsem = %d\n", semctl(semid, 0, GETVAL));

			sleep(2);

			for (int i = 0; i < 20; ++i)
			{
				ret = semop(semid, ops, 1);
				if (ret == -1)
				{
					printf("semop fail\n");
					break;
				}
				printf("6 nsem = %d\n", semctl(semid, 0, GETVAL));
			}

			struct semid_ds sds;
			union semun sun;
			sun.buf = &sds;
			
			ret = semctl(semid, 0, IPC_STAT, sun);
			if (ret == -1)
			{
				printf("semctl fail\n");
			}
			printf("%ld %ld %ld\n", sds.sem_nsems, sds.sem_otime, sds.sem_ctime);

			ret = semctl(semid, 0, IPC_RMID);
			if (ret == -1)
			{
				printf("semctl fail\n");
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

void do2(key_t key)
{
	int flag = IPC_CREAT | IPC_EXCL;
	int semid = semget(key, 1, flag | 0666);
	if (semid == -1)
	{
		perror("semget fail\n");
		return;
	}

	union semun sun;
	sun.val = 10;

	int ret = 0;
	ret = semctl(semid, 0, SETVAL, sun);
	if (ret == -1)
	{
		printf("semctl fail\n");
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

		printf("1 nsem = %d\n", semctl(semid, 0, GETVAL));

		struct sembuf ops[2];
		ops[0].sem_num = 0;
		ops[0].sem_op = -1;
		ops[0].sem_flg = 0;

		ret = semop(semid, ops, 1);
		if (ret == -1)
		{
			printf("semop fail\n");
		}
		printf("2 nsem = %d\n", semctl(semid, 0, GETVAL));

		for (int i = 0; i < 10; ++i)
		{
			ret = semop(semid, ops, 1);
			if (ret == -1)
			{
				printf("semop fail\n");
				break;
			}
			printf("3 nsem = %d\n", semctl(semid, 0, GETVAL));
		}
		
		struct semid_ds sds;
		sun.buf = &sds;
		ret = semctl(semid, 0, IPC_STAT, sun);
		if (ret == -1)
		{
			perror("semctl fail\n");
		}
		printf("%ld %ld %ld\n", sds.sem_nsems, sds.sem_otime, sds.sem_ctime);

		printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
		sleep(2);

		printf("4 nsem = %d\n", semctl(semid, 0, GETVAL));

		struct sembuf ops[2];
		ops[0].sem_num = 0;
		ops[0].sem_op = 1;
		ops[0].sem_flg = 0;

		ret = semop(semid, ops, 1);
		if (ret == -1)
		{
			printf("semop fail\n");
		}
		printf("5 nsem = %d\n", semctl(semid, 0, GETVAL));

		sleep(2);

		for (int i = 0; i < 20; ++i)
		{
			ret = semop(semid, ops, 1);
			if (ret == -1)
			{
				printf("semop fail\n");
				break;
			}
			printf("6 nsem = %d\n", semctl(semid, 0, GETVAL));
		}

		struct semid_ds sds;
		union semun sun;
		sun.buf = &sds;
		
		ret = semctl(semid, 0, IPC_STAT, sun);
		if (ret == -1)
		{
			printf("semctl fail\n");
		}
		printf("%ld %ld %ld\n", sds.sem_nsems, sds.sem_otime, sds.sem_ctime);

		ret = semctl(semid, 0, IPC_RMID);
		if (ret == -1)
		{
			printf("semctl fail\n");
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
	key = ftok(IPC_PATH, 'b');
	if (key == -1)
	{
		printf("ftok fail\n");
		return -1;
	}

	printf("key:%d\n", key);

	//do1(key);
	do2(key);

	return 0;
}