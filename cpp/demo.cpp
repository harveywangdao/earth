#include <iostream>

using namespace std;

/*
const 修饰指针 修饰指针指向的数据 修饰函数返回值
int f() const
constexpr
& 引用
struct
class XX {
  public:
  private:
}
abstract
virtual
explicit
构造函数
析构函数
纯虚函数
static 全局变量
static 局部变量
static 函数
static class成员
多态
nullptr 空指针
reinterpret_cast
noexcept
f() = delete
static_cast
thread_local
std::array
std::atomic_bool
[[noreturn]]
[[nodiscard]]
[[maybe_unused]]
*/

void do1()
{
  bool ok;
  cout << "ok: " << ok << endl;
  ok = true;
  cout << "ok: " << ok << endl;
  ok = false;
  cout << "ok: " << ok << endl;

  if (1 == true)
  {
    cout << "1 == true" << endl;
  }
  else 
  {
    cout << "1 != true" << endl;
  }

  if (2 == true)
  {
    cout << "2 == true" << endl;
  }
  else 
  {
    cout << "2 != true" << endl;
  }

  if (0 == false)
  {
    cout << "0 == false" << endl;
  }
  else 
  {
    cout << "0 != false" << endl;
  }
}

typedef unsigned char uchar;
using byte = unsigned char; //g++ -std=c++11 -o demo demo.cpp 

void do2()
{
  int num1;
  cout << "int: " << sizeof(num1) << endl;

  long num2;
  cout << "long: " << sizeof(num2) << endl;

  long long num3;
  cout << "long long : " << sizeof(num3) << endl;

  float num4;
  cout << "float: " << sizeof(num4) << endl;

  double num5;
  cout << "double: " << sizeof(num5) << endl;

  short num6;
  cout << "short: " << sizeof(num6) << endl;

  char num7;
  cout << "char: " << sizeof(num7) << endl;

  byte num8;
  cout << "byte: " << sizeof(num8) << endl;

  uchar num9;
  cout << "uchar: " << sizeof(num9) << endl;
}

int main(int argc, char const *argv[])
{
  do2();
  return 0;
}