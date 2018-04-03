#include <string>
#include <vector>
#include <iostream>
#include <algorithm>


int LCS_length(std::string str_a, std::string str_b) {
    int **dp = new int*[str_a.size()+1];
    for (int i = 0; i < str_a.size()+1; ++i ) {
        dp[i] = new int[str_b.size()+1];
        for (int j = 0; j < str_b.size()+1; ++j ) {
            dp[i][j] = 0;
        }
    }

    for (int i = 1; i < str_a.size(); ++i ) {
        for (int j = 1; j < str_b.size(); ++j ) {
            dp[i][j] = std::max(dp[i][j-1], std::max(dp[i-1][j], dp[i-1][j-1]));
            if (str_a[i-1] == str_b[j-1])
                dp[i][j]++;
            std::cout << dp[i][j] << " ";
        }
        std::cout << std::endl;
    }


    for (int i = 0; i < str_b.size(); ++i) {
        delete[] dp[i];
        dp[i] = 0;
    }
    delete[] dp;
    dp = 0;
    return 1;
}


int main() {
    LCS_length("ABCDGH", "AEDFHR");
    LCS_length("AGGTAB", "GXTXAYB");
    return 1;
}
