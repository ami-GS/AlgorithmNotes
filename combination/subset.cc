#include <vector>
#include <iostream>
using namespace std;

// https://leetcode.com/problems/subsets/description/
vector<vector<int>> subsets(vector<int>& nums) {
    vector<vector<int>> out(1, vector<int>(0));
    for (int i = 0; i < nums.size(); ++i) {
        vector<int> tmp;
        int curSize = out.size();
        for (int j = 0; j < curSize; ++j) {
            tmp = out[j];
            tmp.push_back(nums[i]);
            out.push_back(tmp);
        }
    }
    return out;
}

vector<vector<int>> subsets_bit(vector<int>& nums) {
    vector<vector<int>> out;
    for (int i = 0; i < (1<<nums.size()); ++i) {
        vector<int> tmp;
        for (int j=0, mask = 1; j < nums.size(); mask <<= 1, ++j) {
            if (i & mask) {
                tmp.push_back(nums[j]);
            }
        }
        out.push_back(tmp);
    }
    return out;
}


int subsets_rec(vector<int>& nums, vector<int> buff, vector<vector<int>> *out, int idx) {
    if (nums.size() == idx) {
        out->push_back(buff);
        return 1;
    }

    subsets_rec(nums, buff, out, idx+1);
    buff.push_back(nums[idx]);
    subsets_rec(nums, buff, out, idx+1);
    return 1;
}


int main() {
    vector<int> nums{1,2,3,4,5};
    auto out = subsets(nums);
    cout << "[ ";
    for(auto v: out) {
        cout << "[";
        for (auto elm: v) {
            cout << elm << ",";
        }
        cout << "],";
    }
    cout << " ]" << endl;

    out = subsets_bit(nums);
    cout << "[ ";
    for(auto v: out) {
        cout << "[";
        for (auto elm: v) {
            cout << elm << ",";
        }
        cout << "],";
    }
    cout << " ]" << endl;

    vector<int> buff;
    vector<vector<int>> out2;
    subsets_rec(nums, buff, &out2, 0);
    cout << "[ ";
    for(auto v: out2) {
        cout << "[";
        for (auto elm: v) {
            cout << elm << ",";
        }
        cout << "],";
    }
    cout << " ]" << endl;
}
