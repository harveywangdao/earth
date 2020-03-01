#include <cstring>
#include <string>
#include <cstdio>
#include <iostream>
#include <unistd.h>

using namespace std;

class People
{
public:
  int age;

  People()
  {
    cout << "People object is being created1" << endl;
  }

  People(int ag):age(ag)
  {
    cout << "People object is being created2" << endl;
  }

  virtual int getAge()
  {
    cout << "People getAge" << endl;
    return age;
  }
  virtual void setAge(int a);

  virtual ~People()
  {
    cout << "People object is being deleted " << age << endl;
  }
};

void People::setAge(int a)
{
  age = a;
}

class Student: public People
{
public:
  int no;

  Student(int a, int n)
  {
    age = a;
    no = n;
    cout << "Student object is being created" << endl;
  }

  virtual void setNo(int num)
  {
    no = num;
  }

  virtual int getNo()
  {
    return no;
  }

  virtual int getAge()
  {
    cout << "Student getAge" << endl;
    return age;
  }

  virtual ~Student()
  {
    cout << "Student object is being deleted " << age << endl;
  }
};

int main(int argc, char const *argv[])
{
  /*People p1;
  cout<< p1.getAge() << endl;
  p1.setAge(10);
  cout<< p1.getAge() << endl;

  People p2 = People(20);
  cout<< p2.getAge() << endl;
  p2.setAge(30);
  cout<< p2.getAge() << endl;

  People *p3 = new People(40);
  cout<< p3->getAge() << endl;
  p3->setAge(50);
  cout<< p3->getAge() << endl;
  delete p3;*/

  Student *s1 = new Student(10, 111);
  cout<< s1->getAge() << endl;
  s1->setAge(20);
  cout<< s1->getAge() << endl;
  delete s1;

  cout<< endl;

  People *s2 = new Student(30, 111);
  cout<< s2->getAge() << endl;
  s2->setAge(40);
  cout<< s2->getAge() << endl;
  delete s2;

  sleep(2);

  return 0;
}