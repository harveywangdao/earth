#include <stdio.h>
#include <pthread.h>
#include <unistd.h>
#include <sched.h>
#include <stdlib.h>
#include <semaphore.h>

//gcc -o app rwlock.c -lpthread

int count;
pthread_rwlock_t rw;//PTHREAD_RWLOCK_INITIALIZER
int running = 1;

void *task1(void *arg)
{
  sleep(2);
  while(running)
  {
    printf("task1 wrlock start\n");
    pthread_rwlock_wrlock(&rw);
    printf("task1 wrlock end\n");

    count++;

    printf("task1 count:%d\n", count);

    printf("task1 unlock start\n");
    pthread_rwlock_unlock(&rw);
    printf("task1 unlock end\n");

    abort();
    break;
  }
}

void *task2(void *arg)
{
  sleep(1);
  while(running)
  {
    printf("task2 rdlock start\n");
    pthread_rwlock_rdlock(&rw);
    printf("task2 rdlock end\n");

    printf("task2 count:%d\n", count);

    sleep(5);

    printf("task2 unlock start\n");
    pthread_rwlock_unlock(&rw);
    printf("task2 unlock end\n");
    break;
  }
}

void *task3(void *arg)
{
  int sl = (int)arg;
  printf("sleep %ds\n", sl);

  sleep(3);
  while(running)
  {
    sleep(sl);
    printf("task3 rdlock start\n");
    pthread_rwlock_rdlock(&rw);
    printf("task3 rdlock end\n");

    printf("task3 count:%d\n", count);

    sleep(5);

    printf("task3 unlock start\n");
    pthread_rwlock_unlock(&rw);
    printf("task3 unlock end\n");
    break;
  }
}

void do1()
{
  int count = 100;
  pthread_t pt1;
  pthread_t pt2;
  pthread_t pt3[count];
  int ret = -1;

  //pthread_rwlockattr_t attr;
  //pthread_rwlock_init(&rw, NULL);

  pthread_rwlockattr_t attr;
  pthread_rwlockattr_init(&attr);
  pthread_rwlockattr_setkind_np(&attr, PTHREAD_RWLOCK_PREFER_WRITER_NONRECURSIVE_NP);
  pthread_rwlock_init(&rw, &attr);

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

  for (int i = 0; i < count; ++i)
  {
    ret = pthread_create(&pt3[i], NULL, (void*)task3, (void*)i);
    if (ret != 0)
    {
      printf("create thread fail\n");
      return;
    }
  }
  //sleep(10);

  //running = 0;

  pthread_join(pt1, NULL);
  pthread_join(pt2, NULL);

  for (int i = 0; i < count; ++i)
  {
    pthread_join(pt3[i], NULL);
  }

  pthread_rwlock_destroy(&rw);
}

int main(int argc, char const *argv[])
{
  do1();
  return 0;
}