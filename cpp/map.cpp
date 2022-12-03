#include <iostream>
#include <map>

using namespace std;

void printMap(map<int, string> m)
{
  map<int, string>::iterator iter;
  for (iter = m.begin(); iter != m.end(); ++iter)
  {
    cout << iter->first << " " << iter->second << " ";
  }
  cout << endl;
}

void do1()
{
  map<int, string> m1;
  pair<map<int, string>::iterator, bool> ret;

  m1.insert(pair<int, string>(1, "a"));
  ret = m1.insert(pair<int, string>(1, "aa"));
  if (ret.second == false)
  {
    cout << "already existed " << ret.first->first << " " << ret.first->second << endl;
  }

  m1.insert(map<int, string>::value_type(2, "b"));
  m1.insert(map<int, string>::value_type(2, "bb"));
  m1[3] = "c";
  m1[4] = "d";
  m1[5] = "e";
  m1[5] = "ee";
  printMap(m1);

  map<int, string>::iterator iter;
  iter = m1.find(3);
  if (iter != m1.end())
  {
    cout << iter->first << " " << iter->second << endl;
    m1.erase(iter);
  }
  printMap(m1);

  cout << m1[5] << endl;

  cout << "m1.erase(2) " << m1.erase(2) << endl;
  cout << "m1.erase(21) " << m1.erase(21) << endl;
  printMap(m1);

  cout << "m1.size() " << m1.size() << endl;
}

int main(int argc, char const *argv[])
{
  do1();
  return 0;
}