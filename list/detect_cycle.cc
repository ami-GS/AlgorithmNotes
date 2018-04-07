#include <iostream>
#include <set>
using namespace std;


// https://leetcode.com/problems/linked-list-cycle/description/
// use floyd's cycle finding
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

bool has_cycle(ListNode* node) {
    set<ListNode*> S;
    auto nn = node;
    while (nn) {
        if (S.count(nn) != 0) {
            return true;
        }
        S.insert(nn);
        nn = nn->next;
    }
    return false;
}


bool floyd_find(ListNode* node) {
    auto slow = node;
    auto fast = node;
    while(slow->next != nullptr && fast->next->next != nullptr) {
        slow = slow->next;
        fast = fast->next->next;
        if (slow == fast) {
            return true;
        }
    }
    return false;
}

int main() {
    int N = 10;
    ListNode* n = new ListNode(0);
    auto np = n;
    for (int i = 1; i < N; ++i) {
        np->next = new ListNode(i);
        np = np->next;
    }
    n->next->next->next->next = n->next->next;

    cout << has_cycle(n) << endl;
    cout << floyd_find(n) << endl;
    
    return 1;
}
