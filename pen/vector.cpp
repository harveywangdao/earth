#include <iostream>
#include <vector>

using namespace std;

void printVector(vector<int> v)
{
  for (int i = 0; i < v.size(); ++i)
  {
    cout << v[i];
    if (i != v.size()-1)
    {
      cout << ",";
    }
  }

  cout << endl;
}

void func1()
{
  vector<int> v1;
  vector<int> v2(10);
  vector<int> v3(10, 5);
  vector<int> v4(v3);

  printVector(v1);
  printVector(v2);
  printVector(v3);
  printVector(v4);

  for (int i = 0; i < 10; ++i)
  {
    v1.push_back(i);
  }

  printVector(v1);

  for (vector<int>::iterator iter = v1.begin(); iter != v1.end(); ++iter)
  {
    cout << *iter << ",";
  }
  cout << endl;

  v1.pop_back();
  cout << "after v1 pop_back" << endl;
  printVector(v1);

  v1.clear();
  cout << "after v1 clear" << endl;
  printVector(v1);

  cout << "v1.empty() " << v1.empty() << endl;
  cout << "v1.size() " << v1.size() << endl;
  cout << "v1.capacity() " << v1.capacity() << endl;
  cout << "v1.max_size() " << v1.max_size() << endl;
}

int main(int argc, char const *argv[])
{
  func1();

  return 0;
}
