#include <iostream>

namespace wang
{
  extern void do1();

  extern int n1;
  //extern int n2;
}

int main(int argc, char const *argv[])
{
  using namespace wang;

  std::cout << "n1: " << n1 << std::endl;
  //std::cout << "n2: " << n2 << std::endl;

  do1();
  return 0;
}
