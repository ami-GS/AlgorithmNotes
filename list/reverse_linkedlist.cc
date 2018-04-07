#include <iostream>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};


ListNode* reverse(ListNode *node) {
    ListNode* prev = nullptr;
    ListNode* nxt = nullptr;
    while (node) {
        nxt = node->next;
        node->next = prev;
        prev = node;
        node = nxt;
    }
    return prev;
}

int main() {
    int N = 10;
    ListNode* n = new ListNode(0);
    auto np = n;
    for (int i = 1; i < N; ++i) {
        np->next = new ListNode(i);
        np = np->next;
    }

    np = n;
    for (int i = 0; i < N; ++i) {
        cout << np->val << "->";
        np = np->next;
    }
    cout << endl;
    auto nn = reverse(n);

    for (int i = 0; i < N; ++i) {
        cout << nn->val << "->";
        nn = nn->next;
    }
    cout << endl;

    return 1;
}
