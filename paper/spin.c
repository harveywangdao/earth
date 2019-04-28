#include <stdio.h>
#include <pthread.h>
#include <unistd.h>
#include <sched.h>
#include <semaphore.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <string.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <errno.h>

//gcc -o app spin.c -lpthread

int count = 0;
int count2 = 3;
pthread_spinlock_t spin;
int running = 1;

void *task1(void *arg)
{
  int ret = 0;

  while(running)
  {
    ret = pthread_spin_lock(&spin);
    if (ret != 0)
    {
      printf("pthread_spin_lock fail:%s\n", strerror(ret));
    }

    count++;
    count2++;
    printf("task1 count:%d;count2:%d\n", count, count2);

    if (count2 - count != 3)
    {
      printf("task1 spin fail\n");
      exit(1);
    }

    pthread_spin_unlock(&spin);
    usleep(4);
  }
}

void *task2(void *arg)
{
  int ret = 0;

  while(running)
  {
    ret = pthread_spin_lock(&spin);
    if (ret != 0)
    {
      printf("pthread_spin_lock fail:%s\n", strerror(ret));
    }

    count--;
    count2--;
    printf("task2 count:%d;count2:%d\n", count, count2);

    if (count2 - count != 3)
    {
      printf("task2 spin fail\n");
      exit(1);
    }

    pthread_spin_unlock(&spin);
    usleep(1);
  }
}

void do1()
{
  pthread_t pt1;
  pthread_t pt2;
  int ret = -1;

  pthread_spin_init(&spin, PTHREAD_PROCESS_PRIVATE);//PTHREAD_PROCESS_SHARED
  //pthread_spin_trylock(&spin);

  ret = pthread_create(&pt1, NULL, (void*)task1, NULL);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  ret = pthread_create(&pt2, NULL, (void*)task2, NULL);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  sleep(1);

  running = 0;

  pthread_join(pt1, NULL);
  pthread_join(pt2, NULL);

  pthread_spin_destroy(&spin);
}

int main(int argc, char const *argv[])
{
  do1();

  return 0;
}