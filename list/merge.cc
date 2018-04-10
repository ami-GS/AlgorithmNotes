#include <iostream>
using namespace std;
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(nullptr) {}
};

//https://leetcode.com/problems/merge-two-sorted-lists/description/
ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
    ListNode dummy(0);
    ListNode* out = &dummy;
    dummy.next = nullptr;
    while(1) {
        if(!l1) {
            out->next = l2;
            break;
        } else if (!l2) {
            out->next = l1;
            break;
        }
        if (l1->val < l2->val) {
            out->next = l1;
            l1 = l1->next;
        } else {
            out->next = l2;
            l2 = l2->next;
        }
        out = out->next;
    }
    return dummy.next;
}

ListNode* mergeTwoLists_r(ListNode* l1, ListNode* l2) {
    if (!l1) {
        return l2;
    } else if (!l2) {
        return l1;
    }

    if (l1->val < l2->val) {
        l1->next = mergeTwoLists(l1->next, l2);
        return l1;
    } else {
        l2->next = mergeTwoLists(l1, l2->next);
        return l2;
    }
}

int main() {
    int N = 5;
    ListNode* n = new ListNode(0);
    auto np = n;
    for (int i = 1; i < N; ++i) {
        np->next = new ListNode(2*i);
        np = np->next;
    }

    ListNode* n2 = new ListNode(0);
    np = n2;
    for (int i = 1; i < N; ++i) {
        np->next = new ListNode(3*i);
        np = np->next;
    }

    auto out = mergeTwoLists(n, n2);
    //auto out = mergeTwoLists_r(n, n2);

    for (int i = 0; i < 2*N; ++i) {
        cout << out->val << "->";
        out = out->next;
    }
}
