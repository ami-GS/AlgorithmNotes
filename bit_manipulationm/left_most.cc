#include <iostream>
#include <cmath>
using namespace std;

int calc(int x) {
    return log2(x);
}

int main() {
    for (int i = 1; i < 36; ++i) {
        cout << i << ":" << calc(i) << ", ";
    }
    cout << endl;
}
