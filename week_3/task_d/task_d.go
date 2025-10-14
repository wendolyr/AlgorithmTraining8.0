package main

import (
	"fmt"
	"math"
	"sort"
)

type Pair struct {
	first, second int
}

func main() {
	var n, p int

	fmt.Scan(&n, &p)

	c := make([]Pair, n+1)
	for i := 1; i <= n; i++ {
		c[i].first = i
		fmt.Scan(&c[i].second)
	}

	sort.Slice(c, func(a, b int) bool {
		return c[a].second < c[b].second
	})

	result := 10e18
	res_i, res_j := -1, -1
	for j := 1; j <= n; j++ {
		cj := c[j].second
		need := float64(p) * float64(cj)

		left, right := 1, len(c)-1

		for left < right {
			mid := left + (right-left)/2
			if float64(c[mid].second) < need {
				left = mid + 1
			} else {
				right = mid
			}
		}

		for k := right - 1; k <= right+1; k++ {
			if k < 1 || k > n || k == j {
				continue
			}
			ci := c[k].second
			cur_diff := math.Abs(float64(ci)/float64(cj) - float64(p))
			if cur_diff < result {
				result = cur_diff
				res_i = c[k].first
				res_j = c[j].first
			}
		}
	}

	fmt.Println(res_i, res_j)
}
