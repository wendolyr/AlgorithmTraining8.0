package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	p := make([]int, n)
	a := make([]int, n)
	children := make([][]int, n)
	p[0] = -1

	for i := 1; i < n; i++ {
		fmt.Scan(&p[i])
		children[p[i]] = append(children[p[i]], i)
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sum := make([]int, n)
	for i := 0; i < n; i++ {
		for _, t := range children[i] {
			sum[i] += a[t]
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		result += abs((-a[i] + sum[i]))
	}

	fmt.Println(result)

}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
