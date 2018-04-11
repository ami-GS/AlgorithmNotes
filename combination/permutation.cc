#include <vector>
#include <iostream>
using namespace std;


vector<vector<int>> permute_dup_loop(vector<int> nums, int len) {
    vector<vector<int>> results{{}};
    int curLen = 0;
    while (curLen < len) {
        vector<vector<int>> tmp;
        for (int i = 0; i < results.size(); ++i) {
            for (int j = 0; j < nums.size(); ++j) {
                auto tmptmp = results[i];
                tmptmp.push_back(nums[j]);
                tmp.push_back(tmptmp);
            }
        }
        curLen++;
        results = tmp;
    }
    return results;
}

vector<vector<int>> permute_loop(vector<int> nums, int len) {
    vector<vector<int>> results{{}};
    int curLen = 0;
    while (curLen < len) {
        vector<vector<int>> tmp;
        for (int i = 0; i < results.size(); ++i) {
            for (int j = 0; j < nums.size(); ++j) {
                auto tmptmp = results[i];
                tmptmp.push_back(nums[j]);
                tmp.push_back(tmptmp);
            }
        }
        curLen++;
        results = tmp;
    }
    return results;
}




int main() {
    vector<int> nums{1,2,3,4,5};
    auto out = permute_dup_loop(nums, 3);
    cout << "[ ";
    for(auto v: out) {
        cout << "[";
        for (auto elm: v) {
            cout << elm << ",";
        }
        cout << "],";
    }
    cout << " ]" << endl;
}
