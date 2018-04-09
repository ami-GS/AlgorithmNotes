#include <algorithm>
#include <vector>
#include <unordered_map>
#include <iostream>
// clap
void calc(std::vector<int> &nums, std::vector<std::vector<int>> *ans, std::vector<int> buf, int idx, int sum, int depth) {
    if (depth == 3 && sum == 0) {
        ans->push_back(buf);
    }

    for (int i = idx; i < nums.size(); ++i) {
        buf[depth] = nums[i];
        calc(nums, ans, buf, i, sum+nums[i], depth+1);
    }
    calc(nums, ans, buf, idx+1, sum, depth);
}

void calc(std::vector<int> &nums) {
    std::unordered_map<int, int> hash;
    std::sort(nums.begin(), nums.end());

    for (int i = 0; i < nums.size(); ++i) {
        hash[nums[i]] = 1;
    }

    for (int i = 0; i < nums.size(); ++i) {
        for (int j = 1; j < nums.size(); ++j) {
            hash[0 - nums[i]+nums[j]];
        }
    }
}

// after understand
std::vector<std::vector<int>> calc2(std::vector<int> nums, int target) {
    std::sort(nums.begin(), nums.end());
    std::vector<std::vector<int>> ans;
    for (int i = 0; i < nums.size(); ++i) {
        for (int j = i+1, k = nums.size(); j < k;) {
            int sum = nums[i] + nums[j] + nums[k];
            if (sum == target) {
                ans.push_back({nums[i], nums[j], nums[k]});
                k--;
            } else if (sum > target) {
                k--;
            } else {
                j++;
            }
        }
    }
    return ans;
}

int main() {
    std::vector<int> nums{-1, 0, 1, 2, -1, -4};
    std::vector<std::vector<int>> ans;
    std::vector<int> buf(3, 0);
    //calc(nums, &ans, buf, 0, 0, 0);

    auto out = calc2(nums, -1);
    std::cout << "[ ";
    for(auto v: out) {
        std::cout << "[";
        for (auto elm: v) {
            std::cout << elm << ",";
        }
        std::cout << "],";
    }
    std::cout << " ]" << std::endl;
}
