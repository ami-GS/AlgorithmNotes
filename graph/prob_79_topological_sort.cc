#include <iostream>
#include <fstream>
#include <string>
#include <vector>


struct node {
    int num_connect;
    std::vector<char> connect_list;
    char ID;
};

const int NODE_NUM = 10;

int hasI(std::vector<char> v, char I) {
    for (int i = 0; i < v.size(); i++) {
        if (v[i] == I) {
            return 1;
        }
    }
    return 0;
}

int set_node(node* all_nodes, char* line) {
    // i can stand for connection from left to right
    // abc: b has one connection from a, c has two connections from a and b
    for (int i = 0; i < 3; ++i) {
        char II = line[i]-'0';
        for (int j = i+1; j < 3; j++) {
            // II -> JJ
            char JJ = line[j]-'0';
            if ( !hasI(all_nodes[II].connect_list, JJ) ) {
                all_nodes[II].connect_list.push_back(JJ);
                all_nodes[JJ].num_connect++;
            }
        }
    }


    return 1;
}

int tsort(node* all_nodes) {
    int stack[10];
    int sIdx = 0;
    int fin[10];
    for (int i = 0; i < NODE_NUM; ++i) {
        fin[i]=i;
    }
    fin[4] = -1;
    fin[5] = -1;

    for (int i = 0; i < NODE_NUM; ++i) {
        if (fin[i] != -1&& all_nodes[i].num_connect == 0) {
            stack[sIdx++] = i;
        }
    }
    while (sIdx) {
        node* dat = &all_nodes[stack[--sIdx]];
        fin[dat->ID] = -1;
        for (int i = 0; i < dat->connect_list.size(); ++i) {
            all_nodes[dat->connect_list[i]].num_connect--;
        }

        for (int i = 0; i < NODE_NUM; ++i) {
            if (fin[i] != -1 && all_nodes[i].num_connect == 0) {
                stack[sIdx++] = i;
            }
        }
        std::cout << (int)dat->ID;
    }
    std::cout << std::endl;
    return 1;
}

int main() {
    std::ifstream ifs("./p079_keylog.txt");
    char str[256];
    if (ifs.fail())
    {
        std::cerr << "fail to read file" << std::endl;
        return -1;
    }

    node all_nodes[NODE_NUM];
    for (int i = 0; i < NODE_NUM; ++i) {
        all_nodes[i].num_connect = 0;
        all_nodes[i].connect_list = std::vector<char>();
        all_nodes[i].ID = static_cast<char>(i);
    }
    // no 5,6. 7 should be first
    while (ifs.getline(str, 256 - 1))
    {
        set_node(all_nodes, str);
    }

    tsort(all_nodes);
    return 1;
}
