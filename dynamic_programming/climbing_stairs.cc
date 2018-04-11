#include <vector>
#include <iostream>
using namespace std;
// https://leetcode.com/problems/climbing-stairs/description/

int climbStairs(int n) {
    vector<int> dp(n+1, 0);
    dp[0] = 0;
    dp[1] = 1;
    dp[2] = 2;
    for (int i = 3; i < dp.size(); ++i) {
        dp[i] = dp[i-1] + dp[i-2];
    }
    return dp[n];
}

int climbStairs_div(int n) {
    if(n == 0) {
        return 0;
    } else if (n == 1) {
        return 1;
    } else if (n == 2) {
        return 2;
    }
    int a = climbStairs_div(n-1);
    a += climbStairs_div(n-2);
    return a;
}

int main() {
    vector<int> nums{1,2,3,4,5,6,7,8,9,10,11,12,13,14};

    for (auto v: nums) {
        cout << v << ": "<< climbStairs(v) << " ";
        cout << climbStairs_div(v) << endl;
    }
}
