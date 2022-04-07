#include <cstring>
#include <string>
#include <cstdio>
#include <iostream>
#include <unistd.h>

using namespace std;

class People1
{
public:
  int age1;

  People1()
  {
    cout << "People1 object is being created" << endl;
  }

  virtual int getAge()
  {
    cout << "People1 getAge" << endl;
    return age1;
  }

  virtual void setAge(int a)
  {
    cout << "People1 setAge" << endl;
    age1 = a;
  }

  virtual ~People1()
  {
    cout << "People1 object is being deleted " << age1 << endl;
  }
};

class People2
{
public:
  int age2;

  People2()
  {
    cout << "People2 object is being created" << endl;
  }

  virtual int getAge()
  {
    cout << "People2 getAge" << endl;
    return age2;
  }
  
  virtual void setAge(int a)
  {
    cout << "People2 setAge" << endl;
    age2 = a;
  }

  virtual ~People2()
  {
    cout << "People2 object is being deleted " << age2 << endl;
  }
};

class Student: public People1, public People2
{
public:
  int no;
  int age;

  Student(int a, int n)
  {
    age = a;
    no = n;
    cout << "Student object is being created" << endl;
  }

  virtual void setNo(int n)
  {
    no = n;
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

  virtual void setAge(int a)
  {
    cout << "Student setAge" << endl;
    age = a;
  }

  virtual ~Student()
  {
    cout << "Student object is being deleted " << age << endl;
  }
};

int main(int argc, char const *argv[])
{
  Student *s1 = new Student(10, 111);
  cout<< s1->getAge() << endl;
  s1->setAge(20);
  cout<< s1->getAge() << endl;
  delete s1;

  return 0;
}