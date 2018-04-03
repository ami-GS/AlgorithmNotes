#include <stdio.h>

void swap(int* a, int *b) {
    if (a != b) {
        *a ^= *b;
        *b ^= *a;
        *a ^= *b;
    }
}

void swapr(int& a, int& b) {
    if (&a != &b) {
        a ^= b;
        b ^= a;
        a ^= b;
    }
}

int main() {
    int a = 123;
    int b = 432;
    printf("%d, %d\n", a, b);
    swap(&a, &b);
    printf("%d, %d\n", a, b);
    swapr(a, b);
    printf("%d, %d\n", a, b);
}
