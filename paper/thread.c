#include <stdio.h>
#include <pthread.h>
#include <unistd.h>
#include <sched.h>
#include <semaphore.h>

//gcc -o app thread.c -lpthread

static int run = 1;
static int retvalue;

void *start_routine(void *arg)
{
  int *running = arg;
  printf("thread 1 start\n");

  while(*running)
  {
    printf("thread 1 running\n");
    usleep(1);
  }

  printf("thread 1 stop\n");

  retvalue = 6;
  pthread_exit((void*)&retvalue);
}

void do1()
{
  pthread_t pt;
  int ret =-1;
  int times = 3;
  int i = 0;
  int *ret_join = NULL;

  ret = pthread_create(&pt, NULL, (void*)start_routine, &run);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  usleep(1);

  for (i = 0; i < times; i++)
  {
    printf("main thread running\n");
    usleep(1);
  }

  run = 0;

  pthread_join(pt, (void*)&ret_join);

  printf("thread 1 return value:%d\n", *ret_join);
}

void do2()
{
  pthread_attr_t attr;
  struct sched_param sch;

  pthread_attr_init(&attr);

  pthread_attr_getschedparam(&attr, &sch);
  sch.sched_priority = 256;
  pthread_attr_setschedparam(&attr, &sch);

  pthread_attr_setscope(&attr, PTHREAD_SCOPE_SYSTEM);//PTHREAD_SCOPE_PROCESS

  pthread_attr_setdetachstate(&attr, PTHREAD_CREATE_JOINABLE);//PTHREAD_CREATE_DETACHED

  pthread_t pt;
  int ret =-1;
  int times = 3;
  int i = 0;
  int *ret_join = NULL;

  ret = pthread_create(&pt, &attr, (void*)start_routine, &run);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  usleep(1);

  for (i = 0; i < times; i++)
  {
    printf("main thread running\n");
    usleep(1);
  }

  run = 0;

  pthread_join(pt, (void*)&ret_join);

  printf("thread 1 return value:%d\n", *ret_join);
}

int count;
pthread_mutex_t mutex;
int running = 1;

void *producer(void *arg)
{
  while(running)
  {
    pthread_mutex_lock(&mutex);
    count++;
    printf("producer count:%d\n", count);
    pthread_mutex_unlock(&mutex);
    usleep(4);
  }
}

void *consumer(void *arg)
{
  while(running)
  {
    pthread_mutex_lock(&mutex);
    count--;
    printf("consumer count:%d\n", count);
    pthread_mutex_unlock(&mutex);
    usleep(1);
  }
}

void do3()
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

  ret = pthread_create(&pt1, NULL, (void*)producer, NULL);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  ret = pthread_create(&pt2, NULL, (void*)consumer, NULL);
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

sem_t sem;

void *producer1(void *arg)
{
  int value = 0;
  printf("producer pthread_self:%ld\n", pthread_self());

  while(running)
  {
    sem_wait(&sem);
    sem_getvalue(&sem, &value);

    count++;
    printf("producer count:%d:%d\n", count, value);
    sem_post(&sem);
    usleep(1);
  }
}

void *consumer1(void *arg)
{
  int value = 0;
  printf("consumer pthread_self:%ld\n", pthread_self());

  while(running)
  {
    sem_wait(&sem);
    sem_getvalue(&sem, &value);

    count--;
    printf("consumer count:%d:%d\n", count, value);
    sem_post(&sem);
    usleep(1);
  }
}

void do4()
{
  pthread_t pt1;
  pthread_t pt2;
  int ret = -1;

  sem_init(&sem, 0, 1);

  ret = pthread_create(&pt2, NULL, (void*)consumer1, NULL);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  ret = pthread_create(&pt1, NULL, (void*)producer1, NULL);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  usleep(1000);

  running = 0;

  ret = pthread_equal(pt1, pt2);
  printf("pthread_equal:%d\n", ret);

  printf("main pthread_self:%ld\n", pthread_self());

  pthread_join(pt1, NULL);
  pthread_join(pt2, NULL);

  sem_destroy(&sem);
}

void cleanup(void *arg)
{
  printf("cleanup pthread_self:%ld:%ld\n", pthread_self(), (long)arg);
}

void *task1(void *arg)
{
  pthread_cleanup_push(cleanup, (void *)1);
  pthread_cleanup_push(cleanup, (void *)2);

  int *running = arg;
  printf("thread 1 start\n");

  while(*running)
  {
    printf("thread 1 running\n");
    sleep(1);
  }

  printf("thread 1 stop\n");

  pthread_exit((void*)1);

  pthread_cleanup_pop(1);
  pthread_cleanup_pop(1);
}

void *task2(void *arg)
{
  pthread_cleanup_push(cleanup, (void *)3);
  pthread_cleanup_push(cleanup, (void *)4);

  int *running = arg;
  printf("thread 2 start\n");

  while(*running)
  {
    printf("thread 2 running\n");
    sleep(1);
  }

  printf("thread 2 stop\n");

  pthread_exit((void*)2);

  pthread_cleanup_pop(1);
  pthread_cleanup_pop(1);
}

void *task3(void *arg)
{
  pthread_cleanup_push(cleanup, (void *)5);
  pthread_cleanup_push(cleanup, (void *)6);

  int *running = arg;
  printf("thread 3 start\n");

  while(*running)
  {
    printf("thread 3 running\n");
    sleep(1);
  }

  pthread_cleanup_pop(1);
  pthread_cleanup_pop(1);

  printf("thread 3 stop\n");

  return ((void*)3);
}

void *task4(void *arg)
{
  pthread_cleanup_push(cleanup, (void *)7);
  pthread_cleanup_push(cleanup, (void *)8);

  int *running = arg;
  printf("thread 4 start\n");

  while(*running)
  {
    printf("thread 4 running\n");
    sleep(1);
  }

  printf("thread 4 stop\n");

  return ((void*)4);

  pthread_cleanup_pop(1);
  pthread_cleanup_pop(1);
}

void *task5(void *arg)
{
  pthread_cleanup_push(cleanup, (void *)9);
  pthread_cleanup_push(cleanup, (void *)10);

  int *running = arg;
  printf("thread 5 start\n");

  while(*running)
  {
    printf("thread 5 running\n");
    sleep(1);
  }

  printf("thread 5 stop\n");

  pthread_exit((void*)5);

  pthread_cleanup_pop(1);
  pthread_cleanup_pop(1);
}

void do5()
{
  pthread_t pt1;
  pthread_t pt2;
  pthread_t pt3;
  pthread_t pt4;
  pthread_t pt5;
  int ret = -1;
  void *ret_join = NULL;

  ret = pthread_create(&pt1, NULL, (void*)task1, &run);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  ret = pthread_create(&pt2, NULL, (void*)task2, &run);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  ret = pthread_create(&pt3, NULL, (void*)task3, &run);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  ret = pthread_create(&pt4, NULL, (void*)task4, &run);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  ret = pthread_create(&pt5, NULL, (void*)task5, &run);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  sleep(2);

  pthread_cancel(pt1);
  pthread_detach(pt5);

  sleep(2);

  run = 0;

  pthread_join(pt1, &ret_join);
  printf("thread 1 return value:%ld\n", (long)ret_join);
  pthread_join(pt2, &ret_join);
  printf("thread 2 return value:%ld\n", (long)ret_join);
  pthread_join(pt3, &ret_join);
  printf("thread 3 return value:%ld\n", (long)ret_join);
  pthread_join(pt4, &ret_join);
  printf("thread 4 return value:%ld\n", (long)ret_join);
  pthread_join(pt5, &ret_join);
  printf("thread 5 return value:%ld\n", (long)ret_join);
}

int main(int argc, char const *argv[])
{
  //do1();
  //do2();
  do3();
  //do4();
  //do5();

  return 0;
}
