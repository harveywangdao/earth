#include <iostream>

using namespace std;

int div(int a, int b)
{
  if (b == 0)
  {
    throw "div zero";
  }

  return a/b;
}

int main(int argc, char const *argv[])
{
  int a=8, b=0, c=999;

  try {
    c = div(a, 3);
    cout << "c=" << c << endl;

    c = div(a, b);
    cout << "c=" << c << endl;
  } catch (const char* msg) {
    cerr << msg << endl;
  }

  cout << "app end" << endl;

  return 0;
}