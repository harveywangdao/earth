#include <iostream>

using namespace std;

class A {
private:
  int n;
};

class B {
private:
  int n;
public:
  B(int _n): n(_n) {}
};

class C {
private:
  int n;
public:
  C() = default;
  C(const C& c);
  C& operator=(const C& c);
  C(int _n): n(_n) {}
};

C::C(const C& c) = default;
C& C::operator=(const C& c) = default;

void do1()
{
  A a;

  //B b1; // B::B()
  B b2(2);
  B b3(b2);

  C c1; // C::C()
  C c2(2);
  C c3(c2);
}

class D
{
public:
  D(int a):x(a) {std:cout<<x<<endl;}
  D(double) = delete;
private:
  int x;  
};

void do2()
{
  D d1(1);
  //D d2(1.2);
}

int main(int argc, char const *argv[])
{
  do2();
  return 0;
}