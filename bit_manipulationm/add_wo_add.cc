#include <iostream>

// TODO: review needed
int add(int a, int b) {
    int sum;

    while (b != 0) {
        sum = b ^ a;
        b = (a & b) << 1;
        a = sum;
    }
    return sum;
}


int main() {
    int A = 5, B = 20;
    for (int i = 0; i < A; ++i) {
        for (int j = 0; j < B; j += 2) {
            std::cout << i << " + " << j << " = " << add(i, j) << "; ";
        }
        std::cout << std::endl;
    }
}
