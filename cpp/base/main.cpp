#include <iostream>

class People {
    public:
        int *pn;
        People() {
            pn = new int(10);
            std::cout << "People create" << std::endl;
        }
        virtual ~People() {
            delete pn;
            std::cout << "People delete" << std::endl;
        }

        void print(){
            std::cout << "People print" << std::endl;
        }
};

class Student: public People {
    public:
        long *pm;
        Student() {
            pm = new long(20);
            std::cout << "Student create" << " pn: "<< *pn << " pm: " << *pm << std::endl;
        }
        virtual ~Student() {
            delete pm;
            std::cout << "Student delete" << std::endl;
        }

        // void print(){
        //     std::cout << "Student print" << std::endl;
        // }
};

void test1() {
    People *p = new Student;
    p->print();
    delete p;
}

void test2() {
    Student *s = new Student;
    delete s;
}

int main(int argc, char const *argv[]) {
    test1();
    return 0;
}
