#include <iostream>

using namespace std;

void do1()
{
  int&& n1 = 10;
  int&& n2 = std::move(n1);
  //int&& n3 = n1;

  int x = 20;

  int&& y1 = x++;
  //int&& y2 = ++x;

  //int& y3 = x++;
  int& y4 = ++x;

  int y5 = x++;
  int y6 = ++x;

  int z = 40;
  // x++是右值
  // ++x是左值
  ++z = z++;     // 左值 = 右值
  std::cout << "z = " << z << std::endl;
  // x++ = ++x;  // 右值 != 右值
  ++z = ++z;     // 左值 = 左值
  std::cout << "z = " << z << std::endl;
  // x++ = ++x   // 右值 != 左值

  std::cout << "x = " << x << std::endl;
  std::cout << "y1 = " << y1 << std::endl;
  std::cout << "y4 = " << y4 << std::endl;
  std::cout << "y5 = " << y5 << std::endl;
  std::cout << "y6 = " << y6 << std::endl;
}

void do2()
{
  int a = 10;
  int b = std::move(a);
  int&& c = 20;

  a++;
  b++;
  c++;

  std::cout << "a: " << a << std::endl;
  std::cout << "b: " << b << std::endl;
  std::cout << "c: " << c << std::endl;
}

void do3()
{
  int a = 10;
  int &b = a;

  //int &c = a * 2;
  const int &d = a * 2;

  int &&e = a * 2;
  //int &&f = a;
  int &&f = move(a);

  e = 30;

  cout << a << endl << b << endl << d << endl << e << endl << f << endl;
  cout << &a << endl << &b << endl << &d << endl << &e << endl << &f << endl;
}

class People
{
  int size;

public:
  People():size(0)
  {
    cout << "default" << ", size: " << size << ", this: " << this << endl;
  }

  //explicit People(int sz):size(sz)
  People(int sz):size(sz)
  {
    cout << "with size" << ", size: " << size << ", this: " << this << endl;
  }

  //explicit People(const People& p):size(p.size)
  People(const People& p):size(p.size)
  {
    cout << "left ref" << ", size: " << size << ", this: " << this << endl;
  }

  //explicit People(People&& p):size(std::move(p.size))
  People(People&& p):size(std::move(p.size))
  {
    cout << "right ref" << ", size: " << size << ", this: " << this << endl;
  }

  People& operator=(const People& p)
  {
    size = p.size;
    cout << "operator= &" << ", size: " << size << ", this: " << this << endl;
    return *this;
  }

  People& operator=(People&& p)
  {
    size = std::move(p.size);
    cout << "operator= &&" << ", size: " << size << ", this: " << this << endl;
    return *this;
  }

  virtual ~People()
  {
    cout << "People destroy" << ", size: " << size << ", this: " << this << endl;
  }
};

void do4()
{
  People p1;
  People p2(2);
  People p3(p1);
  People p4(std::move(p1));

  People p5;
  p5 = p2;

  People p6;
  p6 = std::move(p2);

  People p7 = 7;
  People p8 = p7;
  People p9 = std::move(p7);
}

People m1()
{
  People temp(10);
  std::cout << "&temp: " << &temp << std::endl;
  return temp;
}

/*People& m2()
{
  People temp(20);
  std::cout << "&temp: " << &temp << std::endl;
  return temp;
}*/

/*People&& m3()
{
  People temp(30);
  std::cout << "&temp: " << &temp << std::endl;
  //return temp;
  return std::move(temp);
}*/

void do5()
{
  {
    People p1 = m1();
    std::cout << "&p1: " << &p1 << std::endl; 
  }

  /*{
    std::cout << std::endl;
    People p2 = m2();
    std::cout << "&p2: " << &p2 << std::endl;
  }*/

  /*{
    std::cout << std::endl;
    People p3 = m3();
    std::cout << "&p3: " << &p3 << std::endl; 
  }*/
}

// g++ -std=c++17 -o app move.cpp -w -fno-elide-constructors
// g++ -std=c++17 -o app move.cpp
int main(int argc, char const *argv[])
{
  do5();
  return 0;
}
