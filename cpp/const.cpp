#include <iostream>

using namespace std;

// 保护变量本身
void do1()
{
  const int n1 = 2;
  //n1 = 3;
  cout << "n1: " << n1 << endl;

  int n2 = 11;

  int * const p1 = &n2;

  int n3 = 12;

  //p1 = &n3;

  cout << "p1: " << p1 << endl;
  cout << "*p1: " << *p1 << endl;
}

// 保护指针指向的内存
void do2()
{
  int n1 = 12;

  const int *p1 = &n1;

  //*p1 = 13;

  int n2 = 14;

  p1 = &n2;

  cout << "p1: " << p1 << endl;
  cout << "*p1: " << *p1 << endl;
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

int getn(int n)
{
  return n+1; 
}

// 值不会改变并且在编译过程就能得到计算结果的表达式
// 使用gcc时constexpr指针所指变量必须是全局变量或者static变量
void do4()
{
  constexpr int n1 = 12;
  //n1 = 13;

  const int n2 = getn(10);
  //constexpr int n3 = getn(10);

  int n4 = 15;
  const int *p1 = &n4;
  //constexpr int *p2 = &n4;

  const int n5 = 16;
  const int *p3 = &n5;
  //constexpr int *p4 = &n5;

  static int n6 = 20;
  static int n7 = 30;

  const int *p5 = &n6;
  constexpr int *p6 = &n6;

  //*p5 = 21; // const保护指针指向的内存
  *p6 = 22;   // constexpr不保护指针指向的内存
  
  p5 = &n7;
  //p6 = &n7; // constexpr保护指针变量

  cout << "n6: " << n6 << endl;

  static int n8 = 40;
  //int * constexpr p7 = &n8; // constexpr不能放这个位置

  int n9 = 50;
  //int * constexpr p8 = &n9; // constexpr不能放这个位置
}

constexpr unsigned factorial(unsigned n)
{
  return n < 2 ? 1 : n * factorial(n-1);
}

consteval unsigned combination(unsigned m, unsigned n) {
  return factorial(n) / factorial(m) / factorial(n - m);
}

void do5()
{
  // consteval
  static_assert(factorial(6) == 720);
  static_assert(combination(4,8) == 70);

  constexpr unsigned x{factorial(4)};
  std::cout << x << '\n';
}

int main(int argc, char const *argv[])
{
  do4();
  return 0;
}
