#include <iostream>
#include <thread>
#include <chrono>
#include <mutex>
#include <atomic>
#include <ratio>
#include <algorithm>
#include <iomanip>
#include <vector>
#include <numeric>
#include <ctime>
#include <shared_mutex>

// g++ -std=c++11 -o app mutex.cpp -lpthread
int num = 0;
std::mutex mu;

void fn1()
{
  std::cout << "child thread start" << std::endl;
  for (int i = 0; i < 10000000; ++i)
  {
    std::lock_guard<std::mutex> guard(mu);
    num++;
  }
  std::cout << "child thread end" << std::endl;
}

void do1()
{
  clock_t start = clock();

  std::thread t1(fn1);
  std::thread t2(fn1);
  t1.join();
  t2.join();

  clock_t end = clock();
  std::cout << "num = " << num << " " << end - start << "ms" <<std::endl;
}

void fn2()
{
  std::cout << "child thread start" << std::endl;
  for (int i = 0; i < 10000000; ++i)
  {
    std::unique_lock<std::mutex> unique(mu);
    num++;
  }
  std::cout << "child thread end" << std::endl;
}

void do2()
{
  clock_t start = clock();

  std::thread t1(fn2);
  std::thread t2(fn2);
  t1.join();
  t2.join();

  clock_t end = clock();
  std::cout << "num = " << num << " " << end - start << "ms" << std::endl;
}

using namespace std::chrono_literals;

void do3()
{
  std::timed_mutex tm;
  tm.lock();

  auto start = std::chrono::steady_clock::now();
  if (tm.try_lock_for(2s))
  {
    std::cout << "times lock success" << std::endl;
  }
  else
  {
    std::cout << "times lock fail" << std::endl;
  }
  auto end = std::chrono::steady_clock::now();
  auto diff = std::chrono::duration_cast<std::chrono::milliseconds>(end - start);
  std::cout << diff.count() << std::endl;
}

void do4()
{
  std::recursive_mutex rm;
  //std::mutex rm;
  rm.lock();
  std::cout << "recursive_mutex 1" << std::endl;
  rm.lock();
  std::cout << "recursive_mutex 2" << std::endl;
}

std::recursive_timed_mutex rtm;

void fn3()
{
  auto start = std::chrono::steady_clock::now();
  if (rtm.try_lock_for(2s))
    std::cout << "recursive_timed_mutex success 2" << std::endl;
  else
    std::cout << "recursive_timed_mutex fail 2" << std::endl;

  auto end = std::chrono::steady_clock::now();
  auto diff = std::chrono::duration_cast<std::chrono::milliseconds>(end - start);
  std::cout << diff.count() << std::endl;
}

void do5()
{
  if (rtm.try_lock_for(2s))
    std::cout << "recursive_timed_mutex success 1" << std::endl;
  else
    std::cout << "recursive_timed_mutex fail 1" << std::endl;
  std::thread t1(fn3);
  t1.join();
}

void do6()
{
  std::shared_mutex sm;

  sm.lock_shared();
  std::cout << "shared_mutex 1" << std::endl;
  sm.lock_shared();
  std::cout << "shared_mutex 2" << std::endl;
  sm.unlock_shared();
  sm.unlock_shared();

  sm.lock();
  std::cout << "shared_mutex 3" << std::endl;
  sm.lock_shared();
  std::cout << "shared_mutex 4" << std::endl;
}

void do7()
{
  std::shared_timed_mutex stm;
  stm.lock();
  //stm.try_lock_for(1s);
  //stm.unlock();
  std::cout << "shared_timed_mutex 1" << std::endl;

  if (stm.try_lock_shared_for(4s))   // 超时没起作用
  //if (stm.try_lock_for(3s))
  //if (stm.try_lock())
  //if (stm.try_lock_shared())
    std::cout << "shared_timed_mutex 2" << std::endl;
  else
    std::cout << "shared_timed_mutex 3" << std::endl;
  //stm.unlock_shared();
}

int main(int argc, char const *argv[])
{
  do7();
  return 0;
}