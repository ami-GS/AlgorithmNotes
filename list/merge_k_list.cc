#include <iostream>
#include <vector>
using namespace std;
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(nullptr) {}
};

ListNode* mergeKLists(vector<ListNode*>& lists) {
    ListNode dummy(0);
    ListNode* out = &dummy;
    dummy.next = nullptr;
    while (1) {
        int count = 0;
        int idx;
        int val = 1<<21;
        for (int i = 0; i < lists.size(); ++i) {
            if (lists[i] != nullptr) {
                if (lists[i]->val < val) {
                    val = lists[i]->val;
                    idx = i;
                }
                ++count;
            }
        }
        if (count == 0) {
            break;
        }
        out->next = lists[idx];
        lists[idx] = lists[idx]->next;
        out = out->next;
    }
    return dummy.next;
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

    ListNode* n3 = new ListNode(0);
    np = n3;
    for (int i = 1; i < N; ++i) {
        np->next = new ListNode(5*i);
        np = np->next;
    }

    vector<ListNode*> lists{n, n2, n3};
    auto out = mergeKLists(lists);

    for (int i = 0; i < lists.size()*N; ++i) {
        cout << out->val << "->";
        out = out->next;
    }

}
