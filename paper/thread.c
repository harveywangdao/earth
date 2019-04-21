#include <stdio.h>
#include <pthread.h>
#include <unistd.h>

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

int main()
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
    return 1;
  }

  usleep(1);

  for (i = 0; i < times; i++)
  {
    printf("main thread running\n");
    usleep(1);
  }

  run =0;

  pthread_join(pt, (void*)&ret_join);

  printf("thread 1 return value:%d\n", *ret_join);

  return 0;
}

//gcc -o app thread.c -lpthread