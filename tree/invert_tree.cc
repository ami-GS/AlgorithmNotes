#include <vector>
#include <iostream>
using namespace std;

//https://leetcode.com/problems/invert-binary-tree/description/
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

TreeNode* invertTree(TreeNode* root) {
    if (root == nullptr) {
        return nullptr;
    }
    invertTree(root->left);
    invertTree(root->right);

    TreeNode* tmp = root->left;
    root->left = root->right;
    root->right = tmp;

    return root;
}

void printTree(TreeNode* root) {
    if (root == nullptr) {
        return;
    }
    printTree(root->left);
    printTree(root->right);
    cout << root->val << endl;
}

int main() {
    TreeNode *root;
    root = new TreeNode(4);
    root->right = new TreeNode(7);
    root->left = new TreeNode(2);
    root->right->right = new TreeNode(9);
    root->right->left = new TreeNode(6);
    root->left->right = new TreeNode(3);
    root->left->left = new TreeNode(1);

    printTree(root);
    invertTree(root);
    cout << endl;
    printTree(root);
}
