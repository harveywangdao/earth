#include <iostream>

using namespace std;

struct test1
{
  char a;
  int b;
  long long c;
};

struct alignas(0) test2
{
  char a;
  int b;
  long long c;
};

struct alignas(2) test3
{
  char a;
  int b;
  long long c;
};

struct alignas(4) test4
{
  char a;
  int b;
  long long c;
};

struct alignas(8) test5
{
  char a;
  int b;
  long long c;
};

struct alignas(16) test6
{
  char a;
  int b;
  long long c;
};

struct alignas(32) test7
{
  char a;
  int b;
  long long c;
};

struct alignas(int) test8
{
  char a;
  int b;
  long long c;
};

alignas(64) char cacheline[64];

void do1()
{
  cout << "alignof(test1): " << alignof(test1) << ", ";
  cout << "sizeof(test1): " << sizeof(test1) << endl;

  cout << "alignof(test2): " << alignof(test2) << ", ";
  cout << "sizeof(test2): " << sizeof(test2) << endl;

  cout << "alignof(test3): " << alignof(test3) << ", ";
  cout << "sizeof(test3): " << sizeof(test3) << endl;

  cout << "alignof(test4): " << alignof(test4) << ", ";
  cout << "sizeof(test4): " << sizeof(test4) << endl;

  cout << "alignof(test5): " << alignof(test5) << ", ";
  cout << "sizeof(test5): " << sizeof(test5) << endl;

  cout << "alignof(test6): " << alignof(test6) << ", ";
  cout << "sizeof(test6): " << sizeof(test6) << endl;

  cout << "alignof(test7): " << alignof(test7) << ", ";
  cout << "sizeof(test7): " << sizeof(test7) << endl;

  cout << "alignof(test8): " << alignof(test8) << ", ";
  cout << "sizeof(test8): " << sizeof(test8) << endl;

  cout << "alignof(cacheline): " << alignof(cacheline) << ", ";
  cout << "sizeof(cacheline): " << sizeof(cacheline) << endl;
}

int main(int argc, char const *argv[])
{
  do1();
  return 0;
}