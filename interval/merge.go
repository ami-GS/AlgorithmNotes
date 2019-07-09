package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	start int
	end   int
}

func merge(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	out := make([]Interval, 0, len(intervals))

	for i := 0; i < len(intervals); {
		mergeIdx := i
		for j := i + 1; j < len(intervals); j++ {
			if intervals[mergeIdx].end >= intervals[j].start {
				mergeIdx = j
			} else {
				break
			}
		}
		end := intervals[mergeIdx].end
		if end < intervals[i].end {
			end = intervals[i].end
		}
		out = append(out, Interval{intervals[i].start, end})

		i = mergeIdx + 1
	}
	return out
}

func main() {
	intervals := []Interval{
		Interval{1, 3},
		Interval{2, 6},
		Interval{8, 10},
		Interval{15, 18},
		Interval{19, 21},
		Interval{22, 24},
		Interval{23, 28},
		Interval{25, 29},
		Interval{26, 34},
	}

	out := merge(intervals)
	fmt.Println(out)

	intervals = []Interval{
		Interval{1, 4},
		Interval{2, 3},
	}

	out = merge(intervals)
	fmt.Println(out)

	intervals = []Interval{
		Interval{1, 4},
		Interval{0, 2},
		Interval{3, 5},
	}

	out = merge(intervals)
	fmt.Println(out)

	intervals = []Interval{
		Interval{0, 2},
		Interval{1, 4},
		Interval{3, 5},
	}

	out = merge(intervals)
	fmt.Println(out)

	intervals = []Interval{
		Interval{4, 6},
		Interval{0, 4},
		Interval{2, 3},
		Interval{7, 12},
	}

	out = merge(intervals)
	fmt.Println(out)

	intervals = []Interval{
		Interval{2, 3},
		Interval{4, 5},
		Interval{6, 7},
		Interval{8, 9},
		Interval{1, 10},
	}

	out = merge(intervals)
	fmt.Println(out)

}
