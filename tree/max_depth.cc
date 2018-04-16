#include <vector>
#include <iostream>
using namespace std;

//https://leetcode.com/problems/maximum-depth-of-binary-tree/
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

int maxDepth(TreeNode* root) {
    if (root->right) {
         return maxDepth(root->right) + 1;
    }
    if (root->left) {
        return maxDepth(root->left) + 1;
    }
    return 1;
}

int main() {

    TreeNode* root;
    root = new TreeNode(3);
    root->right = new TreeNode(20);
    root->left = new TreeNode(9);
    root->right->right = new TreeNode(7);
    root->right->left = new TreeNode(15);

    cout << maxDepth(root) << endl;
}
