#include <iostream>

using namespace std;

class People
{
public:
  int age;

  People()
  {
    age = 10;
  }

  People(int _age)
  {
    age = _age;
  }

  virtual int getAge()
  {
    return age;
  }

  virtual void setAge(int a)
  {
    age = a;
  }

  virtual ~People()
  {
  }
};

class Student: public People
{
public:
  int no;

  Student(int _no, int _age):no(_no)
  {
    age = _age;
  }

  virtual int getNo()
  {
    return no;
  }

  int getNoWithoutVirtual()
  {
    return no;
  }

  virtual ~Student()
  {
  }
};

// const_cast 指针
void do1()
{
  int a = 10;
  const int b = a;
  int *p1 = const_cast<int*>(&b);

  cout << "&a = " << &a << endl;
  cout << "&b = " << &b << endl;
  cout << "p1 = " << p1 << endl;

  cout << "before a = " << a << ", b = " << b << endl;
  *p1 = 11;
  cout << "after a = " << a << ", b = " << b << endl;
  cout << "*p1 = " << *p1 << endl;
}

// const_cast 指针 无法修改b
void do1_2()
{
  const int b = 10;
  int *p1 = const_cast<int*>(&b);

  cout << "&b = " << &b << endl;
  cout << "p1 = " << p1 << endl;

  cout << "before b = " << b << endl;
  *p1 = 11;
  cout << "after b = " << b << endl;
  cout << "*p1 = " << *p1 << endl;
}

// const_cast 引用 可以修改a
void do2()
{
  int a = 10;
  const int& b = a;
  int &c = const_cast<int&>(b);

  cout << "&a = " << &a << endl;
  cout << "&b = " << &b << endl;
  cout << "&c = " << &c << endl;

  cout << "before a = " << a << ", b = " << b << ", c = " << c << endl;
  c++;
  cout << "after a = " << a << ", b = " << b << ", c = " << c << endl;
}

// const_cast 引用 无法修改a
void do2_2()
{
  const int a = 10;
  const int& b = a;
  int &c = const_cast<int&>(b);

  cout << "&a = " << &a << endl;
  cout << "&b = " << &b << endl;
  cout << "&c = " << &c << endl;

  cout << "before a = " << a << ", b = " << b << ", c = " << c << endl;
  c++;
  cout << "after a = " << a << ", b = " << b << ", c = " << c << endl;
}

// const_cast 引用 可以修改b
void do2_3()
{
  int a = 10;
  const int b = a;
  int& c = const_cast<int&>(b);

  cout << "&a = " << &a << endl;
  cout << "&b = " << &b << endl;
  cout << "&c = " << &c << endl;

  cout << "before a = " << a << ", b = " << b << ", c = " << c << endl;
  c++;
  cout << "after a = " << a << ", b = " << b << ", c = " << c << endl;
}

// const_cast 引用 无法修改b
void do2_4()
{
  const int a = 10;
  const int b = a;
  int& c = const_cast<int&>(b);

  cout << "&a = " << &a << endl;
  cout << "&b = " << &b << endl;
  cout << "&c = " << &c << endl;

  cout << "before a = " << a << ", b = " << b << ", c = " << c << endl;
  c++;
  cout << "after a = " << a << ", b = " << b << ", c = " << c << endl;
}

void do3()
{
  const int a = 10;
  int& b = const_cast<int&>(a);
  b++;
  int& c = const_cast<int&>(a);
  cout << "a = " << a << endl;
  cout << "b = " << b << endl;
  cout << "c = " << c << endl;

  const People p;
  People& p2 = const_cast<People&>(p);
  cout << "p.age = " << p.age << endl;
  cout << "p2.age = " << p2.age << endl;
  p2.age++;
  cout << "p.age = " << p.age << endl;
  cout << "p2.age = " << p2.age << endl;
}

void do4()
{
  // 用于类层次结构中基类（父类）和派生类（子类）之间指针或引用的转换。
  // 进行上行转换（把派生类的指针或引用转换成基类表示）是安全的；
  // 进行下行转换（把基类指针或引用转换成派生类表示）时，由于没有动态类型检查，所以是不安全的。
  Student s1(111, 11);
  People *p1 = static_cast<People*>(&s1);
  p1->age = 12;
  cout << "s1.age = " << s1.age << endl;
  cout << "p1.age = " << p1->age << endl;

  Student s2(111, 11);
  People *p2 = &s2;
  p2->age = 12;
  cout << "s2.age = " << s2.age << endl;
  cout << "p2.age = " << p2->age << endl;
}

void do5()
{
  float a = 3.6415;
  int b = static_cast<int>(a);
  int c = a;
  int d = (int)a;

  cout << a << endl;
  cout << b << endl;
  cout << c << endl;
  cout << d << endl;
}

void do6()
{
  // 从指针类型到一个足够大的整数类型
  // 从整数类型或者枚举类型到指针类型

  int *i;
  char ptr[] = "dddd";
  i = reinterpret_cast<int*>(ptr);

  cout << ptr << endl;
  cout << &ptr << endl;
  cout << i << endl;
}

void do7()
{
  // 将基类的指针或引用安全地转换成派生类的指针或引用
  // 并用派生类的指针或引用调用非虚函数
  Student s(123456, 20);
  People *p = &s;

  // 区别是啥
  Student *ps1 = dynamic_cast<Student*>(p);
  Student *ps2 = (Student*)p;

  cout << p->age << endl << endl;
  
  cout << ps1->age << endl;                  
  cout << ps1->getNo() << endl;              
  cout << ps1->getNoWithoutVirtual() << endl << endl;
  
  cout << ps2->age << endl;
  cout << ps2->getNo() << endl;
  cout << ps2->getNoWithoutVirtual() << endl;
}

int main(int argc, char const *argv[])
{
  do7();
  return 0;
}
