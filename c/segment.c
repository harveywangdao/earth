#include <stdio.h>

int b = 0;
int c = 1;
int a;

static int d;
static int e = 0;
static int f = 2;

//int arr[100];

int test()
{
  int g;
  int h = 0;
  int i = 3;
  
  static int j;
  static int k = 0;
  static int l = 4;
  
  printf("&a = %p\n", &a);
  printf("&b = %p\n", &b);
  printf("&c = %p\n", &c);

  printf("&d = %p\n", &d);
  printf("&e = %p\n", &e);
  printf("&f = %p\n", &f);

  printf("&g = %p\n", &g);
  printf("&h = %p\n", &h);
  printf("&i = %p\n", &i);

  printf("&j = %p\n", &j);
  printf("&k = %p\n", &k);
  printf("&l = %p\n", &l);
  
  //printf("&arr[0] = %p\n", &arr[0]);
  //printf("&arr[1] = %p\n", &arr[1]);
  //printf("&arr[2] = %p\n", &arr[2]);
  //printf("&arr[99] = %p\n", &arr[99]);

  printf("\nbss:\n");
  printf("&a = %p\n", &a);
  printf("&b = %p\n", &b);
  printf("&d = %p\n", &d);
  printf("&e = %p\n", &e);
  printf("&j = %p\n", &j);
  printf("&k = %p\n", &k);

  printf("\ndata:\n");
  printf("&c = %p\n", &c);
  printf("&f = %p\n", &f);
  printf("&l = %p\n", &l);
}

int main()
{
  test();
  return 0;
}
