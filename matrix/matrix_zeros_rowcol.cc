#include <vector>
#include <iostream>
using namespace std;

void setZeroes(vector<vector<int>>& matrix) {
    int R = matrix.size();
    int C = matrix[0].size();
    int col0loc = 0;
    for (int i = 0; i < R; ++i) {
        for (int j = col0loc; j < C; ++j) {
            if (matrix[i][j] == 0) {
                col0loc = j+1;
                for (int jj = 0; jj < C; ++jj) {
                    matrix[i][jj] = 0;
                }
                for (int ii = 0; ii < R; ++ii) {
                    matrix[ii][j] = 0;
                }
                break;
            }
        }
    }
}



int main() {
    int N = 4;
    vector<vector<int>> mat(4);
    for (int i = 0; i < N; ++i) {
        mat[i] = vector<int>(N, 1);
    }
    mat[2][2] = 0;
    mat[0][0] = 0;
    mat[3][3] = 0;
    setZeroes(mat);

    int R = mat.size();
    int C = mat[0].size();
    for (int i = 0; i < R; ++i) {
        for (int j = 0; j < C; ++j) {
            cout << mat[i][j] << " ";
        }
        cout << endl;
    }

}
