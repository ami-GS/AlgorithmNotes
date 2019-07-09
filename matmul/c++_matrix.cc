
#define M 2048
#define N 512
#define K 1024

void calc() {
    int A[M][N];
    int B[N][K];
    int C[M][K];

    for (int m = 0; m < M; ++m) {
        for (int k = 0; k < K; ++k) {
            for (int n = 0; n < N; ++n) {
                C[m][n] += A[m][n] * B[n][k];
            }
        }
    }
}

int main() {
    calc();
    return 0;
}
