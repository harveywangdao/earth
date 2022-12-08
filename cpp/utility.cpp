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

void m1(int& n1, int& n2, const int& n3)
{
  std::cout << "In function: " << n1 << ' ' << n2 << ' ' << n3 << '\n';
  ++n1; // 增加存储于函数对象的 n1 副本
  ++n2; // 增加 main() 的 n2
  // ++n3; // 编译错误
}

void do6()
{
  int n1 = 1, n2 = 2, n3 = 3;
  std::function<void()> bound_f = std::bind(m1, n1, std::ref(n2), std::cref(n3));
  n1 = 10;
  n2 = 11;
  n3 = 12;
  std::cout << "Before function: " << n1 << ' ' << n2 << ' ' << n3 << '\n';
  bound_f();
  std::cout << "After function: " << n1 << ' ' << n2 << ' ' << n3 << '\n';
}

int main(int argc, char const *argv[])
{
  do6();
  return 0;
}
