#include <iostream>

constexpr void m1()
{
  int o = 2;
  o++;
}

constexpr void m2(int &n)
{
  n++;
}

void do1()
{
  m1();

  int n1 = 20;
  m2(n1);
  std::cout << n1 <<std::endl;
}

struct Road
{
  int n;
  long m;
};

void do2()
{
  constexpr Road r1 = Road{10,20};
  //r1.n = 30;
  std::cout << r1.n << " " << r1.m <<std::endl;
}

int main(int argc, char const *argv[])
{
  do2();
  return 0;
}
