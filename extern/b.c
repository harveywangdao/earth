#include <stdio.h>

static int staticf2();

int sum = 10;
static int ssum = 25;
int same = 2;

int add(int i, int j)
{
  return i + j;
}

int sumaddone()
{
  sum = sum + 1;
}

int printsum()
{
  printf("b.c sum = %d\n", sum);
}

int printdog()
{
  extern int dog;
  printf("b.c dog = %d\n", dog);
}

static int staticf()
{
  printf("b.c staticf\n");
}

int staticf2()
{
  printf("b.c staticf2\n");
}

int div()
{
  printf("b.c div\n");
}

int printsame()
{
  printf("b.c same=%d\n", same);
}

int rabbit = 3;

int printrabbit()
{
  printf("b.c rabbit=%d\n", rabbit);
}