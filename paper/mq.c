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

#define IPC_PATH "/home/thomas/golang/src/github.com/harveywangdao/earth/paper/mq.c"

struct msgdata
{
	long mtype;
	char mtext[512];
};

void do1(key_t key)
{
	int ret = 0;
	int flag = IPC_CREAT | IPC_EXCL;

	int msqid = msgget(key, flag | 0666);
	if (msqid == -1)
	{
		perror("msgget fail\n");
	}
	else
	{
		struct msgdata sdata;
		char msg[] = "I am msgqueue";

		sdata.mtype = 21;
		memcpy(sdata.mtext, msg, sizeof(msg));

		ret = msgsnd(msqid, &sdata, sizeof(msg), 0/*IPC_NOWAIT*/);
		if (ret == -1)
		{
			printf("msgsnd fail\n");
		}
		else
		{
			struct msqid_ds mds;
			struct msgdata rdata;
			ret = msgctl(msqid, IPC_STAT, &mds);
			if (ret == -1)
			{
				printf("msgctl fail\n");
			}
			printf("%ld %ld %ld\n", mds.msg_stime, mds.msg_rtime, mds.msg_qnum);

			int nbytes = msgrcv(msqid, &rdata, sizeof(rdata.mtext), 21, 0);
			if (nbytes == -1)
			{
				printf("msgrcv fail\n");
			}
			printf("nbytes:%d:%s\n", nbytes, rdata.mtext);

			ret = msgctl(msqid, IPC_STAT, &mds);
			if (ret == -1)
			{
				printf("msgctl fail\n");
			}
			printf("%ld %ld %ld\n", mds.msg_stime, mds.msg_rtime, mds.msg_qnum);

			ret = msgctl(msqid, IPC_RMID, NULL);
			if (ret == -1)
			{
				printf("msgctl fail\n");
			}
		}
	}
}

void do2(key_t key)
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

		int ret = 0;
		int flag = IPC_CREAT;// | IPC_EXCL;
		int msqid = msgget(key, flag | 0666);
		if (msqid == -1)
		{
			perror("msgget fail\n");
		}
		else
		{
			struct msgdata sdata;
			char msg[] = "I am msgqueue";

			sdata.mtype = 21;
			memcpy(sdata.mtext, msg, sizeof(msg));

			ret = msgsnd(msqid, &sdata, sizeof(msg), 0/*IPC_NOWAIT*/);
			if (ret == -1)
			{
				printf("msgsnd fail\n");
			}
			
			struct msqid_ds mds;
			ret = msgctl(msqid, IPC_STAT, &mds);
			if (ret == -1)
			{
				printf("msgctl fail\n");
			}
			printf("%ld %ld %ld\n", mds.msg_stime, mds.msg_rtime, mds.msg_qnum);

			/*ret = msgctl(msqid, IPC_RMID, NULL);
			if (ret == -1)
			{
				printf("msgctl fail\n");
			}*/
		}

		printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
		sleep(2);
		int ret = 0;
		int flag = IPC_CREAT;// | IPC_EXCL;
		int msqid = msgget(key, flag | 0666);
		if (msqid == -1)
		{
			perror("msgget fail\n");
		}
		else
		{
			struct msqid_ds mds;
			struct msgdata rdata;
			memset(&rdata, 0, sizeof(rdata));
			ret = msgctl(msqid, IPC_STAT, &mds);
			if (ret == -1)
			{
				printf("msgctl fail\n");
			}
			printf("%ld %ld %ld\n", mds.msg_stime, mds.msg_rtime, mds.msg_qnum);

			int nbytes = msgrcv(msqid, &rdata, sizeof(rdata.mtext), 21, IPC_NOWAIT);
			if (nbytes == -1)
			{
				perror("msgrcv fail\n");
			}
			printf("nbytes:%d:%s\n", nbytes, rdata.mtext);

			ret = msgctl(msqid, IPC_STAT, &mds);
			if (ret == -1)
			{
				printf("msgctl fail\n");
			}
			printf("%ld %ld %ld\n", mds.msg_stime, mds.msg_rtime, mds.msg_qnum);

			ret = msgctl(msqid, IPC_RMID, NULL);
			if (ret == -1)
			{
				printf("msgctl fail\n");
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
	key = ftok(IPC_PATH, 'a');
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