#include <iostream>
#include <list>
#include <forward_list>

using namespace std;

void printList(list<int> lst)
{
  list<int>::iterator iter;
  for (iter = lst.begin(); iter != lst.end(); ++iter)
  {
    cout << *iter << " ";
  }
  cout << endl;
}

void do1()
{
  list<int> l1;
  list<int> l2(10);
  list<int> l3(10, 5);
  list<int> l4(l3);
  
  printList(l1);
  printList(l2);
  printList(l3);
  printList(l4);

  for (int i = 0; i < 5; ++i)
  {
    l1.push_front(i);    
  }
  printList(l1);

  for (int i = 5; i < 10; ++i)
  {
    l1.push_back(i);    
  }
  printList(l1);

  cout << "l1.size() " << l1.size() << endl;
  cout << "l1.max_size() " << l1.max_size() << endl;

  list<int>::iterator iter = l1.begin();
  l1.insert(iter, 99);
  l1.insert(iter, 5);
  cout << "after l1 insert" << endl;
  printList(l1);

  l1.remove(5);
  cout << "after l1 remove" << endl;
  printList(l1);
}

void do2()
{
  std::list<int> l = {1, 2, 3, 4};
  l.push_front(0);
  l.push_back(5);
  for (int n : l)
    std::cout << n << " ";
  std::cout << std::endl;
  //l[2];
}

void do3()
{
  std::forward_list<int> f1{1,2,3,4};

  f1.push_front(-1);
  f1.insert_after(f1.begin(), 0);
  //f1.insert_after(f1.end(), 5);

  for (auto e : f1)
  {
    std::cout << e << " ";
  }
  std::cout << std::endl;
}

int main(int argc, char const *argv[])
{
  do3();
  return 0;
}