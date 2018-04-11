#include <vector>
#include <algorithm>
#include <stdio.h>
using namespace std;
//https://leetcode.com/problems/combination-sum-iv/description/

// not finished
int combinationSum(vector<int>& nums, int target) {
    vector<vector<int>> dp(nums.size()+1, vector<int>(target+1, 0));
    for (int i = 0; i < dp.size(); ++i) {
        dp[i][0] = 1;
    }

    for (int j = 1; j < target+1; ++j) {
        for (int i = 1; i < nums.size()+1; ++i) {
            if (j >= nums[i-1]) {
                dp[i][j] += dp[i][j-nums[i-1]];
            }

        }
    }
    return dp[nums.size()][target];
}

int combinationSum_r(vector<int>& nums, int target) {
    if(target < 0) {
        return 0;
    } else if (target == 0) {
        return 1;
    }

    int ans = 0;
    for (auto val: nums) {
        ans += combinationSum_r(nums, target-val);
    }
    return ans;
}

int main() {
    vector<int> nums{1, 2, 3};
    printf("%d\n", combinationSum(nums, 4));
}
