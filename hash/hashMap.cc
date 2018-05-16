#include <string>
#include <vector>
#include <utility>
#include <iostream>
using namespace std;


template <typename T, typename U>
class HashMap {
    int tableSize;
    vector<pair<T, U>> table;
    int calcHash(T str) {
        int len = str.size();
        if (len >= 2) {
            len *= str[0] * str[len-1];
        }
        return len % tableSize;
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
    HashMap<string, int> myMap(256);
    myMap.add({"one", 1});
    myMap.add({"3ke", 2});
    myMap.add({"098765", 123});


    cerr << myMap.get("one") << endl;
    myMap.add({"one", 2});
    cerr << myMap.get("one") << endl;
    cerr << myMap.get("3ke") << endl;
    cerr << myMap.get("098765") << endl;
    return 1;
}
