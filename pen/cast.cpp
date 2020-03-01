#include <iostream>

using namespace std;

class People
{
public:
  int age;

  People()
  {
    cout << "People object is being created " << age << endl;
  }

  People(int _age)
  {
    age = _age;
    cout << "People object is being created with age " << age << endl;
  }

  virtual int getAge()
  {
    cout << "People getAge" << endl;
    return age;
  }

  virtual void setAge(int a)
  {
    cout << "People setAge" << endl;
    age = a;
  }

  virtual ~People()
  {
    cout << "People object is being deleted " << age << endl;
  }
};

class Student: public People
{
public:
  int no;

  Student(int _no, int _age):no(_no)
  {
    age = _age;
    cout << "Student object is being created " << age << endl;
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
    cout << "Student object is being deleted " << age << endl;
  }
};

void func1()
{
  const int a = 10;
  int *p = const_cast<int*>(&a);
  *p = 11;

  // 为什么不一样
  cout << "a = " << a << endl;
  cout << "*p = " << *p << endl;
}

void func2()
{
  const int a = 10;
  int *p = (int*)&a;
  *p = 11;

  cout << "a = " << a << endl;
  cout << "*p = " << *p << endl;
}

void print(int *p)
{
  cout << *p << endl;
}

void func3()
{
  const int a = 10;
  print((int*)&a);
  print(const_cast<int*>(&a));
}

void func4()
{
  // 用于类层次结构中基类（父类）和派生类（子类）之间指针或引用的转换。
  // 进行上行转换（把派生类的指针或引用转换成基类表示）是安全的；
  // 进行下行转换（把基类指针或引用转换成派生类表示）时，由于没有动态类型检查，所以是不安全的。
  Student s1(111, 11);
  People *p1 = static_cast<People*>(&s1);
  p1->age = 12;

  cout << "s1.age = " << s1.age << endl;
  cout << "p1.age = " << p1->age << endl;
}

void func5()
{
  Student s1(111, 11);
  People *p1 = &s1;
  p1->age = 12;

  cout << "s1.age = " << s1.age << endl;
  cout << "p1.age = " << p1->age << endl;
}

void func6()
{
  float f = 3.6415;
  int i = static_cast<int>(f);
  int n = f;

  cout << f << endl;
  cout << i << endl;
  cout << n << endl;
}

void func7()
{
  // 从指针类型到一个足够大的整数类型
  // 从整数类型或者枚举类型到指针类型

  int *i;
  char *ptr = "dddd";
  i = reinterpret_cast<int*>(ptr);

  cout << ptr << endl;
  cout << i << endl;
}

void func8()
{
  // 将基类的指针或引用安全地转换成派生类的指针或引用
  // 并用派生类的指针或引用调用非虚函数
  Student s(111, 999);
  People *p = &s;

  // 区别是啥
  Student *ps1 = dynamic_cast<Student*>(p);
  Student *ps2 = (Student*)p;

  cout << p->age << endl;
  
  cout << ps1->age << endl;                  
  cout << ps1->getNo() << endl;              
  cout << ps1->getNoWithoutVirtual() << endl;
  
  cout << ps2->age << endl;
  cout << ps2->getNo() << endl;
  cout << ps2->getNoWithoutVirtual() << endl;
}

int main(int argc, char const *argv[])
{
  //func1();
  //func2();
  //func3();
  //func4();
  //func5();
  //func6();
  //func7();
  func8();

  return 0;
}
