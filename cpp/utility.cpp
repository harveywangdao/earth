#include <iostream>
#include <utility>

void do1()
{
  int a{10};
  int b(20);

  std::cout << "a = " << a << std::endl;
  std::cout << "b = " << b << std::endl;
  std::swap(a, b);
  std::cout << "a = " << a << std::endl;
  std::cout << "b = " << b << std::endl;

  int c = 30;
  int d{};
  int e();
  std::cout << "c = " << c << std::endl;
  std::cout << "d = " << d << std::endl;
  std::cout << "e = " << e << std::endl;
}

int main(int argc, char const *argv[])
{
  do1();
  return 0;
}