#include <iostream>
#include <vector>
#include <array>
#include <deque>
#include <utility>
#include <algorithm>

using namespace std;

void printVector(vector<int> &v)
{
  for (int i = 0; i < v.size(); ++i)
  {
    cout << v[i];
    if (i != v.size()-1)
    {
      cout << ",";
    }
  }

  if (v.size() == 0)
  {
    cout << "empty vector";
  }
  cout << endl;
}

void do1()
{
  vector<int> v1;
  vector<int> v2(10);
  vector<int> v3(10, 5);
  vector<int> v4(v3);

  printVector(v1);
  printVector(v2);
  printVector(v3);
  printVector(v4);

  for (int i = 0; i < 10; ++i)
  {
    v1.push_back(i);
  }

  printVector(v1);

  for (vector<int>::iterator iter = v1.begin(); iter != v1.end(); ++iter)
  {
    cout << *iter;
    auto temp = iter;
    if (++temp != v1.end())
    {
      cout << ",";
    }
  }
  cout << endl;

  v1.pop_back();
  cout << "after v1 pop_back" << endl;
  printVector(v1);

  v1.clear();
  cout << "after v1 clear" << endl;
  printVector(v1);

  cout << "v1.empty() " << v1.empty() << endl;
  cout << "v1.size() " << v1.size() << endl;
  cout << "v1.capacity() " << v1.capacity() << endl;
  cout << "v1.max_size() " << v1.max_size() << endl;
}

class People
{
  int age;
public:
  People():age(0)
  {
    std::cout << "default constructor" << ", age: " << age << std::endl;
  }
  People(int a):age(a)
  {
    std::cout << "constructor with age" << ", age: " << age << std::endl;
  }
  People(const People& p)
  {
    age = p.age;
    std::cout << "copy constructor" << ", age: " << age << std::endl;
  }
  People(People&& p)
  {
    age = p.age;
    std::cout << "move constructor" << ", age: " << age << std::endl;
  }
  ~People()
  {
    std::cout << "destroy" << ", age: " << age << std::endl;
  }
};

void do2()
{
  std::vector<People> v1;
  v1.reserve(16);
  People p1(10);
  v1.push_back(p1);   // constructor with age AND copy constructor
  std::cout << std::endl;
  v1.push_back(People(20)); // constructor with age AND move constructor
  std::cout << std::endl;
  v1.emplace_back(30);   //constructor with age,少一次构造
}

void do3()
{
  std::deque<int> d1 = {7, 5, 16, 8};
  d1.push_front(13);
  d1.push_back(25);
  for (int n : d1)
  {
    std::cout << n << " ";
  }
  std::cout << std::endl;

  std::cout << d1[3] << std::endl;
}

void do4()
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
  do4();
  return 0;
}
