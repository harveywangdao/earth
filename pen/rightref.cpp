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

  People(const People &p)
  {
    age = p.age;
    cout << "People object is being created by copy with & " << age << endl;
  }

  People(const People &&p)
  {
    age = p.age;
    cout << "People object is being created by copy with && " << age << endl;
  }

  People& operator=(const People &p)
  {
    age = p.age;
    cout << "People object is being created by operator= " << age << endl;
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

People getPeople()
{
  People p(111);
  cout << "creating people" << endl;
  return p;         // 值拷贝,离开作用域会调用析构
}

void setPeople(People p)
{
  p.age = 2;
  // 对象离开作用域会析构
}


void setPeople2(People& p)
{
  p.age = 2;
  // 引用离开作用域不会析构
}

// g++ -std=c++14 -o app rightref.cpp -w -fno-elide-constructors
void func1()
{
  People p1(222);
  cout << "getPeople start " << p1.age << endl;
  p1 = getPeople();
  cout << "getPeople end " << p1.age << endl;

  People p2(333);
  cout << "setPeople start " << p2.age << endl;
  setPeople(p2);
  cout << "setPeople end " << p2.age << endl;

  People p3(444);
  cout << "setPeople start " << p3.age << endl;
  setPeople2(p3);
  cout << "setPeople end " << p3.age << endl;
}

void func2()
{
  People p1(11);
  People p2(22);
  p2 = p1;    // 值传递

  cout << "p1.age = " << p1.age << endl;
  cout << "p2.age = " << p2.age << endl;

  p2.age = 33;
  cout << "p1.age = " << p1.age << endl;
  cout << "p2.age = " << p2.age << endl;
}

void func3()
{
  People p1(33);
  People p2 = p1;  // 值传递
  p2.age = 44;

  cout << "p1.age = " << p1.age << endl;
  cout << "p2.age = " << p2.age << endl; 
}

// 移动构造函数

void func4()
{
  People &&p1 = getPeople();
  //People p1 = getPeople();
  cout << "p1.age = " << p1.age << endl;
}

void func5()
{
  int a = 4;
  int &b = a;
  //int &c = 5;  // 编译不过
  int &&c = 5;

  cout << "a = " << a << endl;
  cout << "b = " << b << endl;
  cout << "c = " << c << endl;
}

void swap(People &p1, People &p2)
{
  People temp(p1);
  p1 = p2;
  p2 = temp;
}

void swap2(People &p1, People &p2)
{
  People temp(move(p1));
  p1 = move(p2);
  p2 = move(temp);
}

void func6()
{
  People p1(11);
  People p2(22);

  //swap(p1, p2);
  swap2(p1, p2);

  cout << "p1.age = " << p1.age << endl;
  cout << "p2.age = " << p2.age << endl;
}

void func7()
{
  People p1(11);
  People p2(p1);
  People p3(move(p1));

  cout << "p1.age = " << p1.age << endl;
  cout << "p2.age = " << p2.age << endl;
  cout << "p3.age = " << p3.age << endl;
}

int main(int argc, char const *argv[])
{
  //func1();
  //func2();
  //func3();
  //func4();
  //func5();
  func6();
  //func7();

  return 0;
}
