#include <iostream>
#include <string>
#include <string_view>
#include <algorithm>

using namespace std;

void do1()
{
  string s1 = "123456";
  string s2 = s1;

  cout << "before:s1 " << s1 << endl;
  cout << "before:s2 " << s2 << endl;

  char* ptr = const_cast<char*>(s1.c_str());
  *ptr = 'f';

  cout << "after:ptr " << ptr << endl;
  cout << "after:s1 " << s1 << endl;
  cout << "after:s2 " << s2 << endl;
}

void do2()
{
  string s1 = "123456";
  string s2 = s1; 

  cout << "before:s1 " << s1 << endl;
  cout << "before:s2 " << s2 << endl;

  s1[0] = 'f';

  cout << "after:s1 " << s1 << endl;
  cout << "after:s2 " << s2 << endl;
}

void do3()
{
  string s1 = "123456";
  string s2(s1);
  string s3("123456");
  string s4("123456789", 6);
  string s5(10, 'a');

  cout << "s1 " << s1 << endl;
  cout << "s2 " << s2 << endl;
  cout << "s3 " << s3 << endl;
  cout << "s4 " << s4 << endl;
  cout << "s5 " << s5 << endl;

  cout << "s1.size " << s1.size() << endl;
  cout << "s1.length " << s1.length() << endl;
  cout << "s1.capacity " << s1.capacity() << endl;

  s1.append("abc");
  cout << "append s1 " << s1 << endl;
  s1.push_back('d');
  cout << "push_back s1 " << s1 << endl;
  s1.swap(s2);
  cout << "swap s1 " << s1 << endl;
  cout << "swap s2 " << s2 << endl;

  s1.insert(2, "ABC");
  cout << "insert s1 " << s1 << endl;

  s1.erase(2, 3);
  cout << "erase s1 " << s1 << endl;

  s1.replace(2, 3, "ABC");
  cout << "replace s1 " << s1 << endl;

  reverse(s1.begin(), s1.end());
  cout << "reverse s1 " << s1 << endl;

  cout << "s1.find(\"AB\") " << s1.find("AB") << endl;

  cout << "s1.substr(2, 3) " << s1.substr(2, 3) << endl;

  const char *ps1 = s1.data();
  cout << "ps1 " << ps1 << endl;

  cout << "s1.compare(\"ABC\") " << s1.compare("ABC") << endl;

  for (string::iterator p = s1.begin(); p != s1.end(); ++p)
  {
    cout << *p << endl;
  }

  for (string::reverse_iterator p = s1.rbegin(); p != s1.rend(); ++p)
  {
    cout << *p << endl;
  }

  cout << "empty s1 " << s1.empty() << endl;
  s1.clear();
  cout << "clear s1 " << s1 << endl;
  cout << "empty s1 " << s1.empty() << endl;

}

void do4()
{
  string s1 = "abc";
  s1 = "123" + s1;
  cout << s1 << endl;
  
  const char *prefix;
  prefix = "ABC";
  s1 = prefix + s1;
  cout << s1 << endl;
}

void do5()
{
  std::string s1 = "abc";
  std::string s2("abc");

  std::cout << "size: " << s1.size() << std::endl;
  std::cout << "length: " << s1.length() << std::endl;
  std::cout << "empty: " << s1.empty() << std::endl;
  std::cout << "capacity: " << s1.capacity() << std::endl;

  std::string s3(4, '=');
  std::cout << "s3: " << s3 << std::endl;

  std::string s4(s1);
  std::cout << "s4: " << s4 << std::endl;

  std::string s5 = s1 + s3;
  std::cout << "s5: " << s5 << std::endl;

  s1[0] = '1';
  std::cout << "s1: " << s1 << std::endl;

  s1.push_back('2');
  std::cout << "s1: " << s1 << std::endl;
}

void do6()
{
  std::string_view s1 = "abc";
  std::cout << "s1: " << s1 << std::endl;
  std::cout << "s1[0]: " << s1[0] << std::endl;
  const char *p = s1.data();
  //char *p = s1.data();
  //std::cout << "data: " << p << std::endl;
  //p++;
  //std::cout << "data: " << *p << std::endl;
  //*p = 'd';

  char *p1 = const_cast<char*>(p);
  std::cout << "*p1: " << *p1 << std::endl;
  *p1 = 'd';
}

void do7()
{
  std::string s("abc");
  std::string_view s1(s);

  std::cout << "s1: " << s1 << std::endl;
  
  const char *p = s1.data();
  char *p1 = const_cast<char*>(p);
  std::cout << "*p1: " << *p1 << std::endl;
  *p1 = 'd';
  std::cout << "s1: " << s1 << std::endl;
}

void do8()
{
  char str[] = "abc";
  std::string_view s1(str);
  std::cout << "s1: " << s1 << std::endl;

  std::cout << "str: " << (void*)str << std::endl;
  std::cout << "s1: " << (void*)s1.data() << std::endl;

  char *p1 = const_cast<char*>(s1.data());
  *p1 = 'd';
  std::cout << "s1: " << s1 << std::endl;
}

void do9()
{
  std::string_view s1;
  {
    std::string s2("abc");
    s1 = s2.data();
    std::cout << "data: " << (void*)s2.data() << std::endl;
    std::cout << "c_str: " << (void*)s2.c_str() << std::endl;
    std::cout << "s1: " << s1 << std::endl;
  }
  std::cout << "s1: " << s1 << std::endl; //未知错误
}

int main(int argc, char const *argv[])
{
  do9();
  return 0;
}
