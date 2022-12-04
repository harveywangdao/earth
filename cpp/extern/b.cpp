#include <iostream>

namespace wang
{
  int n1 = 10;
  static int n2 = 20;
}

namespace wang
{
  void do1()
  {
    std::cout << "n1: " << n1 << std::endl;
    std::cout << "n2: " << n2 << std::endl;
  }
}
