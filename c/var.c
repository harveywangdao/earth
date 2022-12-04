#include <stdio.h>

int incr(int a)
{
  return a + 1;
}

int n1;
int n2 = 10;
//int n3 = incr(10);

static int n4;
static int n5 = 10;
//static int n6 = incr(10);

int n7;
int n7;
//static int n7;

static int n8;
static int n8;

int n9 = 10;
//int n9 = 20;
int n9;
int n9;

int n10 = 10;

void do1()
{
  printf("n1 = %d\n", n1);
  printf("n2 = %d\n", n2);
  //printf("n3 = %d\n", n3);
  printf("n4 = %d\n", n4);
  printf("n5 = %d\n", n5);
  //printf("n6 = %d\n", n6);
  printf("n7 = %d\n", n7);
  printf("n8 = %d\n", n8);
  printf("n9 = %d\n", n9);

  int n10 = 20;
  printf("n10 = %d\n", n10);

  {
    int n10 = 30;
    printf("n10 = %d\n", n10);
    int n11 = 30;
  }
  //printf("n11 = %d\n", n11);

  static int n12;
  printf("n12 = %d\n", n12);
}

int main(int argc, char const *argv[])
{
  do1();
  return 0;
}