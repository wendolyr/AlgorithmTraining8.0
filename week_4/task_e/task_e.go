package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Route struct {
	l int
	r int
}

type Pair[F, S any] struct {
	first  F
	second S
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, k := 0, 0, int64(0)

	fmt.Fscan(reader, &n, &m, &k)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	routes := make([]Route, m)
	pref_dif := make([]int, n+2)

	for i := range m {
		fmt.Fscan(reader, &routes[i].l, &routes[i].r)
		pref_dif[routes[i].l] += 1
		pref_dif[routes[i].r+1] -= 1
	}

	pairs := make([]Pair[int, int], n+1)

	for i := 1; i <= n; i++ {
		plot_coverage := pairs[i-1].first + pref_dif[i]
		pairs[i].first = plot_coverage
		pairs[i].second = a[i]
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].first > pairs[j].first
	})

	result := int64(0)
	for i := 0; i < n; i++ {
		if k > 0 {
			repaired := min(k, int64(pairs[i].second))
			pairs[i].second -= int(repaired)
			k -= repaired
		}

		result += (int64(pairs[i].first) * int64(pairs[i].second))
	}

	fmt.Println(result)

}
