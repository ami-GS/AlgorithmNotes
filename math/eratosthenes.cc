#include <stdio.h>


#define N 30 + 1


void fill_sieve(bool *sieve) {
    for (int i = 2; i < N; ++i) {
        sieve[i] = true;
        if (i%2 == 0) {
            sieve[i] = false;
        }
    }

    sieve[0] = false;
    sieve[1] = false;
    sieve[1] = true;
    //sieve[0] = false;
    for (int i = 2; i*i < N; i += 2) {
        for (int j = i*2; j < N; j += j) {
            sieve[j] = false;
        }
    }

}

int main() {
    bool sieve[N];
    fill_sieve(sieve);
    for (int i = 1; i < N; ++i) {
        if (sieve[i]) {
            printf("%d ", i);
        }
    }
}
