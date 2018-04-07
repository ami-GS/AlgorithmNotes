#include <vector>
#include <iostream>
#include <queue>
#include <algorithm>

//problem set is from http://kuuso1.hatenablog.com/entry/2015/12/20/212620
struct edge {
    int toID, cost;
};


class Node {
public:
    int ID;
    std::vector<edge> ad_list;

    void addAdj(int toID, int cost) {
        edge e = {toID, cost};
        this->ad_list.push_back(e);
    }

    void showAdj() {
        for (int i = 0; i < this->ad_list.size(); ++i) {
            std::cout << this->ad_list[i].toID << "," << this->ad_list[i].cost << " ";
        }
        std::cout << std::endl;
    }
};


int main () {
    Node nodes[7];
    nodes[0].addAdj(1, 4);
    nodes[0].addAdj(2, 5);
    nodes[0].addAdj(4, 2);

    nodes[1].addAdj(0, 4);
    nodes[1].addAdj(2, 6);
    nodes[1].addAdj(3, 4);
    nodes[1].addAdj(4, 3);

    nodes[2].addAdj(0, 5);
    nodes[2].addAdj(1, 6);
    nodes[2].addAdj(3, 6);
    nodes[2].addAdj(6, 10);

    nodes[3].addAdj(2, 6);
    nodes[3].addAdj(1, 4);
    nodes[3].addAdj(4, 6);
    nodes[3].addAdj(5, 2);
    nodes[3].addAdj(6, 6);

    nodes[4].addAdj(0, 2);
    nodes[4].addAdj(1, 3);
    nodes[4].addAdj(3, 6);
    nodes[4].addAdj(5, 9);

    nodes[5].addAdj(3, 2);
    nodes[5].addAdj(4, 9);
    nodes[5].addAdj(6, 3);

    nodes[6].addAdj(2, 10);
    nodes[6].addAdj(3, 6);
    nodes[6].addAdj(5, 3);

    for (int i = 0; i < 7; ++i) {
        nodes[i].showAdj();
    }

    int dp[7] = {0, 1<<21, 1<<21, 1<<21, 1<<21, 1<<21, 1<<21};
    std::priority_queue<std::pair<int, int>, std::vector<std::pair<int, int>>, std::greater<std::pair<int, int>>> prQ;
    prQ.push(std::pair<int, int>(0, 0)); // cost, node // cost 0 to node 0 (because start)


    while (!prQ.empty()) {
        auto node  = prQ.top();
        prQ.pop();
        int nID = node.second;
        if (dp[nID] < node.first) {
            continue;
        }
        for(int i = 0; i < nodes[nID].ad_list.size(); ++i){
            edge e = nodes[nID].ad_list[i];
            if(dp[e.toID] > dp[nID] + e.cost){
                dp[e.toID] = dp[nID] + e.cost;
                prQ.push(std::pair<int, int>(dp[e.toID], e.toID));
            }
        }
    }

    for (int i = 0; i < 7; ++i)  {
        std::cout << dp[i] << " ";
    }
    std::cout << std::endl;
}
