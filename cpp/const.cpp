#include <iostream>

using namespace std;

// 保护变量本身
void do1()
{
  const int n1 = 2;
  //n1 = 3;

  int n2 = 11;
  int * const p1 = &n2;

  int n3 = 12;
  //p1 = &n3;
}

// 保护指针指向的内存
void do2()
{
  int n1 = 12;
  const int *p1 = &n1;
  //*p1 = 13;

  int n2 = 14;
  p1 = &n2;
}

class People {
public:
  int n;

  People(int _n):n(_n)
  {

  }

  int getNum1()
  {
    n = 13;
    return n;
  }

  int getNum2() const
  {
    //n = 14;
    return n;
  }
};

// const修饰类成员函数,其目的是防止成员函数修改被调用对象的值
void do3()
{
  People p1(12);
  auto n1 = p1.getNum1();
  cout << "n1: " << n1 << endl;

  auto n2 = p1.getNum2();
  cout << "n2: " << n2 << endl;
}

// constexpr和const一样不能修改值
void do4()
{
  const int n1 = 10;
  // n1 = 11;

  constexpr int n2 = 20;
  // n2 = 21;
}

int g_n1 = 30;
// 使用gcc时constexpr指针所指变量必须是全局变量或者static变量
void do5()
{
  int n1 = 10;
  // constexpr int *p1 = &n1; // &n1 is not a constant expression

  static int n2 = 20;
  constexpr int *p2 = &n2;

  constexpr int *p3 = &g_n1;
}

void do6()
{
  static int n1 = 10;
  static int n2 = 20;

  const int *p1 = &n1;
  constexpr int *p2 = &n1;

  // *p1 = 11; // const保护指针指向的内存
  *p2 = 12;    // constexpr不保护指针指向的内存
  cout << "n1: " << n1 << endl;
  
  p1 = &n2;
  // p2 = &n2; // constexpr保护指针变量

  static int n3 = 30;
  int constexpr *p3 = &n3;  // constexpr可以放这个位置
  static int n4 = 40;
  // int * constexpr p4 = &n4; // constexpr不能放这个位置
}

int addone1(int n)
{
  return n+1;
}
const int addone2(int n)
{
  return n+1;
}
constexpr int addone3(int n)
{
  return n+1;
}

// constexpr变量 = constexpr函数
void do7()
{
  // constexpr int n1 = addone1(10); // 不能在编译期确定值
  // constexpr int n2 = addone2(10); // 不能在编译期确定值
  constexpr int n3 = addone3(10);
  cout << "n3: " << n3 << endl;

  int x1 = 10;
  // constexpr int n4 = addone3(x1); // x1 is not usable in a constant expression

  const int x2 = 10;
  constexpr int n5 = addone3(x2);
  cout << "n5: " << n5 << endl;

  constexpr int x3 = 10;
  constexpr int n6 = addone3(x3);
  cout << "n6: " << n6 << endl;
}

void do8()
{
  int n1 = addone2(10);
  cout << "n1: " << n1 << endl;
  n1++;
  cout << "n1: " << n1 << endl;

  const int n2 = addone2(10);
  cout << "n2: " << n2 << endl;
  // n2++;
}

void genarr(int n)
{
  int arr1[n]; // c++支持
  arr1[0] = 12;

  const int m = 20;
  int arr2[m];
  arr2[0] = 12;
}

void do9()
{
  genarr(10);

  int a = 10;
  genarr(a);
}

struct A
{
  int a;
  long b;
};

void do10()
{
  constexpr A a1{1,2};
  // a1.a = 2;
  cout << a1.a << endl;
  cout << a1.b << endl;

  const A a2{1,2};
  // a2.a = 2;
}

consteval int addone4(int n)
{
  return n+1;
}

// g++ -std=c++20 -o app const.cpp
void do11()
{
  constexpr int n1 = addone4(10);
  cout << "n1: " << n1 << endl;

  int a = 10;
  // constexpr int n2 = addone4(a); // a is not usable in a constant expression

  // int n2 = addone4(a);  // consteval依然起作用
  int n3 = addone3(a);     // constexpr不起作用,变成运行时计算
}

constinit int g_n2 = 20;
//constinit int g_n4 = addone2(10);
constinit int g_n5 = addone3(10);
constinit int g_n6 = addone4(10);
//constinit int g_n7 = addone3(g_n1);
const int g_n8 = 10;
constinit int g_n9 = addone3(g_n8);

//constinit int g_n3 = addone1(10);
const int g_n10 = addone1(10); // 可能不在编译期计算值

void do12()
{
  // constinit int n1 = 10; // can only be applied to a variable with static or thread storage duration
  constinit static int n2 = 10;

  cout << "n2: " << n2 << endl;
  cout << "g_n2: " << g_n2 << endl;

  n2++;
  cout << "n2: " << n2 << endl;

  g_n2++;
  cout << "g_n2: " << g_n2 << endl;
}

int main(int argc, char const *argv[])
{
  do12();
  return 0;
}
