#include <iostream>
#include <fstream>

//#include <string>
#include <cstring>

using namespace std;

int main(int argc, char const *argv[])
{
  char data[128];
  const char *msg = "this is test file.";

  ofstream outfile;
  outfile.open("temp.txt");    // ios::in | ios::out
  outfile << msg;

  cout << outfile.tellp() << endl;
  outfile.seekp(0, ios::beg);
  cout << outfile.tellp() << endl;

  outfile.close();

  cout << endl;

  ifstream infile;
  infile.open("temp.txt");

  /*infile >> data;
  cout << data << endl;
  infile >> data;
  cout << data << endl;*/

  cout << infile.tellg() << endl;

  while (!infile.eof())
  {
    //infile.getline(data, sizeof(data));
    //cout << data << endl;

    //char c;
    //infile.get(c);
    //cout << c << endl;

    memset(data, 0, sizeof(data));
    infile.read(data, sizeof(data));
    cout << data << endl;
  }
  cout << endl;

  cout << infile.tellg() << endl;
  cout << infile.eof() << endl;
  cout << endl;

  infile.clear();
  cout << infile.tellg() << endl;
  cout << infile.eof() << endl;
  cout << endl;

  infile.seekg(0, infile.beg);
  cout << infile.tellg() << endl;
  cout << infile.eof() << endl;
  cout << endl;

  infile.seekg(0, infile.end);
  cout << infile.tellg() << endl;
  cout << infile.eof() << endl;

  infile.close();

  return 0;
}