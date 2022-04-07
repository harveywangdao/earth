#include <stdio.h>
#include <pthread.h>
#include <unistd.h>
#include <sched.h>
#include <semaphore.h>
#include <string.h> 

//gcc -o app cond.c -lpthread

int count;
pthread_mutex_t mutex;
pthread_cond_t cond;//PTHREAD_COND_INITIALIZER
int running = 1;
//pthread_cond_timedwait(&cond, &mutex, &time);

void *task1(void *arg)
{
  int ret;
  while(running)
  {
    pthread_mutex_lock(&mutex);
    printf("task1 pthread_cond_wait start\n");
    pthread_cond_wait(&cond, &mutex);
    printf("task1 pthread_cond_wait end\n");
    pthread_mutex_unlock(&mutex);
  }

  printf("task1 exit\n");
}

void *task2(void *arg)
{
  int ret;
  while(running)
  {
    pthread_mutex_lock(&mutex);
    printf("task2 pthread_cond_wait start\n");
    pthread_cond_wait(&cond, &mutex);
    printf("task2 pthread_cond_wait end\n");
    pthread_mutex_unlock(&mutex);
  }
  
  printf("task2 exit\n");
}

void *task3(void *arg)
{
  int ret;
  while(running)
  {
    printf("task3 pthread_cond_signal start\n");
    pthread_cond_signal(&cond);
    printf("task3 pthread_cond_signal end\n");

    sleep(1);

    printf("task3 pthread_cond_broadcast start\n");
    pthread_cond_broadcast(&cond);
    printf("task3 pthread_cond_broadcast end\n");

    sleep(1);
  }

  printf("task3 exit\n");
}

void do1()
{
  pthread_t pt1;
  pthread_t pt2;
  pthread_t pt3;
  int ret = -1;

  //pthread_condattr_t attr;
  pthread_mutex_init(&mutex, NULL);
  pthread_cond_init(&cond, NULL);

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

  sleep(3);

  pthread_cond_broadcast(&cond);

  pthread_join(pt1, NULL);
  pthread_join(pt2, NULL);
  pthread_join(pt3, NULL);

  pthread_cond_destroy(&cond);
  pthread_mutex_destroy(&mutex);
}

void *taska(void *arg)
{
  while(running)
  {
    pthread_mutex_lock(&mutex);

    pthread_cond_wait(&cond, &mutex);

    printf("taska count:%d\n", count);

    pthread_mutex_unlock(&mutex);
  }

  printf("taska exit\n");
}

void *taskb(void *arg)
{
  while(running)
  {
    pthread_mutex_lock(&mutex);

    pthread_cond_wait(&cond, &mutex);

    printf("taskb count:%d\n", count);

    pthread_mutex_unlock(&mutex);
  }
  
  printf("taskb exit\n");
}

void *taskc(void *arg)
{
  int broadcast;

  while(running)
  {
    broadcast = 0;

    pthread_mutex_lock(&mutex);
    count++;

    if (count % 17 == 0)
    {
      broadcast = count / 17;
    }

    pthread_mutex_unlock(&mutex);

    if (broadcast)
    {
      if (broadcast % 2 == 0)
      {
        pthread_cond_broadcast(&cond);
      }
      else
      {
        pthread_cond_signal(&cond);
      }
    }

    usleep(10000);
  }

  printf("taskc exit\n");
}

void do2()
{
  pthread_t pt1;
  pthread_t pt2;
  pthread_t pt3;
  int ret = -1;

  pthread_mutex_init(&mutex, NULL);
  pthread_cond_init(&cond, NULL);

  ret = pthread_create(&pt1, NULL, (void*)taska, NULL);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  ret = pthread_create(&pt2, NULL, (void*)taskb, NULL);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  ret = pthread_create(&pt3, NULL, (void*)taskc, NULL);
  if (ret != 0)
  {
    printf("create thread fail\n");
    return;
  }

  sleep(4);

  running = 0;

  sleep(3);

  pthread_cond_broadcast(&cond);

  pthread_join(pt1, NULL);
  pthread_join(pt2, NULL);
  pthread_join(pt3, NULL);

  pthread_cond_destroy(&cond);
  pthread_mutex_destroy(&mutex);
}

int main(int argc, char const *argv[])
{
  //do1();
	do2();

	return 0;
}