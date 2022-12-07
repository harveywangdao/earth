#include <iostream>
#include <cstring>

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
    cout << "People default" << ", size: " << size << ", this: " << this << endl;
  }

  //explicit People(int sz):size(sz)
  People(int sz):size(sz)
  {
    cout << "People with size" << ", size: " << size << ", this: " << this << endl;
  }

  //explicit People(const People& p):size(p.size)
  People(const People& p):size(p.size)
  {
    cout << "People left ref" << ", size: " << size << ", this: " << this << endl;
  }

  //explicit People(People&& p):size(std::move(p.size))
  People(People&& p):size(std::move(p.size))
  {
    cout << "People right ref" << ", size: " << size << ", this: " << this << endl;
  }

  People& operator=(const People& p)
  {
    size = p.size;
    cout << "People operator= &" << ", size: " << size << ", this: " << this << endl;
    return *this;
  }

  People& operator=(People&& p)
  {
    size = std::move(p.size);
    cout << "People operator= &&" << ", size: " << size << ", this: " << this << endl;
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

class Sky
{
  char *_data;

public:
  Sky():_data(nullptr)
  {
    cout << "Sky default" << ", data: " << (void *)_data << ", this: " << this << endl;
  }

  Sky(int sz)
  //explicit Sky(int sz)
  {
    _data = new char[sz];
    cout << "Sky with size" << ", data: " << (void *)_data << ", this: " << this << endl;
  }
  //explicit only declarations of constructors and conversion operators
  //Sky& operator=(int sz) 输入和返回类型必须一样

  Sky(const Sky& p)
  {
    _data = new char[std::strlen(p._data)+1];
    std::strcpy(_data, p._data);
    cout << "Sky left ref" << ", data: " << (void *)_data << ", this: " << this << endl;
  }

  Sky(Sky&& p)
  {
    _data = p._data;
    p._data = nullptr;
    cout << "Sky right ref" << ", data: " << (void *)_data << ", this: " << this << endl;
  }

  Sky& operator=(const Sky& p)
  {
    if (&p == this)
    {
      std::cout << "Sky operator= &, same Sky" << std::endl;
      return *this;
    }

    if (_data != nullptr)
    {
      delete[] _data;
      _data = nullptr;
    }

    _data = new char[std::strlen(p._data)+1];
    std::strcpy(_data, p._data);
    cout << "Sky operator= &" << ", data: " << (void *)_data << ", this: " << this << endl;
    return *this;
  }

  Sky& operator=(Sky&& p)
  {
    if (&p == this)
    {
      std::cout << "Sky operator= &&, same Sky" << std::endl;
      return *this;
    }

    if (_data != nullptr)
    {
      delete[] _data;
      _data = nullptr;
    }

    _data = p._data;
    p._data = nullptr;
    cout << "Sky operator= &&" << ", data: " << (void *)_data << ", this: " << this << endl;
    return *this;
  }

  void dothing()
  {
    if (_data == nullptr)
    {
      std::cout << "Sky dothing, data is nullptr" << std::endl;
    }
    else
    {
      std::cout << "Sky dothing, data: " << (void *)_data << std::endl;
    }
  }

  virtual ~Sky()
  {
    cout << "Sky destroy"  << ", this: " << this << endl;
    if (_data != nullptr)
    {
      delete[] _data;
      _data = nullptr;
    }
  }
};

void do6()
{
  Sky s1;
  Sky s2(10);
  Sky s3(s2);

  s2.dothing();
  Sky s4(std::move(s2));
  s2.dothing();
}

void do7()
{
  Sky s5 = 20;
  Sky s6 = s5;
  Sky s7 = std::move(s5);
}

void do8()
{
  std::cout << "赋值 size:" << endl;
  Sky s8;
  s8 = 30;  //先Sky with size,再Sky operator= &&,再释放Sky with size

  std::cout << endl;
  std::cout << "Copy assignment operator:" << endl;
  Sky s9(40);
  s9 = s8;

  std::cout << endl;
  std::cout << "Move assignment operator:" << endl;
  Sky s10(50);
  s10 = std::move(s8);

  std::cout << endl;
}

template<typename T>
void func(T& param)
{
  std::cout << "传入的是左值" << std::endl;
}
template<typename T>
void func(T&& param)
{
  std::cout << "传入的是右值" << std::endl;
}
template<typename T>
void warp1(T&& param)
{
  func(param);  // param变成左值了
}
template<typename T>
void warp2(T&& param)
{
  func(std::forward<T>(param)); // 左值还是左值,右值还是右值
}
void do9()
{
  int num = 2019;
  warp1(num);
  warp1(2019);

  warp2(num);
  warp2(2019);
}

template<typename T>
void func1(T&& param) // && && -> &&
{
  cout << param << endl;
}
void do10()
{
  int&& val = 4;
  func1(val);
}

template<typename T>
void func2(T&& param) // & && -> &
{
  cout << param << endl;
}
void do11()
{
  int num = 2021;
  int& val = num;
  func2(val);
}

template<typename T>
void func3(T& param) // && & -> &
{
  cout << param << endl;
}
void do12()
{
  int&& val = 2021;
  func3(val);
}

template<typename T>
void func4(T& param) // & & -> &
{
  cout << param << endl;
}
void do13()
{
  int num = 2021;
  int& val = num;
  func4(val);
}

template<typename T>
void func5(T&& param) // 无论是左值还是右值都能接收
{
  cout << param << endl;
}
void do14()
{
  int num = 2019;
  func5(num);
  func5(2019);
}

template<typename T>
void func6(T& param)
{
  cout << "传入的是左值" << endl;
}
template<typename T>
void func6(T&& param)
{
  cout << "传入的是右值" << endl;
}
void do15()
{
  int num = 2019;
  func6(num); // void func6(T& param)
  func6(2019);// void func6(T&& param)
}

template<typename T>
void func7(T& param)
{
  cout << param << endl;
}
void do16()
{
  int num = 2019;
  func7(num);
}
void do17()
{
  // func7(2019); // 不能传右值
}

void do18()
{
  Sky s1 = Sky(2);

  s1 = s1;
  s1 = std::move(s1);

  s1 = 3;
}

// g++ -std=c++17 -o app move.cpp -w -fno-elide-constructors
// g++ -std=c++17 -o app move.cpp
int main(int argc, char const *argv[])
{
  do18();
  return 0;
}
