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
  int b = move(a);
  a++;
  b++;
  cout << "a: " << a << ", b: " << b << endl;
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
public:
  int size;
  int *data;

  People()
  {
    size = 0;
    data = nullptr;
    cout << "default" << endl;
  }

  People(int sz)
  {
    size = sz;
    data = new int[size];
    cout << "with size" << endl;
  }

  People(const People &p)
  {
    size = p.size;
    data = new int[p.size];
    cout << "left ref" << endl;
  }

  People(People &&p)
  {
    size = p.size;
    data = p.data;
    p.data = nullptr;
    cout << "right ref" << endl;
  }

  /*People& operator=(const People &p)
  {
    size = p.size;
    data = p.data;
    cout << "operator=" << endl;
  }*/

  virtual ~People()
  {
    if (data != nullptr)
    {
      delete[] data;
      data = nullptr;
    }
    cout << "People destroy" << endl;
  }
};

// g++ -std=c++14 -o app rightref.cpp -w -fno-elide-constructors
int main(int argc, char const *argv[])
{
  do1();
  return 0;
}
