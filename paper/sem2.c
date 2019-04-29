#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <string.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <errno.h>
#include <semaphore.h>
#include <limits.h>

// gcc -o app sem2.c -lpthread

/*
无名信号量：
1.同一个线程加锁解锁                                 ok
2.同进程内的多线程间，一个线程加锁，另一个线程解锁      ok
3.亲缘进程间，一个进程加锁，另一个进程解锁
4.无关进程间，一个进程加锁，另一个进程解锁
*/

/*
有名信号量：
1.同一个线程加锁解锁                                 ok
2.同进程内的多线程间，一个线程加锁，另一个线程解锁      ok
3.亲缘进程间，一个进程加锁，另一个进程解锁             ok
4.无关进程间，一个进程加锁，另一个进程解锁             ok
*/

void do1()
{
  int ret;
  /*ret = sem_unlink("/ipc_sem_cup");
  if (ret == -1)
  {
    printf("sem_unlink fail\n");
    return;
  }*/

  printf("SEM_VALUE_MAX=%d\n", SEM_VALUE_MAX);
  sem_t *sem = sem_open("/ipc_sem_cup", O_CREAT, 0666, 5);//O_EXCL
  if (sem == SEM_FAILED)
  {
    perror("sem_open fail\n");
    return;
  }

  ret = sem_wait(sem);
  if (ret == -1)
  {
    printf("sem_wait fail\n");
    return;
  }

  ret = sem_trywait(sem);
  if (ret == -1)
  {
    printf("sem_trywait fail\n");
    return;
  }

  struct timespec tsc;
  tsc.tv_sec = 1;
  tsc.tv_nsec = 5000000;
  ret = sem_timedwait(sem, &tsc);
  if (ret == -1)
  {
    printf("sem_timedwait fail\n");
    return;
  }

  ret = sem_post(sem);
  if (ret == -1)
  {
    printf("sem_post fail\n");
    return;
  }

  ret = sem_close(sem);
  if (ret == -1)
  {
    printf("sem_close fail\n");
    return;
  }

  ret = sem_unlink("/ipc_sem_cup");
  if (ret == -1)
  {
    printf("sem_unlink fail\n");
    return;
  }
}

void do2()
{
  int ret;
  sem_t sem_origin;
  sem_t *sem = &sem_origin;
  ret = sem_init(sem, 1, 5);
  if (ret == -1)
  {
    printf("sem_init fail\n");
    return;
  }

  int sem_value;
  ret = sem_getvalue(sem, &sem_value);
  if (ret == -1)
  {
    printf("sem_getvalue fail\n");
    return;
  }

  ret = sem_destroy(sem);
  if (ret == -1)
  {
    printf("sem_destroy fail\n");
    return;
  }
}

void do3()
{
  sem_t *sem = sem_open("/ipc_sem_cup", O_CREAT, 0666, 1);//O_EXCL
  if (sem == SEM_FAILED)
  {
    perror("sem_open fail\n");
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

    int value = 0;
    sem_getvalue(sem, &value);
    printf("son sem value:%d\n", value);

    sem_wait(sem);

    printf("son sleep start\n");
    sleep(2);
    printf("son sleep end\n");

    sem_getvalue(sem, &value);
    printf("son sem value:%d\n", value);

    sem_post(sem);

    sem_getvalue(sem, &value);
    printf("son sem value:%d\n", value);

    int ret = sem_close(sem);
    if (ret == -1)
    {
      printf("sem_close fail\n");
      exit(0);
    }

    printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
    exit(0);
  }
  else
  {
    int value = 0;
    sem_getvalue(sem, &value);
    printf("sem value:%d\n", value);

    sem_wait(sem);

    printf("sleep start\n");
    sleep(2);
    printf("sleep end\n");

    sem_getvalue(sem, &value);
    printf("sem value:%d\n", value);

    sem_post(sem);

    sem_getvalue(sem, &value);
    printf("sem value:%d\n", value);

    int ret = sem_close(sem);
    if (ret == -1)
    {
      printf("sem_close fail\n");
      //exit(0);
    }

    ret = sem_unlink("/ipc_sem_cup");
    if (ret == -1)
    {
      printf("sem_unlink fail\n");
      //return;
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

void do4()
{
  int ret;
  sem_t sem_origin;
  sem_t *sem = &sem_origin;
  ret = sem_init(sem, 1, 1);
  if (ret == -1)
  {
    printf("sem_init fail\n");
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

    int value = 0;
    sem_getvalue(sem, &value);
    printf("son sem value:%d\n", value);

    sem_wait(sem);

    printf("son sleep start\n");
    sleep(2);
    printf("son sleep end\n");

    sem_getvalue(sem, &value);
    printf("son sem value:%d\n", value);

    sem_post(sem);

    sem_getvalue(sem, &value);
    printf("son sem value:%d\n", value);

    ret = sem_destroy(sem);
    if (ret == -1)
    {
      printf("sem_destroy fail\n");
      exit(0);
    }

    printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
    exit(0);
  }
  else
  {
    int value = 0;
    sem_getvalue(sem, &value);
    printf("sem value:%d\n", value);

    sem_wait(sem);

    printf("sleep start\n");
    sleep(2);
    printf("sleep end\n");

    sem_getvalue(sem, &value);
    printf("sem value:%d\n", value);

    ret = sem_trywait(sem);
    if (ret == -1)
    {
      printf("sem_trywait fail\n");
    }

    sem_post(sem);

    sem_getvalue(sem, &value);
    printf("sem value:%d\n", value);

    ret = sem_destroy(sem);
    if (ret == -1)
    {
      printf("sem_destroy fail\n");
      //return;
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

void do5()
{
  
}

int main(int argc, char const *argv[])
{
  //do1();
  //do2();
  //do3();
  //do4();
  do5();

  return 0;
}