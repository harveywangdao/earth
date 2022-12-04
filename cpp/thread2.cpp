#include <iostream>
#include <string>
#include <mutex>
#include <thread>

std::mutex coutMutex;

thread_local std::string s("hello from ");

void dothing()
{
  thread_local int n = 2;
  n++;
  std::cout << "n: " << n << std::endl;
}

void addThreadLocal(std::string const& s2){
  s += s2;
  // 加锁是为了保护 std::cout 输出正常
  std::lock_guard<std::mutex> guard(coutMutex);
  std::cout << s << std::endl;
  std::cout << "&s: " << &s << std::endl;

  for (int i = 0; i < 3; ++i)
  {
    dothing();
  }

  std::cout << std::endl;
}

int main(){
  std::thread t1(addThreadLocal, "t1");
  std::thread t2(addThreadLocal, "t2");
  std::thread t3(addThreadLocal, "t3");
  std::thread t4(addThreadLocal, "t4");

  t1.join();
  t2.join();
  t3.join();
  t4.join();
}
