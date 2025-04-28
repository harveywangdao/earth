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
	char c1[] = "中";
	cout << "sizeof(c1): " << sizeof(c1) << endl;
	for (int i = 0; i < sizeof(c1); i++)
	{
		int num = int(c1[i]);
		cout << hex << num << endl;
	}
}

void test4() {
	wchar_t s1[] = L"中";
	//wcout << s1 << endl;
  cout << "sizeof(s1): " << sizeof(s1) << endl;

	for (int i = 0; i < sizeof(s1)/sizeof(wchar_t); i++)
	{
		cout << hex << s1[i] << endl;
	}
	
  wstring s2 = L"中";
	wcout << L"s2: " << s2 << endl;
  cout << "s2.length(): " << s2.length() << endl;
  cout << "s2.capacity(): " << s2.capacity() << endl;
}

int test5() {
	std::wstring ws;
	
	// 输入
	std::wcout << L"请输入宽字符串: ";
	std::getline(std::wcin, ws);
	
	// 输出
	std::wcout << L"你输入的是: " << ws << std::endl;
	
	return 0;
}

int main()
{
	test5();
	return 0;
}
