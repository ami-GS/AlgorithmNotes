#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;

//https://leetcode.com/problems/binary-tree-maximum-path-sum/description/
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};


int helper(TreeNode* root, int *sum) {
    if (root == nullptr) {
        return 0;
    }

    int left = max(0, helper(root->left, sum));
    int right = max(0, helper(root->right, sum));
    *sum = max(*sum, left + right + root->val);

    return max(left, right) + root->val;
}

int maxPathSum(TreeNode* root) {
    int sum = 0;
    int out = helper(root, &sum);
    cout << sum << endl;
}

int main() {
    TreeNode *root;

    root = new TreeNode(1);
    root->left = new TreeNode(9);
    root->left->right = new TreeNode(5);
    root->left->left = new TreeNode(5);
    root->right = new TreeNode(20);
    //root->right->left = new TreeNode(15);
    root->right->right = new TreeNode(7);

    maxPathSum(root);
}
