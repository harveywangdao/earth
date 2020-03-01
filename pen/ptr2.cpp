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

int main(int argc, char const *argv[])
{
  func1();
  func2();

  int *pp;

  if (nullptr == 0)
  {
    cout << "FFFF" << endl;
  } else {
    cout << "EEEE" << endl;
  }

  cout << "app stop" << endl;

  return 0;
}