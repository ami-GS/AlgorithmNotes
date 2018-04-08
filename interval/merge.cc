#include <vector>
#include <iostream>

using namespace std;

//https://leetcode.com/problems/merge-intervals/description/

// ask to intervier
// - the intervals are sorted by 'start'? -> if no, need sort by 'start'
// - how many overrapping appears? -> over 1, then need to consider scale
struct Interval {
    int start;
    int end;
    Interval() : start(0), end(0) {}
    Interval(int s, int e) : start(s), end(e) {}
};

vector<Interval> merge(vector<Interval> &intervals) {
    vector<Interval> ret;
    int count = 0;

    while (count < intervals.size()) {
        int newstart = 0;
        int newend = 0;
        while (count < intervals.size()-1 && intervals[count].end < intervals[count+1].start) {
            ret.push_back(intervals[count++]);
        }
        if (count == intervals.size()-1) {
            ret.push_back(intervals[count]);
            return ret;
        }

        newstart = intervals[count++].start;
        while (count < intervals.size()-1 && intervals[count].end > intervals[count+1].start) {
            count++;
        }
        newend = intervals[count++].end;
        ret.push_back(Interval(newstart, newend));
    }
    return ret;
}

int main() {
    vector<Interval> intervals{
        Interval(1, 3),
        Interval(2, 6),
        Interval(8, 10),
        Interval(15, 18),
        Interval(19, 21),
        Interval(22, 24),
        Interval(23, 28),
        Interval(25, 29),
        Interval(26, 34)
            };
    auto out = merge(intervals);
    cout << "[　";
    for (int i = 0; i < out.size(); ++i) {
        cout << "[" << out[i].start << ", " << out[i].end << "] ";
    }
    cout << "]" << endl;
}
