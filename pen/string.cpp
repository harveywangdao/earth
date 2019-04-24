#include <cstring>
#include <string>
#include <cstdio>
#include <iostream>

using namespace std;

void fun1()
{
  string s1 = "hello, world!";
  string s2 = s1; 

  cout << "before: " << s2 << endl;
  char* ptr = const_cast<char*>(s1.c_str());
  *ptr = 'f';
  cout << "after:ptr " << ptr << endl;
  cout << "after:s1 " << s1 << endl;
  cout << "after:s2 " << s2 << endl;
}

void fun2()
{
  string s1 = "hello, world!";
  string s2 = s1; 

  cout << "before: " << s2 << endl;
  s1[0] = 'f';
  cout << "after:s1 " << s1 << endl;
  cout << "after:s2 " << s2 << endl;
}

int main(int argc, char const *argv[])
{
  cout << "fun1: " << endl;
  fun1();

  cout << "fun2: " << endl;
  fun2();

  return 0;
}
