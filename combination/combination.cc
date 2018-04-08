#include <string>
#include <vector>
#include <iostream>
#include <stack>

void make_all_combinations(std::vector<char> str, std::vector<char> buff, std::vector<std::vector<char>> *ret, int i) {
    if (str.size() == i) {
        ret->push_back(buff);
        return;
    }

    make_all_combinations(str, buff, ret, i+1);
    buff.push_back(str[i]);
    make_all_combinations(str, buff, ret, i+1);
}

//http://blog.gainlo.co/index.php/2017/01/05/uber-interview-questions-permutations-array-arrays/
void uber(std::vector<std::vector<char>> list_list, std::stack<char> buff, std::vector<std::stack<char>> *ret, int i) {
    if (list_list.size() == i) {
        ret->push_back(buff);
        return;
    }
    for (int j = 0; j < list_list[i].size();  ++j) {
        buff.push(list_list[i][j]);
        uber(list_list, buff, ret, i+1);
        buff.pop();
    }
}

//https://leetcode.com/problems/permutations/description/
void permutate(std::vector<int> list, std::vector<int> buff, std::vector<std::vector<int>> *ret, int depth) {
    if (list.size() == depth) {
        ret->push_back(buff);
        return;
    }

    for (int i = 0; i < list.size(); ++i) {
        buff[depth] = list[i];
        permutate(list, buff, ret, depth+1);
    }
}

int main() {
    std::vector<char> v{'a', 'b', 'c', 'e', 'f', 'g'}, c;
    std::vector<std::vector<char>> ret;
    make_all_combinations(v, c, &ret, 0);

    for (int i = 0; i < ret.size(); ++i) {
        for (int j = 0; j < ret[i].size(); ++j) {
            std::cout << ret[i][j];
        }
        std::cout << std::endl;
    }

    std::vector<std::vector<char>> list_list{{'a', 'b', 'c'}, {'d'}, {'e', 'f'}};
    std::stack<char> cc;
    //manual stack is better to print
    std::vector<std::stack<char>> retret;
    uber(list_list, cc, &retret, 0);

    for (int i = 0; i < retret.size(); ++i) {
        int size = retret[i].size();
        for (int j = 0; j < size; ++j) {
            std::cout << retret[i].top();
            retret[i].pop();
        }
        std::cout << std::endl;
    }

    std::vector<int> list{1,2,3}, ccc(3);
    std::vector<std::vector<int>> retretret;
    permutate(list, ccc, &retretret, 0);
    for (int i = 0; i < retretret.size(); ++i) {
        for (int j = 0; j < retretret[i].size(); ++j) {
            std::cout << retretret[i][j];
        }
        std::cout << std::endl;
    }
}
