#include <iostream>

using namespace std;

void do1()
{
  int a = 12;
  int b = move(a);

  b = 13;

  cout << "a: " << a << ", b: " << b << endl;
}

void do2()
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

  People(const People &&p)
  {
    size = p.size;
    data = p.data;
    p.data = nullptr;
    cout << "right ref" << endl;
  }

  People& operator=(const People &p)
  {
    size = p.size;
    data = p.data;
    cout << "operator=" << endl;
  }

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

void do3()
{

}

// g++ -std=c++14 -o app rightref.cpp -w -fno-elide-constructors
int main(int argc, char const *argv[])
{
  do3();
  return 0;
}
