#include <stdio.h>

//extern int num = 20;
extern int num;
extern int num2; // 可以申明但是不能使用

int n1 = 10;

void do2()
{
  // printf("b.c num = %d, num2 = %d\n", num, num2);
  printf("b.c do2 num = %d\n", num);
}

void do3()
{
  printf("b.c do3 num = %d\n", num);
}

void do4()
{
  n1++;
  printf("b.c do4 n1 = %d\n", n1);
}