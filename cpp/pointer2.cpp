#include <iostream>
#include <memory>

using namespace std;

struct A;
struct B;

struct A {
    shared_ptr<B> bptr;
    ~A() {
        cout << "A is delete" << endl;
    }
};

struct B {
    shared_ptr<A> aptr;
    ~B() {
        cout << "B is delete " << endl;
    }
};

void func1(){
    shared_ptr<A> ap(new A);
    shared_ptr<B> bp(new B);

    cout << ap.use_count() << endl;
    cout << bp.use_count() << endl;

    ap->bptr = bp;
    bp->aptr = ap;

    cout << ap.use_count() << endl;
    cout << bp.use_count() << endl;

    cout << "AAAA" << endl;
}

struct C;
struct D;

struct C {
    weak_ptr<D> dptr;
    ~C() {
        cout << "C is delete" << endl;
    }
};

struct D {
    weak_ptr<C> cptr;
    ~D() {
        cout << "D is delete " << endl;
    }
};

void func2(){
    shared_ptr<C> cp(new C);
    shared_ptr<D> dp(new D);

    cout << cp.use_count() << endl;
    cout << dp.use_count() << endl;

    cp->dptr = dp;
    dp->cptr = cp;

    cout << cp.use_count() << endl;
    cout << dp.use_count() << endl;

    cout << "AAAA" << endl;
}

void do3()
{
  if (nullptr == 0)
  {
    cout << "nullptr == 0" << endl;
  }
  else 
  {
    cout << "nullptr != 0" << endl;
  }
}

void do4()
{
  const int a = 10;
  const int *p1 = &a;
  // int *p2 = &a; // invalid conversion from const int* to int*
  int *p3 = (int*)&a;

  cout << "&a: " << &a << endl;
  cout << "p1: " << p1 << endl;
  cout << "p3: " << p3 << endl;

  cout << "a: " << a << endl;
  cout << "*p3: " << *p3 << endl;
  (*p3)++; // 结果未定义
  cout << "a: " << a << endl;
  cout << "*p3: " << *p3 << endl;
}

void do5()
{
  int n1 = 10;
  int *p1 = &n1;
  //std::unique_ptr<int> up1(p1);  // free(): invalid pointer munmap_chunk(): invalid pointer
  //std::cout << *up1 << std::endl;

  int *p2 = new int;
  *p2 = 20;
  std::unique_ptr<int> up2(p2);
  std::cout << *up2 << std::endl;
  //delete p2; // free(): double free detected in tcache 2

  //std::unique_ptr<int> up3(up2);
  //std::unique_ptr<int> up3 = up2;
  std::unique_ptr<int> up3(std::move(up2));
  //std::cout << *up2 << std::endl;
  std::cout << *up3 << std::endl;

  std::unique_ptr<int> up4 = std::make_unique<int>();
  *up4 = 30;
  std::cout << *up4 << std::endl;

  std::unique_ptr<int[]> up5 = std::make_unique<int[]>(4);
  up5[0] = 40;
  up5[1] = 41;
  up5[2] = 42;
  up5[3] = 43;
  std::cout << up5[0] << std::endl;
  std::cout << up5[1] << std::endl;
  std::cout << up5[2] << std::endl;
  std::cout << up5[3] << std::endl;
}

void do6()
{
  std::shared_ptr<int> sp1(new int);
  *sp1 = 10;
  std::cout << *sp1 << std::endl;

  std::shared_ptr<int> sp2 = std::make_shared<int>();
  *sp2 = 20;
  std::cout << *sp2 << std::endl;
}

void do7()
{
  std::weak_ptr<int> wp1;
  
  auto sp1 = std::make_shared<int>();
  *sp1 = 10;
  wp1 = sp1;
  std::cout << *sp1 << std::endl;
  
  std::shared_ptr<int> sp2 = wp1.lock();
  *sp2 = 20;
  std::cout << *sp2 << std::endl;
  std::cout << wp1.use_count() << std::endl;

  std::cout << std::endl;
  std::weak_ptr<int> wp2;
  {
    auto sp3 = std::make_shared<int>();
    *sp3 = 30;
    wp2 = sp3;
    std::cout << *sp3 << std::endl;
  }
  std::cout << wp2.use_count() << std::endl;
  std::cout << wp2.lock() << std::endl;
}

int main(int argc, char const *argv[])
{
  do7();
  return 0;
}