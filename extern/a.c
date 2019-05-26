#include <stdio.h>
#include "b.h"

extern int add(int i, int j);
extern int sumaddone();
extern int printsum();
extern int printdog();
extern int staticf();
extern int staticf2();
extern int printsame();
int div();

//extern int same;

int a;
int a;
int same;
int big = 4;

int printbig()
{
  printf("big=%d\n", big);
}

int testret()
{
  if (1)
  {
    /* code */
  }
}

typedef void (*functype)();

functype wang()
{
  int cup = 11;

  printf("wang cup = %d\n", cup);

  void pan()
  {
    printf("pan cup = %d\n", cup);
  }

  cup = 23;

  return pan;
}

int main(int argc, char const *argv[])
{
  extern int sum;
  printf("sum = %d\n", sum);

  sumaddone();
  printf("sum = %d\n", sum);

  sum = sum + 1;
  printf("sum = %d\n", sum);
  printsum();

  printf("dog = %d\n", dog);
  dog++;
  printf("dog = %d\n", dog);
  printdog();

  printf("add(1, 2) = %d\n", add(1, 2));

  //staticf();
  //staticf2();

  printf("hello extern\n");
  //f1();
  div();

  //extern int ssum;
  //printf("ssum = %d\n", ssum);

  printf("same = %d\n", same);
  same = 999;
  printf("same = %d\n", same);
  printsame();

  extern int rabbit;

  printf("rabbit = %d\n", rabbit);
  rabbit = 888;
  printf("rabbit = %d\n", rabbit);
  int printrabbit();
  printrabbit();

  #include "c.h"

  int big = 3;
  printf("big = %d\n", big);
  big = 777;
  printf("big = %d\n", big);
  int ret = printbig();
  printf("ret = %d\n", ret);

  int pen=90;
  int paper()
  {
    printf("pen=%d\n", pen);
    printf("paper\n");
  }

  paper();

  ret = testret();
  printf("ret = %d\n", ret);

  functype f1 = wang();
  f1();

  return 0;
}
