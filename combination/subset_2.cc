#include <vector>
#include <set>
#include <iostream>
using namespace std;
// https://leetcode.com/problems/subsets-ii/description/

vector<vector<int>> subsetsWithDup(vector<int>& nums) {
    //set<vector<int>> ret(1, vector<int>{});
    set<vector<int>> ret;
    ret.insert(vector<int>{});
    for (int i = 0; i < nums.size(); ++i) {
        vector<vector<int>> tmp(ret.begin(), ret.end());
        for (auto vec: tmp) {
            vec.push_back(nums[i]);
            ret.insert(vec);
        }
    }
    return vector<vector<int>>(ret.begin(), ret.end());
}


// ????? difficult
// int subsetsWithDup_rec(vector<int>& nums, vector<int> buff, vector<vector<int>> *out, int idx) {
//     if (nums.size() == idx) {
//         out->push_back(buff);
//         return 1;
//     }

//     subsets_rec(nums, buff, out, idx+1);
//     buff.push_back(nums[idx]);
//     subsets_rec(nums, buff, out, idx+1);
//     return 1;
// }

int main() {
    vector<int> nums{1,2,2};
    auto out = subsetsWithDup(nums);
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
