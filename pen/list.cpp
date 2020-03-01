#include <iostream>
#include <list>

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

void func1()
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

int main(int argc, char const *argv[])
{
  func1();

  return 0;
}