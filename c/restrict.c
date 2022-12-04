#include <stdio.h>

// gcc -std=c99 -o app restrict.c

// gcc -std=c99 -S -o restrict1.s restrict.c
// gcc -std=c99 -S -O1 -o restrict2.s restrict.c
// gcc -std=c99 -S -O3 -o restrict2.s restrict.c
int func1(int * restrict p1, int * restrict p2)
{
  *p1 = 10;
  *p2 = 20;
  return *p1 + *p2;
}

void do1()
{
  int a = 1;
  int b = 2;
  int c = func1(&a, &b);

  printf("a = %d, b = %d, c = %d\n", a, b, c);
}

// register不能取地址
void do2()
{
  int arr[16] = {7,2,3,5,78,89,66,4,3,3,3,344,54,6,7};

  int sum = 0;

  for (int i = 0; i < 16; ++i)
  {
    register int tp = arr[i];

    //int *p = &tp; // address of register variable tp requested

    sum += tp;
  }
}

void do3()
{
  // auto n1 = 3.3; // warning: type defaults to int in declaration of n1 [-Wimplicit-int]
  auto double n2 = 3.3;
}

int main(int argc, char const *argv[])
{
  do3();
  return 0;
}