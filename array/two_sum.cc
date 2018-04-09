#include <unordered_set>
#include <vector>
#include <utility>
#include <iostream>
using namespace std;

vector<pair<int, int>> calc(const vector<int> &nums, int target) {
    vector<pair<int, int>> ans;
    unordered_set<int> st;
    for (int i = 0; i < nums.size(); ++i) {
        int tmp = target-nums[i];
        if (st.count(tmp) == 0) {
            st.insert(tmp);
        }
        if (st.count(nums[i]) == 1) {
            ans.push_back(pair<int, int>(tmp, nums[i]));
        }
    }
    return ans;
}

int main() {
    vector<int> nums{2, 3,4, 5,6, 7, 9, 11, 15};
    auto list = calc(nums, 9);

    std::cout << "[ ";
    for(auto v: list) {
        std::cout << "[";
        std::cout << v.first << ", " << v.second;
        std::cout << "], ";
    }
    std::cout << "]" << std::endl;
}
