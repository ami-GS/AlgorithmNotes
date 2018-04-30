#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;

//https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

int kthSmallest(TreeNode* root, int &k) {
    if (root == nullptr) {
        return -1;
    }
    int left = kthSmallest(root->left, k);
    k--;
    if (k == 0) {
        return root->val;
    }
    int right = kthSmallest(root->right, k);
    return max(left, right);
}

void makeBST(TreeNode *root, vector<TreeNode> *nodes) {
    for (int i = 0; i < nodes->size(); ++i) {
        auto p = root;
        auto node = &(*nodes)[i];
        while (1) {
            if (p->val <= node->val) {
                if (p->right == nullptr) {
                    p->right = node;
                    break;
                } else {
                    p = p->right;
                }
            } else {
                if (p->left == nullptr) {
                    p->left = node;
                    break;
                } else {
                    p = p->left;
                }
            }
        }
    }
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
    TreeNode root(1);
    vector<TreeNode> nodes{TreeNode(8), TreeNode(9), TreeNode(5), TreeNode(6), TreeNode(20), TreeNode(3)};
    makeBST(&root, &nodes);
    //printTree(&root);
    for (int kth = 1; kth < nodes.size()+2; ++kth) {
        int tmp = kth;
        cout << kthSmallest(&root, tmp) << endl;
    }
}
