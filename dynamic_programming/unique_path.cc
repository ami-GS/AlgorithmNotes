#include <vector>
#include <iostream>
using namespace std;
//https://leetcode.com/problems/unique-paths/description/

int uniquePaths(int m, int n) {
    vector<vector<int>> dp(m, vector<int>(n, 0));
    for (int i = 0; i < m; ++i)  {
        dp[i][0] = 1;
    }
    for (int i = 0; i < n; ++i)  {
        dp[0][i] = 1;
    }

    for (int i = 1; i < m; ++i) {
        for (int j = 1; j < n; ++j) {
            dp[i][j] = dp[i-1][j] + dp[i][j-1];
        }
    }

    return dp[m-1][n-1];
}

int main(int argc, char **argv) {
    cout <<uniquePaths(3, 2) << endl;
    cout <<uniquePaths(7, 3) << endl;
}
