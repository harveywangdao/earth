#include <iostream>
#include <ratio>
#include <chrono>
#include <algorithm>
#include <iomanip>
#include <ctime>
#include <thread>

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
  std::chrono::time_point<std::chrono::system_clock> now = std::chrono::system_clock::now();
  std::time_t t_c = std::chrono::system_clock::to_time_t(now - 24h);
  std::time_t oldt = std::time(nullptr);
  std::this_thread::sleep_for(2700ms);

  std::time_t newt = std::chrono::system_clock::to_time_t(std::chrono::system_clock::now());
  std::cout << "oldt-newt == " << oldt-newt << " s\n";

  std::chrono::time_point<std::chrono::steady_clock> start = std::chrono::steady_clock::now();
/*  std::cout << "Different clocks are not comparable: \n"
                << " System time: " << now.time_since_epoch() << "\n"
                << " Steady time: " << start.time_since_epoch() << "\n";*/

  const auto end = std::chrono::steady_clock::now();
      std::cout
        << "Slow calculations took "
        << std::chrono::duration_cast<std::chrono::microseconds>(end - start).count() << "µs ≈ "
        << (end - start) / 1ms << "ms ≈ " // almost equivalent form of the above, but
        << (end - start) / 1s << "s.\n";  // using milliseconds and seconds accordingly
}

int main(int argc, char const *argv[])
{
  do3();
  return 0;
}