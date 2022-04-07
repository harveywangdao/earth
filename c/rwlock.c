#include <stdio.h>
#include <pthread.h>
#include <unistd.h>
#include <sched.h>
#include <semaphore.h>

//gcc -o app rwlock.c -lpthread

int count;
pthread_rwlock_t rw;//PTHREAD_RWLOCK_INITIALIZER
int running = 1;

void *task1(void *arg)
{
  while(running)
  {
    pthread_rwlock_wrlock(&rw);
    //pthread_rwlock_trywrlock(&rw);
    //pthread_rwlock_timedwrlock(&rw, &time);
    count++;
    printf("task1 count:%d\n", count);
    pthread_rwlock_unlock(&rw);
    usleep(4);
  }
}

void *task2(void *arg)
{
  while(running)
  {
    pthread_rwlock_wrlock(&rw);
    count--;
    printf("task2 count:%d\n", count);
    pthread_rwlock_unlock(&rw);
    usleep(1);
  }
}

void *task3(void *arg)
{
  while(running)
  {
    pthread_rwlock_rdlock(&rw);
    //pthread_rwlock_tryrdlock(&rw);
    //pthread_rwlock_timedrdlock(&rw, &time);

    printf("task3 count:%d\n", count);

    pthread_rwlock_unlock(&rw);
    usleep(1);
  }
}

void do1()
{
  pthread_t pt1;
  pthread_t pt2;
  pthread_t pt3;
  int ret = -1;

  //pthread_rwlockattr_t attr;
  pthread_rwlock_init(&rw, NULL);

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

  sleep(1);

  running = 0;

  pthread_join(pt1, NULL);
  pthread_join(pt2, NULL);
  pthread_join(pt3, NULL);

  pthread_rwlock_destroy(&rw);
}

int main(int argc, char const *argv[])
{
	do1();
	return 0;
}