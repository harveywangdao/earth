#include <iostream>
#include <type_traits>

using namespace std;

void do1()
{
  //static_assert(false);
  static_assert(true, "this is true");
  //static_assert(false, "this is false");

  const int a = 3;
  const int b = 3;
  static_assert((a == b), "(a == b)");
  static_assert(a == b, "a == b");
}

void do2()
{
  int a = 10;
  static_assert(is_same_v<int,decltype(a)>, "type diff");
}

int main(int argc, char const *argv[])
{
  do2();
  return 0;
}
