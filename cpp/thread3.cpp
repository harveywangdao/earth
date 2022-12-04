#include <iostream>
#include <string>
#include <mutex>
#include <thread>

// g++ -std=c++11 -o app thread3.cpp -lpthread
std::mutex mu;

thread_local int n1 = 10;
thread_local static int n2 = 20;

void m1(std::string const& s)
{
  thread_local int n3 = 30;
  thread_local static int n4 = 40;

  n1++;
  n2++;
  n3++;
  n4++;
  std::lock_guard<std::mutex> guard(mu);
  std::cout << s << " " << n1 << " " << n2 << " " << n3 << " " << n4 << std::endl;
}

void f1(std::string const& s){
  for (int i = 0; i < 4; ++i)
  {
    m1(s);
  }
}

int main(){
  std::thread t1(f1, "t1");
  std::thread t2(f1, "t2");
  std::thread t3(f1, "t3");
  std::thread t4(f1, "t4");

  t1.join();
  t2.join();
  t3.join();
  t4.join();
}

