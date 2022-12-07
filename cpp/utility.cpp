#include <iostream>
#include <utility>
#include <array>
#include <algorithm>

void do1()
{
  int a{10};
  int b(20);

  std::cout << "a = " << a << std::endl;
  std::cout << "b = " << b << std::endl;
  std::swap(a, b);
  std::cout << "a = " << a << std::endl;
  std::cout << "b = " << b << std::endl;

  int c = 30;
  int d{};
  int e();
  std::cout << "c = " << c << std::endl;
  std::cout << "d = " << d << std::endl;
  std::cout << "e = " << e << std::endl;
}

void do2()
{
  //std::swap(10, 20);
  //std::swap(std::move(10), std::move(20));
}

void do3()
{
  int arr1[3] = {1,2,3};
  int arr2[3] = {10,20,30};

  std::cout << "arr1 = " << arr1 << std::endl;
  std::cout << "arr2 = " << arr2 << std::endl;
  std::cout << "&arr1 = " << &arr1 << std::endl;
  std::cout << "&arr2 = " << &arr2 << std::endl;
  std::swap(arr1, arr2);
  std::cout << "arr1 = " << arr1 << std::endl;
  std::cout << "arr2 = " << arr2 << std::endl;
  std::cout << "&arr1 = " << &arr1 << std::endl;
  std::cout << "&arr2 = " << &arr2 << std::endl;

  for (int i = 0; i < 3; ++i)
  {
    std::cout << arr1[i] << std::endl;
  }
  for (int i = 0; i < 3; ++i)
  {
    std::cout << arr2[i] << std::endl;
  }
}

void do4()
{
  int n1 = 10;
  int n2 = 10;
  std::cout << "&n1 = " << &n1 << std::endl;
  std::cout << "&n2 = " << &n2 << std::endl;
  std::swap(n1, n2);
  std::cout << "&n1 = " << &n1 << std::endl;
  std::cout << "&n2 = " << &n2 << std::endl;
}

void do5()
{
  int n = 10;
  std::exchange(n, 20);
  std::cout << "n = " << n << std::endl;
}

void do6()
{
  std::array<int,4> arr1 = {1,2,3,4};
  std::array<int,4> arr2 = arr1;
  std::array<int,4> arr3{1,2,3,4};
  std::array<int,4> arr4{{1,2,3,4}};

  for (int i = 0; i < arr4.size(); ++i)
  {
    std::cout << arr4[i] << std::endl;
  }

  std::cout << std::endl;

  for(const auto& e: arr1)
  {
    std::cout << e << std::endl;
  }

  std::cout << std::endl;

  std::cout << "front: " << arr2.front() << std::endl;
  std::cout << "back: " << arr2.back() << std::endl;
  std::cout << "data: " << arr2.data() << std::endl;
  std::cout << "*data: " << *arr2.data() << std::endl;
  std::cout << "begin: " << arr2.begin() << std::endl;
  std::cout << "end: " << arr2.end() << std::endl;
  //std::cout << "rbegin: " << arr2.rbegin() << std::endl;
  //std::cout << "rend: " << arr2.rend() << std::endl;
  std::cout << "empty: " << arr2.empty() << std::endl;
  std::cout << "max_size: " << arr2.max_size() << std::endl;

  std::cout << std::endl;
  std::array<int,4>::iterator iter;
  for (iter = arr1.begin(); iter != arr1.end(); ++iter)
  {
    std::cout << *iter << std::endl;
  }

  std::cout << std::endl;
  std::array<int,4>::reverse_iterator riter;
  for (riter = arr2.rbegin(); riter != arr2.rend(); ++riter)
  {
    std::cout << *riter << std::endl;
  }

  std::cout << std::endl;
  std::array<int,4> arr5 = {2,5,3,1};
  std::array<int,4> arr6;
  std::sort(arr5.begin(), arr5.end());
  for(const auto& e: arr5)
  {
    std::cout << e << std::endl;
  }
  std::cout << std::endl;
  std::reverse(arr5.begin(), arr5.end());
  for (auto iter = arr5.begin(); iter != arr5.end(); ++iter)
  {
    std::cout << *iter << std::endl;
  }
  std::reverse_copy(arr5.begin(), arr5.end(), arr6.begin()); // 不改变arr5
  std::cout << "\narr5:" << std::endl;
  for(const auto& e: arr5)
  {
    std::cout << e << std::endl;
  }
  std::cout << "\narr6:" << std::endl;
  for(const auto& e: arr6)
  {
    std::cout << e << std::endl;
  }
}

int main(int argc, char const *argv[])
{
  do6();
  return 0;
}
