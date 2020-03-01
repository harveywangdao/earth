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
    cout << "People object is being created" << endl;
  }

  virtual int getAge() = 0;
  virtual void setAge(int a) = 0;

  virtual ~People()
  {
    cout << "People object is being deleted " << age << endl;
  }
};

class Student: public People
{
public:
  int no;

  Student()
  {
    cout << "Student object is being created whitout param" << endl;
  }

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

  virtual Student operator+(const Student& s)
  {
    Student stu;
    stu.age = this->age + s.age;
    return stu;
  }

  //virtual void add();

  virtual ~Student()
  {
    cout << "Student object is being deleted " << age << endl;
  }
};

int main(int argc, char const *argv[])
{
  Student s1(10,111);
  Student s2(20,222);

  Student s3;
  s3 = s1 + s2;

  cout << s3.age << endl;

  return 0;
}