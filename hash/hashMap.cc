#include <string>
#include <vector>
#include <utility>
#include <iostream>
using namespace std;


template <typename T, typename U>
class HashMap {
    int tableSize;
    vector<pair<T, U>> table;

    int hashing(string key) {
        int hashBase;
        int len = key.size();
        if (len >= 2) {
            hashBase = len * key[0] * key[len-1];
        } else if (len == 1){
            hashBase = len * key[0];
        } else {
            //error?
        }
        return hashBase;
    }
    int hashing(int key) {
        return key;
    }

    int calcHash(T key) {
        int hashBase = 1;
        return hashing(key);
    }
public:
    HashMap(int tableSize) : tableSize(tableSize), table(vector<pair<T, U>>(tableSize, {})) {}
    void add(const pair<T, U> data) {
        int hash = this->calcHash(data.first);
        if (this->table[hash] == pair<T, U>{} ) {
            this->table[hash] = data;
        } else {
            if (this->table[hash].first == data.first) {
                goto INSERT;
            }
            while (this->table[++hash] != pair<T, U>{}) {}
        INSERT:
            this->table[hash] = data;
        }
    }

    U get (const T &key) {
        int hash = this->calcHash(key);
        return this->table[hash].second;
    }
};

int main() {
    HashMap<string, int> myStrMap(256);
    HashMap<int, int> myIntMap(256);
    myStrMap.add({"one", 1});
    myStrMap.add({"3ke", 2});
    myStrMap.add({"098765", 123});
    cerr << myStrMap.get("one") << endl;
    myStrMap.add({"one", 2});
    cerr << myStrMap.get("one") << endl;
    cerr << myStrMap.get("3ke") << endl;
    cerr << myStrMap.get("098765") << endl;

    myIntMap.add({1, 1});
    myIntMap.add({2, 2});
    myIntMap.add({3, 123});
    cerr << myIntMap.get(1) << endl;
    myIntMap.add({1, 2});
    cerr << myIntMap.get(1) << endl;
    cerr << myIntMap.get(2) << endl;
    cerr << myIntMap.get(3) << endl;


    return 1;
}
