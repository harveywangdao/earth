#include <iostream>

void do1()
{
  //std::cout << "sizeof char8_t: " << sizeof(char8_t) << std::endl;     // 1
  std::cout << "sizeof char16_t: " << sizeof(char16_t) << std::endl;   // 2
  std::cout << "sizeof char32_t: " << sizeof(char32_t) << std::endl;   // 4
  std::cout << "sizeof wchar_t: " << sizeof(wchar_t) << std::endl;     // 4
}

void do2()
{
  char c1 = u8'我';
  char16_t c2 = u'我';
  char32_t c3 = U'我';

  std::cout << "sizeof c1: " << sizeof(c1) << std::endl;     // 1
  std::cout << "sizeof c2: " << sizeof(c2) << std::endl;     // 2
  std::cout << "sizeof c3: " << sizeof(c3) << std::endl;     // 4

  std::cout << "c1: " << c1 << std::endl;
  std::cout << "c2: " << c2 << std::endl;
  std::cout << "c3: " << c3 << std::endl;

  char s1[] = "我是猪";
  char16_t s2[] = u"我是猪";    // 最后一个元素是0
  char32_t s3[] = U"我是猪";
  char s4[] = u8"我是猪";
  wchar_t s5[] = L"我是猪";

  std::cout << "sizeof s1: " << sizeof(s1) << std::endl;     // 10
  std::cout << "sizeof s2: " << sizeof(s2) << std::endl;     // 8
  std::cout << "sizeof s3: " << sizeof(s3) << std::endl;     // 16
  std::cout << "sizeof s4: " << sizeof(s4) << std::endl;     // 10
  std::cout << "sizeof s5: " << sizeof(s5) << std::endl;     // 16

  std::wcout << "s5: " << s5 << std::endl;
  std::wcout << L"我是猪" << std::endl;

  for (int i = 0; i < 4; ++i)
  {
    std::cout << s2[i] << std::endl;
  }
}

int main(int argc, char const *argv[])
{
  do2();
  return 0;
}