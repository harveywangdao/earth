#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/sem.h>
#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>

union semun
{
  int val;
  struct semid_ds *buf;
  unsigned short *array;
};

int sem_init(int sem_id, int init_valve)
{
  union semun sem_union;

  sem_union.val = init_valve; //初始化信号量的值

  if( semctl(sem_id, 0, SETVAL, sem_union) )
  {
    perror("Initinal semaphore");
    return -1;
  }

  return 0;
}

int sem_v(int sem_id)
{
  struct sembuf sem_b;

  sem_b.sem_num = 0; //信号量编号
  sem_b.sem_op = 1;
  sem_b.sem_flg = SEM_UNDO;

  printf("v start\n");

  if( semop(sem_id, &sem_b, 1)==-1 )
  {
    perror("V operation");

    return -1;
  }

  printf("v end\n");

  return 0;
}

int sem_p(int sem_id)
{
  struct sembuf sem_b;

  sem_b.sem_num = 0;
  sem_b.sem_op = -1;
  sem_b.sem_flg = SEM_UNDO;

  printf("p start\n");

  if( semop(sem_id, &sem_b, 1)==-1 )
  {
    perror("P operation\n");

    return -1;
  }

  printf("p end\n");

  return 0;
}

int main()
{
  pid_t pid;

  int sem_id;

  sem_id = semget(ftok(".", 'a'), 1, 0666|IPC_CREAT );

  sem_init(sem_id, 1);  //初始化信号量的值

  printf("sem_id:%d\n", sem_id);

  pid = fork();
  if(pid == -1)
  {
    perror("fork\n");
    return -1;
  }
  else if( pid == 0) //子
  {
    printf("sem_id:%d\n", sem_id);
    printf("1 nsem = %d\n", semctl(sem_id, 0, GETVAL));

    sem_p(sem_id);
    printf("2 nsem = %d\n", semctl(sem_id, 0, GETVAL));

    printf("child progress, pid=%d\n", getpid());

    sleep(4);

    printf("son exit\n");
    //sem_v(sem_id);    
  }
  else //父
  {
    sleep(1);
    printf("3 nsem = %d\n", semctl(sem_id, 0, GETVAL));

    sem_p(sem_id); //为什么不阻塞？
    printf("4 nsem = %d\n", semctl(sem_id, 0, GETVAL));

    printf("farther pid=%d\n", getpid());

    //sem_v(sem_id);
  }

  return 0;
}
