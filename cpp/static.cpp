#include <iostream>
#include <type_traits>

using namespace std;

void do1()
{
  static_assert(true); // c++17
  static_assert(true, "this is true");

  const int a = 3;
  const int b = 3;
  static_assert((a == b), "(a == b)");
  static_assert(a == b, "a == b");
}

// g++ -std=c++17 -o app static.cpp 
void do2()
{
  int a = 10;
  int &b = a;
  static_assert(is_same_v<int,decltype(a)>, "decltype(a) is not int");
  static_assert(is_same_v<int&,decltype(b)>, "decltype(b) is not int&");
  static_assert(is_same_v<int&,decltype(a)>, "decltype(a) is not int&");
}

int main(int argc, char const *argv[])
{
  do1();
  return 0;
}
