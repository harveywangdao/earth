#include <iostream>
#include <map>
#include <set>
#include <unordered_map>
#include <unordered_set>
#include <stack>
#include <queue>
#include <tuple>

using namespace std;

void printMap(map<int, string> m)
{
  map<int, string>::iterator iter;
  for (iter = m.begin(); iter != m.end(); ++iter)
  {
    cout << iter->first << " " << iter->second << " ";
  }
  cout << endl;
}

void do1()
{
  map<int, string> m1;
  pair<map<int, string>::iterator, bool> ret;

  m1.insert(pair<int, string>(1, "a"));
  ret = m1.insert(pair<int, string>(1, "aa"));
  if (ret.second == false)
  {
    cout << "already existed " << ret.first->first << " " << ret.first->second << endl;
  }

  m1.insert(map<int, string>::value_type(2, "b"));
  m1.insert(map<int, string>::value_type(2, "bb"));
  m1[3] = "c";
  m1[4] = "d";
  m1[5] = "e";
  m1[5] = "ee";
  printMap(m1);

  map<int, string>::iterator iter;
  iter = m1.find(3);
  if (iter != m1.end())
  {
    cout << iter->first << " " << iter->second << endl;
    m1.erase(iter);
  }
  printMap(m1);

  cout << m1[5] << endl;

  cout << "m1.erase(2) " << m1.erase(2) << endl;
  cout << "m1.erase(21) " << m1.erase(21) << endl;
  printMap(m1);

  cout << "m1.size() " << m1.size() << endl;
}

void do2()
{
  std::set<int> s1;
  s1.insert(1);
  s1.insert(3);
  s1.insert(5);
  s1.insert(7);

  for (auto n : s1)
  {
    std::cout << n << " ";
  }
  std::cout << std::endl;

  std::cout << "empty: " << s1.empty() << std::endl;
  std::cout << "size: " << s1.size() << std::endl;
  std::cout << "max_size: " << s1.max_size() << std::endl;
  
  std::set<int>::iterator p = s1.find(1);
  if (p != s1.end())
  {
    std::cout << *p << std::endl;
  }
}

void do3()
{
  std::multiset<int> a;
  a.insert(4);
  a.insert(3);
  a.insert(4);
  a.insert(2);
  a.insert(1);
  a.insert(4);
  for (auto n : a)
  {
    std::cout << n << " ";
  }
  std::cout << std::endl;

  std::set<int>::iterator p = a.find(1);
  if (p != a.end())
  {
    std::cout << *p << std::endl;
  }

  std::cout << a.count(4) << std::endl;
  a.erase(a.find(4));
  std::cout << a.count(4) << std::endl;
  a.erase(4);
  std::cout << a.count(4) << std::endl;
}

void do4()
{
  std::map<std::string, std::string> m;
  m.emplace(std::make_pair(std::string("111"), std::string("aaa")));
  m.emplace(std::make_pair("222", "bbb"));
  m.emplace("333", "ccc");

  for (auto e: m)
  {
    std::cout << e.first << " : " << e.second << std::endl;
  }
}

void do5()
{
  std::multimap<std::string, std::string> m;
  m.emplace(std::make_pair(std::string("111"), std::string("aaa")));
  m.emplace(std::make_pair("222", "bbb"));
  m.emplace("333", "ccc");
  m.emplace("222", "ddd");

  for (auto e: m)
  {
    std::cout << e.first << " : " << e.second << std::endl;
  }
}

void do6()
{
  std::unordered_set<int> us1;
  us1.insert(1);
  us1.insert(2);
  us1.insert(3);
  us1.insert(4);

  for (auto e: us1)
  {
    std::cout << e << " ";
  }
  std::cout << std::endl;
}

void do7()
{
  std::unordered_multiset<int> us1;
  us1.insert(4);
  us1.insert(1);
  us1.insert(4);
  us1.insert(2);
  us1.insert(3);
  us1.insert(4);

  for (auto e: us1)
  {
    std::cout << e << " ";
  }
  std::cout << std::endl;
}

void do8()
{
  std::unordered_map<int,int> um1;
  um1[1] = 10;
  um1[2] = 20;
  um1[3] = 30;

  for (auto e: um1)
  {
    std::cout << e.first << " : " << e.second << std::endl;
  }
  std::cout << std::endl;
}

void do9()
{
  std::unordered_multimap<int,int> um1;
  um1.emplace(3, 30);
  um1.emplace(1, 10);
  um1.emplace(2, 20);
  um1.emplace(3, 30);

  for (auto e: um1)
  {
    std::cout << e.first << " : " << e.second << std::endl;
  }
  std::cout << std::endl;
}

void do10()
{
  std::stack<int> s1;
  s1.push(1);
  s1.push(2);
  s1.push(3);
  s1.push(4);

  std::cout << s1.top() << std::endl;
  s1.pop();
  std::cout << s1.top() << std::endl;

  /*for (auto e: s1)
  {
    std::cout << e << " ";
  }
  std::cout << std::endl;*/
}

void do11()
{
  std::queue<int> q1;
  q1.push(1);
  q1.push(2);
  q1.push(3);
  q1.push(4);
  q1.push(5);

  std::cout << q1.front() << std::endl;
  std::cout << q1.back() << std::endl;

  q1.pop();
  std::cout << q1.front() << std::endl;

  /*for (auto e: q1)
  {
    std::cout << e << " ";
  }
  std::cout << std::endl;*/
}

void do12()
{
  std::priority_queue<int> q1;
  q1.push(1);
  q1.push(2);
  q1.push(3);
  q1.push(4);
  q1.push(7);
  q1.push(5);

  /*for (auto e: q1)
  {
    std::cout << e << " ";
  }
  std::cout << std::endl;*/

  std::cout << q1.top() << std::endl;
  q1.pop();
  std::cout << q1.top() << std::endl;
}

void do13()
{
  std::tuple<int,std::string,double> t1(1,"aaa",1.2);
  std::tuple<int,std::string,double> t2 = std::make_tuple(2,"bbb",2.2);
  //std::cout << t1 << std::endl;
}

int main(int argc, char const *argv[])
{
  do13();
  return 0;
}