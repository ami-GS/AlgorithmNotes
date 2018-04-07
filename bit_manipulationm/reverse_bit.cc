#include <stdio.h>
#include <vector>

// Be careful to do indexing (i - 1)
uint32_t calc(uint32_t num) {
    uint32_t ans = 0;
    for (int i = 0; i < 32; ++i) {
        ans |= ((num >> i) & 1) << (32-i-1);
    }

    return ans;
}

int main() {
    printf("%d\n", calc(43261596));
}
