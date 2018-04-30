#include <iostream>
#include <vector>

int hamming_weight_naive(uint32_t val) {
    int out = 0;
    while (val != 0) {
        if (val & 1)
            out++;
        val >>= 1;
    }
    return out;
}

// invert lowest bit '1'
int hamming_weight(uint32_t val) {
    int count = 0;
    while (val != 0) {
        count++;
        val &= val-1;
    }
    return count;
}
// O(logN)
int hamming_weight_2(uint32_t val) {
    int ret = 0;
    while (val) {
        val = val & (val-1);
        ++ret;
    }
    return ret;
}

std::vector<int> table;

int hamming_weight_table(uint32_t val) {
    if (table.size() != 255) {
        for (int i = 0; i < 255; i++) {
            int tmp = i;
            table.push_back(0);
            while (tmp) {
                tmp = tmp & (tmp - 1);
                ++table[i];
            }
        }
    }
    int out = 0;
    while (val) {
        out += table[val & 0xff];
        val >>= 8;
    }
    return out;
}



int main() {
    int N = 33;
    for (int i = 0; i < N; ++i) {
        std::cout << i << ": " << hamming_weight(i) << " " << hamming_weight_2(i) << " " << __builtin_popcount(i) << " " << hamming_weight_table(i) << std::endl; //2
    }
}
