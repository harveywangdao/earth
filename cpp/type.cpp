#include <iostream>
#include <typeinfo>
#include <typeindex>
#include <unordered_map>

using namespace std;

// g++ -std=c++11 -o app type.cpp
void do1()
{
  int arr1[8] = {1};
  auto arr2 = arr1;
  auto arr3 = {10,20,30};  // std::initializer_list

  cout << "arr1: " << arr1 << endl;
  cout << "typeid arr1: " << typeid(arr1).name() << endl;
  cout << "sizeof arr1: " << sizeof(arr1) << endl;

  cout << "arr2: " << arr2 << endl;
  cout << "typeid arr2: " << typeid(arr2).name() << endl;
  cout << "sizeof arr2: " << sizeof(arr2) << endl;

  cout << "typeid arr3: " << typeid(arr3).name() << endl;
  cout << "sizeof arr3: " << sizeof(arr3) << endl;
  for (auto &v : arr3)
  {
    cout << v << endl;
  }

  //auto arr4[] = {1};
  //auto arr5[8] = {1};
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

void do6()
{
  int a = 10;
  decltype(a) b;
  int &c = a;
  decltype(c) d = b;
  decltype((a)) e = a;   // 加上括号表示引用
  decltype(auto) f = (a);// 加上括号表示引用
  cout << "typeid(a): " << typeid(a).name() << endl;
  cout << "typeid(b): " << typeid(b).name() << endl;
  cout << "typeid(c): " << typeid(c).name() << endl;
  cout << "typeid(d): " << typeid(d).name() << endl;
  cout << "typeid(e): " << typeid(e).name() << endl;
  cout << "typeid(f): " << typeid(f).name() << endl;

  cout << a << endl;
  e++;
  cout << a << endl;
  f++;
  cout << a << endl;
}

template<typename T, typename U>
//auto add(T t, U u) -> decltype(t + u)
auto add(T t, U u)
{
  return t + u;
}

void do7()
{
  int a = 10, b = 20;
  auto c = add<int,int>(a, b);
  cout << "c: " << c << endl;
  cout << "typeid(c): " << typeid(c).name() << endl;
}

template<auto n>
auto fff() -> std::pair<decltype(n), decltype(n)>
{
  return {n, n};
}

void do8()
{
  auto a = 1 + 2;
  auto b = add(1, 1.2);
  static_assert(std::is_same_v<decltype(a), int>);
  static_assert(std::is_same_v<decltype(b), double>);

  auto [v, w] = fff<13>();
  cout << "v: " << v << endl;
  cout << "w: " << w << endl;

  auto d = {1, 2}; // std::initializer_list<int>
  //auto e{1, 2};
  auto m{5};
  //decltype(auto) z = { 1, 2 }
}

int main(int argc, char const *argv[])
{
  do8();
  return 0;
}
