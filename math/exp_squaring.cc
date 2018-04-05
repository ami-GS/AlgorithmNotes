#include <iostream>

const long long M = 10000000000;

long long run(int n, int p) {
    if (p == 0) {
        return 1;
    } else if (p % 2 == 1) {
        return (run(n, p-1) * n) % M;
    } else {
        uint64_t tmp = run(n, p/2) % M;
        return (tmp * tmp) % M;
    }
}

int main() {
    long long ans = 0;
    for (int i = 1; i < 1001; i++) {
        ans += run(i, i) % M;
    }
    std::cout << ans << std::endl;
}
