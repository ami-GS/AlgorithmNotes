#include <vector>

int calc(std::vector<int> nums) {
    int start = 0;
    int end = nums.size()-1;
    while (start < end) {
        if (nums[start] < nums[end]) {
            return nums[start];
        }

        int pivot = (start + end)/2;
        if (nums[pivot] < nums[end]) {
            end = pivot;
        } else {
            start = pivot+1;
        }
    }
    return nums[start];
}

int main() {
    //std::vector<int> nums{4, 5, 6, 7, 0, 1, 2};
    //std::vector<int> nums{0, 1, 2,4,5,6,7};
    std::vector<int> nums{1,2,4,5,6,7,0};
    auto out = calc(nums);
    printf("%d\n", out);
}
