#include <vector>
#include <iostream>
using namespace std;

//https://leetcode.com/problems/same-tree/description/
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};


bool isSameTree(TreeNode* p, TreeNode* q) {
    if (!p && !q) {
        return true;
    } else if (!p && q || p && !q) {
        return false;
    }
    bool out = isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
    if (p->val != q->val) {
        return false;
    }
    return out;
}


int main() {

    TreeNode *root1, *root2;
    root1 = new TreeNode(3);
    root1->right = new TreeNode(20);
    root1->left = new TreeNode(9);
    root1->right->right = new TreeNode(7);
    root1->right->left = new TreeNode(15);

    root2 = new TreeNode(3);
    root2->right = new TreeNode(20);
    root2->left = new TreeNode(9);
    root2->right->right = new TreeNode(7);
    root2->right->left = new TreeNode(15);

    cout << isSameTree(root1, root2) << endl;
}

