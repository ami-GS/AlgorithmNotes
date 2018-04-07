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

int hamming_weight(uint32_t val) {
    int count = 0;
    while (val != 0) {
        count++;
        val &= val-1;
    }
    return count;
}

int main() {
    std::cout << hamming_weight(11) << std::endl; //2
}
