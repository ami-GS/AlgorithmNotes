#include <stack>
#include <vector>
#include <iostream>
#include <stdio.h>
using namespace std;

class Node;


class Node {
public:
    int ID;
    vector<Node*> adj_nodes;
    Node(int ID) : ID(ID), adj_nodes({}) {}
    void addNode(Node *child) {
        this->adj_nodes.push_back(child);
    }
};


bool has_cycle(Node* root) {
    vector<bool> visited(5, false);
    vector<bool> walking(5, false);
    stack<Node*> stk;
    stk.push(root);

    while (!stk.empty()) {
        auto nd = stk.top();
        visited[nd->ID] = true;
        walking[nd->ID] = true;

        bool visitedSomewhere = false;
        for (auto adj : nd->adj_nodes) {
            if (!visited[adj->ID]) {
                stk.push(adj);
                visitedSomewhere = true;
            } else if (walking[adj->ID]) {
                cout << "detect" << endl;
            }
        }

        if (!visitedSomewhere) {
            stk.pop();
            walking[nd->ID] = false;
        }
    }
    return false;
}



int main() {
    auto root = Node(0);
    auto n1 = Node(1);
    auto n2 = Node(2);
    auto n3 = Node(3);
    auto n4 = Node(4);
    auto n5 = Node(5);

    root.addNode(&n1);
    root.addNode(&n2);
    n2.addNode(&n3);
    n3.addNode(&n5);
    n5.addNode(&n4);
    n4.addNode(&n2);

    has_cycle(&root);
    return 1;
}
