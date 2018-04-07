#include <vector>


int calc(const std::vector<int> &nums) {
    int N = nums.size()+1;
    int ans = (N-1)*N/2;
    if (N%2 == 1) {
        ans += N/2;
    }
    int acc = 0;
    for (int i = 0; i < N-1; ++i) {
        acc += nums[i];
    }
    return ans - acc;
}

// need review
int calc2(const std::vector<int> &nums) {
    int tmp = 0;
    for (int i = 0 ; i < nums.size(); i++) {
        tmp ^= nums[i] ^ (i+1);
        printf("%d ", tmp);
    }
    return tmp;
}

int main() {
    std::vector<int> nums{9,6,4,2,3,5,7,0,1};
    printf("%d\n",calc(nums));
    printf("%d\n", calc2(nums));
}
