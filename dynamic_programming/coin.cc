#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;

//https://leetcode.com/problems/coin-change/description/
int coinChange(vector<int>& coins, int amount) {
    vector<int> dp(amount+1, 1<<21);
    dp[0] = 0;
    for (int i = 1; i < dp.size(); ++i) {
        for (auto coin: coins) {
            if (i >= coin) {
                dp[i] = min(dp[i], dp[i-coin]+1);
            }
        }
    }
    if (dp[amount] == 1<<21) {
        return -1;
    }
    return dp[amount];
}

int main() {
    vector<vector<int>> coinsets{vector<int>{1, 2, 5}, vector<int>{2}};
    vector<int> amounts{11, 3};
    for (int i = 0; i < amounts.size(); ++i) {
        cout << coinChange(coinsets[i], amounts[i]) << endl;
    }
}
