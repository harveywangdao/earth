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

//gcc -o app barrier.c -lpthread

int count = 0;
pthread_barrier_t barrier;
int running = 1;

void *task1(void *arg)
{
  while(running)
  {
    printf("task1 pthread_barrier_wait start\n");
    pthread_barrier_wait(&barrier);
    printf("task1 pthread_barrier_wait end\n");
  }
}

void *task2(void *arg)
{
  while(running)
  {
    printf("task2 pthread_barrier_wait start\n");
    pthread_barrier_wait(&barrier);
    printf("task2 pthread_barrier_wait end\n");
  }
}

void *task3(void *arg)
{
  while(running)
  {
    printf("task3 pthread_barrier_wait start\n");
    pthread_barrier_wait(&barrier);
    printf("task3 pthread_barrier_wait end\n");
    sleep(1);
  }
}

void do1()
{
  pthread_t pt1;
  pthread_t pt2;
  pthread_t pt3;
  int ret = -1;

  //pthread_barrierattr_t attr;
  pthread_barrier_init(&barrier, NULL, 3);

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

  ret = pthread_create(&pt3, NULL, (void*)task3, NULL);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  sleep(4);

  running = 0;

  sleep(1);
  
  pthread_barrier_wait(&barrier);

  pthread_join(pt1, NULL);
  pthread_join(pt2, NULL);
  pthread_join(pt3, NULL);

  pthread_barrier_destroy(&barrier);
}

int main(int argc, char const *argv[])
{
  do1();

  return 0;
}