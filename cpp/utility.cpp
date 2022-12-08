#include <iostream>
#include <utility>
#include <algorithm>

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

void do2()
{
  //std::swap(10, 20);
  //std::swap(std::move(10), std::move(20));
}

void do3()
{
  int arr1[3] = {1,2,3};
  int arr2[3] = {10,20,30};

  std::cout << "arr1 = " << arr1 << std::endl;
  std::cout << "arr2 = " << arr2 << std::endl;
  std::cout << "&arr1 = " << &arr1 << std::endl;
  std::cout << "&arr2 = " << &arr2 << std::endl;
  std::swap(arr1, arr2);
  std::cout << "arr1 = " << arr1 << std::endl;
  std::cout << "arr2 = " << arr2 << std::endl;
  std::cout << "&arr1 = " << &arr1 << std::endl;
  std::cout << "&arr2 = " << &arr2 << std::endl;

  for (int i = 0; i < 3; ++i)
  {
    std::cout << arr1[i] << std::endl;
  }
  for (int i = 0; i < 3; ++i)
  {
    std::cout << arr2[i] << std::endl;
  }
}

void do4()
{
  int n1 = 10;
  int n2 = 10;
  std::cout << "&n1 = " << &n1 << std::endl;
  std::cout << "&n2 = " << &n2 << std::endl;
  std::swap(n1, n2);
  std::cout << "&n1 = " << &n1 << std::endl;
  std::cout << "&n2 = " << &n2 << std::endl;
}

void do5()
{
  int n = 10;
  std::exchange(n, 20);
  std::cout << "n = " << n << std::endl;
}

int main(int argc, char const *argv[])
{
  do5();
  return 0;
}
