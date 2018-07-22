#include <iostream>
#include "PersistenceMap.h"
int main() {
    std::cout << "Hello, World!" << std::endl;
    PersistenceMap kv;
    kv.Set("C","c");
    kv.Set("B","b");
    kv.Set("A","a");

    auto r = kv.Get("C");

    std::cout << r << std::endl;

    auto r2 = kv.Get("B");
    std::cout << r2 << std::endl;

    auto r3 = kv.Get("A");
    std::cout << r3 << std::endl;

//    ofstream out;
//    out.open("/Users/liiiyu/Documents/project/c++/PersistenceMap/kv.txt",ios::binary|ios::out|ios::app);
//    out.seekp(0,std::ios::end);
//    auto offset = out.tellp();
//    std::cout <<offset << std::endl;
//    out.write("wo",2);
//    out.close();
//
//    ifstream in;
//    in.open("/Users/liiiyu/Documents/project/c++/PersistenceMap/kv.txt",ios::binary|ios::in);
//    in.seekg(2, ios::beg);
//    char buf[2];
//    in.read(buf,2);
//
//    std::cout << buf << std::endl;

    return 0;
}