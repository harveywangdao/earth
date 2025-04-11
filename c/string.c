#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <stdint.h>

void do1()
{
  char str[8];

  for (int i = 0; i < sizeof(str); ++i)
  {
    str[i] = 'A';
  }
  printf("str:%s\n", str);

  snprintf(str, sizeof(str), "123456");   //自动加'\0'
  printf("str:%s\n", str);

  memset(str, 0, sizeof(str));
  snprintf(str, sizeof(str), "1234567");
  printf("str:%s\n", str);

  memset(str, 0, sizeof(str));
  snprintf(str, sizeof(str), "12345678");   //最多拷贝sizeof(str)-1
  printf("str:%s\n", str);

  memset(str, 0, sizeof(str));
  snprintf(str, sizeof(str), "123456789");
  printf("str:%s\n", str);
}

void do2()
{
  char c1[] = "中";
  printf("sizeof(c1): %ld\n", sizeof(c1));
  for (int i = 0; i < sizeof(c1); i++)
  {
    printf("%x\n", c1[i]);
  }
}

int main(int argc, char const *argv[])
{
  do2();

  return 0;
}