#include <cstring>
#include <string>
#include <cstdio>
#include <iostream>

using namespace std;

void selfplusplus(int &v)
{
  v++;
}

void swap(int& x, int& y)
{
  int temp = x;
  x = y;
  y = temp;
}

int& add(int a, int b)
{
  int w = a + b;
  int& d = w;
  return d;
}

int c;

int& add2(int a, int b)
{
  c = a + b;
  return c;
}

void do1()
{
  int value = 12;
  int &refValue = value;

  int v2 = 88;
  refValue = v2;

  refValue++;
  selfplusplus(refValue);
  selfplusplus(value);

  cout << value << endl << refValue << endl;
  cout << &value << endl << &refValue << endl;

  int a = 100;
  int b = 200;

  swap(a,b);

  cout << a << endl << b << endl;

  cout << add(1,2) << endl;
  cout << add2(1,2) << endl;
}

int &add3(int &a, int &b)
{
  static int c;
  c = a+b;
  return c;
}

int &add4(int &a, int &b, int &c)
{
  c = a+b;
  return c;
}

void do2()
{
  int a=12, b=13;
  int c;
  cout << "c: " << add4(a,b,c) << endl;
}

// 引用在语法层面只能在申明的时候初始化,无法重新赋值,因为后面对引用的所有操作都相当于操作指向的那个变量
void do3()
{
  int a = 10;
  int &refa = a;

  int b = 20;
  int &refb = b;

  refa = b;
  refb = 30;
  refa = refb;

  cout << "refa: " << refa << ", &refa: " << &refa << ", &a: " << &a << endl;
}

void f1(int &a, int const & b)
{
  a = 30;
  //b = 40; // const引用不能改变值
}

void do4()
{
  int a=10,b=20;
  f1(a,b);
  cout << "a: " << a << ", b: " << b << endl;
}

void do5()
{
  const int a = 20;
  //int &b = a; // binding reference of type int& to const int discards qualifiers
  const int &b = a;
  //b++;
  cout << "a = " << a << endl;
  cout << "b = " << b << endl;
}

int main(int argc, char const *argv[])
{
  do5();
  return 0;
}