#include <iostream>
#include <locale>
#include <iomanip>

using namespace std;

void test1() {
	int num = 255;
	std::cout << "Hex (lower): " << std::hex << num << std::endl;            // ff
	std::cout << "Hex (upper): " << std::uppercase << num << std::endl;      // FF
	std::cout << "Hex (0x): " << std::showbase << num << std::endl;          // 0xFF
	std::cout << "Hex (padded): 0x" << std::setw(4) << std::setfill('0') << num << std::endl;  // 0x00FF
}

void test2() {
  char c1[] = "中";
  printf("sizeof(c1): %ld\n", sizeof(c1));
  for (int i = 0; i < sizeof(c1); i++)
  {
    printf("%x\n", c1[i]);
  }
}

void test3() {
	//setlocale(LC_ALL, "chs");
	//wchar_t wt[] = L"a中";
	//wcout << wt << endl;
  //cout << "sizeof(wt): " << sizeof(wt) << endl;

  //wstring s1 = L"a中";
  //cout << "s1.length(): " << s1.length() << s1.capacity() << endl;

	char c1[] = "中";
	cout << "sizeof(c1): " << sizeof(c1) << endl;
	for (int i = 0; i < sizeof(c1); i++)
	{
		int num = int(c1[i]);
		cout << hex << num << endl;
	}
}

int main()
{
	test3();
	return 0;
}
