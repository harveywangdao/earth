#include <iostream>
#include <typeinfo>
#include <typeindex>
#include <unordered_map>

using namespace std;

// g++ -std=c++11 -o app decltype.cpp
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
  cout << "sizeof arr3: " << sizeof(arr3) << endl;

  for (auto v : arr3)
  {
    cout << v << endl;
  }
  for (int i = 0; i < 3; ++i)
  {
    //cout << arr3[i] << endl;
  }
}

void do2()
{
  int a = 10;
  long b = 20;
  cout << "typeid(a): " << typeid(a).name() << endl;
  cout << "typeid(b): " << typeid(b).name() << endl;

  cout << "typeid(2): " << typeid(2).name() << endl; // 默认int
  cout << "typeid(2+3): " << typeid(2+3).name() << endl;
  cout << "typeid(int): " << typeid(int).name() << endl;
  cout << "typeid(b+2): " << typeid(b+2).name() << endl;

  const type_info& t1 = typeid(int);
  const type_info& t2 = typeid(a);
  cout << "t1 == t2: " << (t1 == t2) << endl;
  cout << "type_index(t1) == type_index(t2): " << (type_index(t1) == type_index(t2)) << endl;

  cout << "t1.hash_code(): " << t1.hash_code() << endl;
  cout << "t2.hash_code(): " << t2.hash_code() << endl;

  cout << "type_index(t1): " << type_index(t1).name() << endl;
}

struct Base1 {}; // non-polymorphic
struct Derived1 : Base1 {};
 
struct Base2 { virtual void foo() {} }; // polymorphic
struct Derived2 : Base2 {};

void do3()
{
  Derived1 d1;
  Base1& b1 = d1;
  Derived2 d2;
  Base2& b2 = d2;
  cout << "typeid(d1): " << typeid(d1).name() << endl;
  cout << "typeid(b1): " << typeid(b1).name() << endl;
  cout << "typeid(d2): " << typeid(d2).name() << endl;
  cout << "typeid(b2): " << typeid(b2).name() << endl;

  Derived2 *pd2 = nullptr;
  cout << "typeid(pd2): " << typeid(pd2).name() << endl;
  //cout << "typeid(*pd2): " << typeid(*pd2).name() << endl;
}

void do4()
{
  double *mydoubleptr = nullptr;
  try {
    cout << "typeid(*mydoubleptr): " << typeid(*mydoubleptr).name() << endl;

    Derived2* bad_ptr = nullptr;
    cout << "typeid(*bad_ptr): " << typeid(*bad_ptr).name() << endl;
  } catch (const bad_typeid& e) {
    cout << "typeid(*bad_ptr) catch: " << e.what() << endl;
  }
}

void do5()
{
  unordered_map<type_index, string> type_names;
  type_names[type_index(typeid(int))] = "int";
  type_names[type_index(typeid(double))] = "double";
  cout << "type_names[type_index(typeid(2))] = " << type_names[type_index(typeid(2))] << endl;

  //unordered_map<type_info, string> type_names2;
  //type_names2[typeid(int)] = "int";
  //type_names2[typeid(double)] = "double";
  //cout << "type_names2[typeid(2)] = " << type_names2[typeid(2)] << endl;
}

int main(int argc, char const *argv[])
{
  do1();
  return 0;
}
