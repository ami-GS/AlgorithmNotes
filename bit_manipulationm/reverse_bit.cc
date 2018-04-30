#include <stdio.h>
#include <vector>
using namespace std;

// Be careful to do indexing (i - 1)
uint32_t calc(uint32_t num) {
    uint32_t ans = 0;
    for (int i = 0; i < 32; ++i) {
        ans |= ((num >> i) & 1) << (32-i-1);
    }

    return ans;
}

vector<uint8_t> table;

uint32_t calc_table(uint32_t num) {
    if (table.size() != 255) {
        for (uint8_t i = 0; i < 255; ++i) {
            uint8_t tmp = i;
            table.push_back(0);
            for (int j = 0; tmp; ++j) {
                table[i] |= (tmp & 1) << (7-j);
                tmp >>= 1;
            }
        }
    }
    uint32_t out = 0;
    int i = 0;
    while (num) {
        out |= table[num & 0xff] << (sizeof(out)-1-(i++)) * 8;
        num >>= 8;
    }
    return out;
}


int main() {
    for (uint32_t i = 1; i < 123456789; i *= 2) {
        printf("%u %u\n", calc(i), calc_table(i));
    }
}
