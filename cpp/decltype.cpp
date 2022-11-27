#include <iostream>
#include <typeinfo>
#include <typeindex>

using namespace std;

void do1()
{
  int arr1[8] = {1};
  auto arr2 = arr1;
  auto arr3 = {10,20,30};

  cout << "arr1: " << arr1 << endl;
  cout << "typeid arr1: " << typeid(arr1).name() << endl;
  cout << "sizeof arr1: " << sizeof(arr1) << endl;

  cout << "arr2: " << arr2 << endl;
  cout << "typeid arr2: " << typeid(arr2).name() << endl;
  cout << "sizeof arr2: " << sizeof(arr2) << endl;

  //cout << "arr3: " << arr3 << endl;
  cout << "typeid arr3: " << typeid(arr3).name() << endl;
  //cout << "sizeof arr3: " << sizeof(arr3) << endl;
}

void do2()
{
  int a = 2;
  unsigned int b = 3;
  long c = 34;

  cout << "typeid a: " << typeid(a).name() << endl;
  cout << "typeid b: " << typeid(b).name() << endl;
  cout << "typeid c: " << typeid(c).name() << endl;

  cout << "a is int: " << (typeid(a).name() == typeid(int).name()) << endl;
  cout << "b is unsigned int: " << (typeid(b).name() == typeid(unsigned int).name()) << endl;
  cout << "c id long: " << (typeid(c).name() == typeid(long).name()) << endl;

  const type_info &t1 = typeid(int);
  const type_info &t2 = typeid(a);
  const type_info *t3 = &typeid(int);

  cout << t1.name() << endl;
  cout << t3->name() << endl;

  cout << "t1.hash_code: " << t1.hash_code() << endl;
  cout << "t2.hash_code: " << t2.hash_code() << endl;

  cout << "t1 == t2: " << (t1 == t2) << endl;
  cout << "type_index(t1) == type_index(t2): " << (type_index(t1) == type_index(t2)) << endl;
}

void do3()
{
  cout << "type_index(typeid(int)): " << type_index(typeid(int)) << endl;
}

int main(int argc, char const *argv[])
{
  do3();
  return 0;
}
