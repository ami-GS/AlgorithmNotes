#include <vector>
#include <algorithm>
#include <stdio.h>

//std::vector calc(std::vector<int> nums) {
int calc(std::vector<int> nums) {
    int N = nums.size();

    int dp = 0; // will reset if dp+nums[i] is smaller than nums[i]
    int ans = 0;
    int curMax = 0;
    for (int i = 0; i < N; i++) {
        curMax = std::max(nums[i], curMax+nums[i]);
        ans = std::max(ans, curMax);
        printf("%d : %d %d\n", i, curMax, ans);
    }

    return ans;
}

int main() {
    std::vector<int> nums{-2,1,-3,4,-1,2,1,-5,4};
    auto out = calc(nums);
    printf("%d\n", out);
}
