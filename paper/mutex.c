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

//gcc -o app mutex.c -lpthread

int count = 0;
int count2 = 3;
pthread_mutex_t mutex;
int running = 1;

void *task1(void *arg)
{
  while(running)
  {
    pthread_mutex_lock(&mutex);
    count++;
    count2++;
    printf("task1 count:%d;count2:%d\n", count, count2);

    if (count2 - count != 3)
    {
      printf("task1 mutex fail\n");
      exit(1);
    }

    pthread_mutex_unlock(&mutex);
    usleep(4);
  }
}

void *task2(void *arg)
{
  while(running)
  {
    pthread_mutex_lock(&mutex);
    count--;
    count2--;
    printf("task2 count:%d;count2:%d\n", count, count2);

    if (count2 - count != 3)
    {
      printf("task2 mutex fail\n");
      exit(1);
    }

    pthread_mutex_unlock(&mutex);
    usleep(1);
  }
}

void do1()
{
  //PTHREAD_RECURSIVE_MUTEX_INITIALIZER_NP PTHREAD_ERRORCHECK_MUTEX_INITIALIZER_NP
  //pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER; 
  //pthread_mutexattr_t mutexattr;
  //pthread_mutex_trylock();
  //struct timespec tsc;
  //tsc.tv_sec = 1;
  //tsc.tv_nsec = 5000000;
  //pthread_mutex_timedlock(&mutex, &tsc);
  //clock_gettime(CLOCK_REALTIME, &tsc);//CLOCK_MONOTONIC

  pthread_t pt1;
  pthread_t pt2;
  int ret = -1;

  pthread_mutex_init(&mutex, NULL);

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

  pthread_mutex_destroy(&mutex);
}

int main(int argc, char const *argv[])
{
	do1();

	return 0;
}