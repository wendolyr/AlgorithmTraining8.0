package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Pair struct {
	b int
	t int
}

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fscan(reader, &n)

	a := make([]Pair, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i].b, &a[i].t)
	}

	fmt.Fscan(reader, &m)
	cars := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &cars[i])
	}

	sorted := make([]int, m)
	copy(sorted, cars)
	slices.Sort(sorted)

	car_map := make(map[int]int64)
	current := 0
	for _, car := range sorted {
		for current < n-1 && car > a[current+1].b {
			current++
		}

		res := int64(car) * int64(a[current].t)
		car_map[car] = res
	}

	for _, car := range cars {
		fmt.Fprintln(writer, car_map[car])
	}
}
