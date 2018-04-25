#include <vector>
#include <iostream>
using namespace std;


int find(vector<int> nums) {
    int tmp = 0;
    for (auto x: nums) {
        tmp ^= x;
    }
    return tmp;
}

int main() {
    vector<vector<int>> nums{{1, 2, 3, 2, 3, 1, 3}, {5, 7, 2, 7, 5, 2, 5}, {2, 3, 5, 4, 5, 2, 4, 3, 5, 2, 4, 4, 2}};
    for (auto num: nums) {
        cout << find(num) << endl;
    }
}
