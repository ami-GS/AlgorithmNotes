#include <vector>
#include <unordered_map>
#include <queue>
#include <iostream>
using namespace std;

//https://leetcode.com/problems/top-k-frequent-elements/description/
vector<int> topKFrequent(vector<int>& nums, int k) {
    unordered_map<int, int> table;
    //priority_queue<pair<int, int>, vector<pair<int, int>>, less<pair<int, int>>> heap;
    priority_queue<pair<int, int>, vector<pair<int, int>>> heap;
    for (auto v: nums) {
        table[v]++;
    }
    for (auto in: table) {
        heap.push({in.second, in.first});
    }

    vector<int> out(k, 0);
    for (int i = 0; i < k; ++i) {
        out[i] = heap.top().second;
        heap.pop();
    }
    return out;
}

int main() {
    vector<int> nums{1,1,1,2,2,3, 4,4,4,4,4};
    vector<int> out = topKFrequent(nums, 3);

    cout << "[";
    for (auto v: out) {
        cout << v << ", ";
    }
    cout << "]" << endl;
}
