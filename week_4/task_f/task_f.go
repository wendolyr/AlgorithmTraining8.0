package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Interval struct {
	start, end float64
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	var x int64
	fmt.Fscan(reader, &n, &m, &x)

	intervals := make([]Interval, 0)

	for i := 0; i < n; i++ {
		var a, b, v int64
		fmt.Fscan(reader, &a, &b, &v)
		if a > b {
			v = -v
		}

		var t_min, t_max float64

		t_min = float64(x-b) / float64(v)
		t_max = float64(x-a) / float64(v)

		if t_max < 0 {
			continue
		}
		if t_min < 0 {
			t_min = 0
		}

		intervals = append(intervals, Interval{t_min, t_max})

	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	merged_intervals := make([]Interval, 0)

	for i := 0; i < len(intervals); i++ {
		if i == 0 {
			merged_intervals = append(merged_intervals, intervals[0])
			continue
		}

		last := &merged_intervals[len(merged_intervals)-1]
		if intervals[i].start <= last.end {
			last.end = max(last.end, intervals[i].end)
		} else {
			merged_intervals = append(merged_intervals, intervals[i])
		}
	}

	for i := 0; i < m; i++ {
		var t float64
		fmt.Fscan(reader, &t)

		left, right := 0, len(merged_intervals)-1
		idx := -1
		for left <= right {
			mid := left + (right-left)/2
			if merged_intervals[mid].start <= t {
				if t <= merged_intervals[mid].end {
					idx = mid
					break
				}
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		if idx == -1 {
			fmt.Printf("%f\n", t)
		} else {
			fmt.Printf("%f\n", merged_intervals[idx].end)
		}
	}

}
