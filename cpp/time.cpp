#include <iostream>
#include <ratio>
#include <chrono>
#include <algorithm>
#include <iomanip>
#include <ctime>
#include <thread>
#include <vector>
#include <numeric>
#include <ctime>

void do1()
{
  using two_third = std::ratio<2, 3>;
  using one_sixth = std::ratio<1, 6>;
  
  using sum = std::ratio_add<two_third, one_sixth>;
  std::cout << "2/3 + 1/6 = " << sum::num << '/' << sum::den << std::endl;;
}

void do2()
{
  std::nano v;
  std::chrono::nanoseconds n1(1);
  std::chrono::microseconds n2(2);
  std::chrono::milliseconds n3(3);
  std::chrono::seconds n4(4);
  std::chrono::minutes n5(5);
  std::chrono::hours n6(6);

  std::chrono::duration<int, std::kilo> ks(3); // 3000 seconds
  std::chrono::duration<double, std::ratio<1, 30>> hz30(3.5);

  n1 = n4;
  std::cout << n1.count() << std::endl;
}

using namespace std::literals;

void do3()
{
  auto start = std::chrono::steady_clock::now();
  std::this_thread::sleep_for(2700ms);
  auto end = std::chrono::steady_clock::now();
  auto diff = std::chrono::duration_cast<std::chrono::microseconds>(end - start);
  std::cout << diff.count() << "µs ≈ " << (end - start) / 1ms << "ms ≈ " << (end - start) / 1s << "s" << std::endl;
}

using namespace std::chrono_literals;

void do4()
{
  auto t1 = std::chrono::system_clock::now();  
  std::time_t tt1 = std::chrono::system_clock::to_time_t(t1);

  std::this_thread::sleep_for(2700ms);

  auto t2 = std::chrono::system_clock::now();  
  std::time_t tt2 = std::chrono::system_clock::to_time_t(t2);

  std::cout << "tt2-tt1 == " << tt2-tt1 << "s" << std::endl;
  
  std::time_t t3 = std::time(nullptr);
  auto from = std::chrono::system_clock::from_time_t(t3);
  
  std::this_thread::sleep_for(500ms);
  
  auto diff = std::chrono::system_clock::now() - from;
  
  //std::cout << diff.count() << " (" << std::chrono::round<std::chrono::milliseconds>(diff).count() << ")" << std::endl;
}

void do5()
{
  auto start = std::chrono::steady_clock::now();
  std::vector<int> v(20000000, 42);
  std::accumulate(v.begin(), v.end(), 0u);
  auto end = std::chrono::steady_clock::now();
  std::chrono::duration<double> diff = end - start;
  std::cout << "Time to fill and iterate a vector of " << std::setw(9) << " ints : " << diff.count() << " s\n";
}

void do6()
{
  auto start = std::chrono::high_resolution_clock::now();
  std::vector<int> v(20000000, 42);
  std::accumulate(v.begin(), v.end(), 0u);
  auto end = std::chrono::high_resolution_clock::now();
  std::chrono::duration<double> diff = end - start;
  std::cout << "Time to fill and iterate a vector of " << std::setw(9) << " ints : " << diff.count() << " s\n";
}

void do7()
{
  const auto p0 = std::chrono::time_point<std::chrono::system_clock>{};
  const auto p1 = std::chrono::system_clock::now();
  const auto p2 = p1 - std::chrono::hours(24);
  
  std::time_t epoch_time = std::chrono::system_clock::to_time_t(p0);
  std::cout << "epoch: " << std::ctime(&epoch_time);
  std::time_t today_time = std::chrono::system_clock::to_time_t(p1);
  std::cout << "today: " << std::ctime(&today_time);
  
  std::cout << "hours since epoch: " << std::chrono::duration_cast<std::chrono::hours>(p1.time_since_epoch()).count() << std::endl;
  std::cout << "yesterday, hours since epoch: " << std::chrono::duration_cast<std::chrono::hours>(p2.time_since_epoch()).count() << std::endl;
}

int main(int argc, char const *argv[])
{
  do7();
  return 0;
}
