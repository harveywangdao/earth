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

int main(int argc, char const *argv[])
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

  return 0;
}