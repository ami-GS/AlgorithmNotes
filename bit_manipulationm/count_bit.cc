#include <iostream>

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

int main() {
    int N = 33;
    for (int i = 0; i < N; ++i) {
        std::cout << i << ": " << hamming_weight(i) << " " << hamming_weight_2(i) << " " << __builtin_popcount(i) << std::endl; //2
    }
}
