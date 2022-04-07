#include <iostream>
#include <thread>
#include <chrono>
#include <mutex>
#include <atomic>

using namespace std;

// g++ -std=c++11 -o app thread.cpp -lpthread
mutex mu;
int num = 0;

void add1()
{
  lock_guard<mutex> guard(mu);
  num++;
}

void fn1()
{
  cout << "child thread start" << endl;

  for (int i = 0; i < 10000000; ++i)
  {
    add1();
  }

  cout << "child thread end" << endl;
}

void func1()
{
  clock_t start = clock();

  thread t1(fn1);
  thread t2(fn1);

  t1.join();
  t2.join();

  clock_t end = clock();

  cout << "num = " << num << " " << end - start << "ms" <<endl;
}

void add2()
{
  unique_lock<mutex> unique(mu);
  num++;
}

void fn2()
{
  cout << "child thread start" << endl;

  for (int i = 0; i < 10000000; ++i)
  {
    add2();
  }

  cout << "child thread end" << endl;
}

void func2()
{
  clock_t start = clock();

  thread t1(fn2);
  thread t2(fn2);

  t1.join();
  t2.join();

  clock_t end = clock();

  cout << "num = " << num << " " << end - start << "ms" <<endl;
}

int main(int argc, char const *argv[])
{
  //func1();
  func2();

  return 0;
}