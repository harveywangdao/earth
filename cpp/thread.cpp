#include <iostream>
#include <thread>
#include <chrono>
#include <mutex>
#include <atomic>

using namespace std;

// g++ -std=c++11 -o app thread.cpp -lpthread
mutex mu;
int num = 0;

atomic_int num2{0};

void fn1()
{
  cout << "child thread start" << endl;

  for (int i = 0; i < 10000000; ++i)
  {
    mu.lock();
    num++;
    mu.unlock();
  }

  for (int i = 0; i < 10000000; ++i)
  {
    num2++;
  }

  //this_thread::sleep_for(chrono::seconds(3));
  //this_thread::yield();
  //cout<< this_thread::get_id() << endl;

  cout << "child thread end" << endl;
}

void func1()
{
  cout << "thread::hardware_concurrency() = " << thread::hardware_concurrency() << endl;

  clock_t start = clock();

  thread t1(fn1);
  thread t2(fn1);

  //cout << "t1 id is " << t1.get_id() << endl;
  //cout << "t2 id is " << t2.get_id() << endl;

  t1.join();
  t2.join();

  clock_t end = clock();

  cout << "num = " << num << " " << end - start << "ms" <<endl;
  cout << "num2 = " << num2 << " " <<end - start << "ms" <<endl;
}

void fn2(int n)
{
  cout << "child thread start " << n << endl;

  cout << "child thread end" << endl;
}

void func2()
{  
  clock_t start = clock();

  thread t1(fn2, 1);
  thread t2(fn2, 2);

  t1.join();
  t2.join();

  clock_t end = clock();

  cout << end - start << "ms" <<endl;
}

int main(int argc, char const *argv[])
{
  //func1();
  func2();

  return 0;
}