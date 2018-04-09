#include <vector>
#include <algorithm>
#include <stdio.h>

void swap(int& a, int& b) {
    a ^= b;
    b ^= a;
    a ^= b;
}

int calc(std::vector<int> nums) {
    int N = nums.size();
    std::vector<int> out(N, -1);

    int curMax = nums[0];
    int curMin = nums[0];
    int ans = 0;
    for (int i = 1; i < N; i++) {
        if (nums[i] < 0) {
            swap(curMax, curMin);
        }
        curMax = std::max(nums[i], curMax*nums[i]);
        curMin = std::min(nums[i], curMin*nums[i]);
        ans = std::max(curMax, ans);
        printf("%d : %d %d\n", i, curMax, curMin);
    }

    return ans;
}

int main() {
    std::vector<int> nums{-2,1,-3,4,-1,2,1,-5,4};
    auto out = calc(nums);
    printf("%d\n", out);
}
