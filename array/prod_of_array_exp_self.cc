#include <vector>
#include <stdio.h>

std::vector<int> calc(std::vector<int> nums) {
    // time:O(N) space:O(1) is accepted
    int N = nums.size();
    std::vector<int> out(N, 1);
    int left = 1;
    int right = 1;
    for (int i = 0; i < N; ++i) {
        out[i] *= left;
        out[N-i-1] *= right;
        left *= nums[i];
        right *= nums[N-i-1];
    }
    return out;
}
int main() {
    std::vector<int> nums{1,2,3,4,5,6};
    auto out = calc(nums);
    for (int i = 0 ; i < out.size(); ++i) {
        printf("%d ", out[i]);
    }
}
