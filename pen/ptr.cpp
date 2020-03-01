#include <iostream>
#include <memory>

using namespace std;

class People
{
public:
  int age;

  People()
  {
    cout << "People object is being created" << endl;
  }

  People(int _age)
  {
    age = _age;
    cout << "People object is being created with age " << _age << endl;
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

void func1()
{
  auto_ptr<People> p(new People());
  p->setAge(22);
  cout << p->getAge() << endl;
}

void func1_5()
{
  auto_ptr<People> p(new People());
  p->setAge(22);
  cout << p->getAge() << endl;

  auto_ptr<People> p2;
  p2 = p;
  //cout << p->getAge() << endl;
  cout << p2->getAge() << endl;
}

void func2()
{
  People *p = new People();

  auto_ptr<People> auto_p1(p);
  auto_p1->setAge(22);
  cout << auto_p1->getAge() << endl;

  auto_ptr<People> auto_p2(p);
  auto_p2->setAge(33);
  cout << auto_p2->getAge() << endl;
  // core dumped
}

void func3()
{
  auto_ptr<People> auto_p1(new People());
  auto_p1->setAge(22);
  cout << auto_p1->getAge() << endl;

  auto_ptr<People> auto_p2(auto_p1);

  auto_p2->setAge(33);
  cout << auto_p2->getAge() << endl;

  auto_p1->setAge(44);             // core dumped
  //cout << auto_p1->getAge() << endl;
}

void func4()
{
  auto_ptr<People> auto_p1(new People());
  auto_p1->setAge(22);
  cout << auto_p1->getAge() << endl;

  auto_ptr<People> auto_p2(new People());
  auto_p2->setAge(33);
  cout << auto_p2->getAge() << endl;

  auto_p1 = auto_p2;
  cout << auto_p1->getAge() << endl;
}

void test1(auto_ptr<People> &ap)
{
  cout << "test1 " << ap->getAge() << endl;
}

void test2(auto_ptr<People> ap)
{
  cout << "test2 " << ap->getAge() << endl;
}

void func5()
{
  auto_ptr<People> auto_p(new People());
  auto_p->setAge(22);
  cout << auto_p->getAge() << endl;

  test1(auto_p);
  cout << "test1 end" << endl;
  cout << auto_p->getAge() << endl;

  test2(auto_p);
  cout << "test2 end" << endl;
  cout << auto_p->getAge() << endl;    // core dumped
}

void func6()
{
  People *p = new People();
  auto_ptr<People> auto_p(p);
  auto_p->setAge(22);
  cout << auto_p->getAge() << endl;

  cout << "p = " << p << endl;
  //cout << "auto_p = " << auto_p << endl;
  cout << "&auto_p = " << &auto_p << endl;
  cout << "auto_p.get() = " << auto_p.get() << endl;

  auto_p.reset(new People());
  cout << "after reset" << endl;
  cout << auto_p->getAge() << endl;
  auto_p->setAge(33);
  cout << auto_p->getAge() << endl;

  auto_ptr<People> auto_p2(auto_p.release());
  cout << auto_p2->getAge() << endl;
  //cout << auto_p->getAge() << endl;     // core dumped

  People *p2 = auto_p2.release();
  cout << p2->getAge() << endl;
  //cout << auto_p2->getAge() << endl;      // core dumped
  delete p2;
}

void func7()
{
  unique_ptr<People> unique_p(new People());
  unique_p->setAge(22);
  cout << unique_p->getAge() << endl;

  unique_ptr<People> unique_p2(new People());
  //unique_p2 = unique_p;      // 编译不过

  unique_p2 = move(unique_p);
  cout << unique_p2->getAge() << endl;
  //cout << unique_p->getAge() << endl;  // core dumped
}

void func8()
{
  unique_ptr<People> p1(new People());
  p1->setAge(22);
  cout << p1->getAge() << endl;

  unique_ptr<People> p2(new People());
  p2->setAge(33);
  cout << p2->getAge() << endl;

  p1.swap(p2);

  cout << p1->getAge() << endl;
  cout << p2->getAge() << endl;

  p1.reset(new People(44));
  cout << p1->getAge() << endl;

  p2.reset();
  //cout << p2->getAge() << endl;    // core dumped

  People *p3 = p1.release();

  delete p3;
}

void func9()
{
  unique_ptr<People> p1 = make_unique<People>(22);
  cout << p1->getAge() << endl;

  shared_ptr<People> p2(new People(33));
  shared_ptr<People> p3 = p2;
  cout << p2->getAge() << endl;
  cout << p3->getAge() << endl;
  p3.reset();
  p2.reset();

  shared_ptr<People> p4 = make_shared<People>(44);
  cout << p4->getAge() << endl;

  shared_ptr<int> p5(new int[10], [](int *p){delete [] p;});
  shared_ptr<int> p6(new int[10], default_delete<int[]>());
}

void func10()
{
  shared_ptr<People> p1(new People(33));
  cout << p1->getAge() << endl;

  weak_ptr<People> p2;
  p2 = p1;

  shared_ptr<People> p3 = p2.lock();
  cout << p3->getAge() << endl;

  shared_ptr<People> p4 = p2.lock();
  cout << p4->getAge() << endl;

  cout << p2.use_count() << endl;
  cout << p3.use_count() << endl;
  cout << p4.use_count() << endl;
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
  //func8();
  //func9();
  //func10();
  func1_5();

  cout << "app stop" << endl;

  return 0;
}