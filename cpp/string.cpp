#include <string>
#include <iostream>
#include <algorithm>

using namespace std;

void func1()
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

void func2()
{
  string s1 = "123456";
  string s2 = s1; 

  cout << "before:s1 " << s1 << endl;
  cout << "before:s2 " << s2 << endl;

  s1[0] = 'f';

  cout << "after:s1 " << s1 << endl;
  cout << "after:s2 " << s2 << endl;
}

void func3()
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

void func4()
{
  string s1 = "abc";
  s1 = "123" + s1;
  cout << s1 << endl;
         
  const char *prefix;
  prefix = "ABC";
  s1 = prefix + s1;
  cout << s1 << endl;
}

int main(int argc, char const *argv[])
{
  //func1();
  //func2();
  //func3();
  func4();

  return 0;
}