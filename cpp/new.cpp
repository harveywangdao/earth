#include <iostream>

class People
{
private:
  int num;
public:
  People()
  {
    std::cout << "People constructor" << std::endl;
  }

  People(int n):num(n)
  {
    std::cout << "People constructor with num" << std::endl;
  }

  virtual ~People()
  {
    std::cout << "People destructor" << std::endl;
  }

  void dothing()
  {
    std::cout << "People dothing" << ", num: " << num << std::endl;
  }
};

void do1()
{
  People *p1 = new People;
  delete p1;
  std::cout << "p1: " << p1 << std::endl;

  People *p2 = new People[3];
  delete[] p2;
  std::cout << "p2: " << p2 << std::endl;
}

void do2()
{
  People *p1 = new People;
  std::cout << "p1: " << p1 << std::endl;
  delete[] p1; // Segmentation fault
}

void do3()
{
  People *p1 = new People[3];
  std::cout << "p1: " << p1 << std::endl;
  delete p1; // 调用一个析构后报Error
}

void do4()
{
  People *p1 = new People;
  std::cout << "p1: " << p1 << std::endl;
  //p1.dothing();
  p1->dothing();
  (*p1).dothing();
  delete p1;
}

void do5()
{
  People *p1 = new People(10);
  std::cout << "p1: " << p1 << std::endl;
  p1->dothing();
  p1++;
  p1->dothing(); // 位置区域的内存,但是不会报错
  p1++;
  p1->dothing(); // 位置区域的内存,但是不会报错
  delete p1;
}

void do6()
{
  People *p1 = new People[3]{(10),(20),(30)};
  std::cout << "p1: " << p1 << std::endl;
  p1->dothing();
  p1++;
  p1->dothing();
  p1++;
  p1->dothing();

  p1++;
  p1->dothing(); // 位置区域的内存,但是不会报错

  delete p1;
}

void do7()
{
  People *p1 = new People[3]{(10),(20),(30)};
  std::cout << "p1: " << p1 << std::endl;
  p1->dothing();
  p1++;
  p1->dothing();
  delete p1;
}

int main(int argc, char const *argv[])
{
  do3();
  return 0;
}