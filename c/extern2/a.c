#include <stdio.h>

int num = 10;
static int num2 = 10;

//int n1 = 20;
int n1;

// extern void do2();
void do2(); // 默认带extern

//void do1();
extern void do1();

void do1()
{
  num++;
  printf("a.c num = %d, num2 = %d\n", num, num2);
  do2();

  void do3();
  do3();

  n1++;
  printf("a.c n1 = %d\n", n1);

  void do4();
  do4();
}

int main(int argc, char const *argv[])
{
  do1();
  return 0;
}
