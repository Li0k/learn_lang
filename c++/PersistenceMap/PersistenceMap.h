//
// Created by liiiyu on 2018/7/20.
//

#ifndef PERSISTENCEMAP_PERSISTENCEMAP_H
#define PERSISTENCEMAP_PERSISTENCEMAP_H

#include <iostream>
#include <unordered_map>
#include <fstream>

using namespace std;

class Entry {
public:
    Entry(string v, int ofs) : value(v), offset(ofs) {}

    ~Entry() = default;

    string value;
    int offset;
    int vLen;
};

class PersistenceMap {
public:
    unordered_map<string, Entry *> kv;
    ifstream in;
    ofstream out;


    string Get(string key) {
        auto it = kv.find(key);
        if (it == kv.end())
            return NULL;

        auto entry = it->second;
        auto offset = entry->offset;

        in.open("/Users/liiiyu/Documents/project/c++/PersistenceMap/kv.txt", ios::binary | ios::in);

        if (!in.good())
            assert(0);


        auto len = entry->vLen;
        char *buf = new char [len];

        in.seekg(entry->offset, ios::beg);
        in.read(buf, len);
        in.close();

        return std::string(buf);
    }

    void Set(string key, const char *value) {
        out.open("/Users/liiiyu/Documents/project/c++/PersistenceMap/kv.txt", ios::binary | ios::out | ios::app);

        if (!out.good()) {
            assert(0);
        }

        out.seekp(0, ios::end);
        auto offset = out.tellp();
        auto e = new Entry(value, offset);
        out.write(value, strlen(value));
        e->vLen = static_cast<int>(strlen(value));

        out.close();

        kv.insert({key, e});
    }
};


#endif //PERSISTENCEMAP_PERSISTENCEMAP_H
