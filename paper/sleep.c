#include <unistd.h>

//time ./app

int main(int argc, char const *argv[])
{
  for( int i = 0; i < 10000; i++ )
  {
      usleep(1);
      usleep(0);
  }

  return 0;
}
