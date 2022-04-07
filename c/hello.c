#include <stdio.h>

/*
gcc -E hello.c -o hello.i
gcc -S hello.i -o hello.s
gcc -c hello.s -o hello.o
gcc hello.o -o hello
*/
int main(int argc, char const *argv[])
{
  printf("%s\n", "hello c");
  return 0;
}
