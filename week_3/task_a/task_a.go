package main

import "fmt"

func main() {
	var a, b, S, left, right int64

	fmt.Scan(&a, &b, &S)

	left = max(a, b)
	right = S + a + b + 1

	for left < right {
		mid := left + (right-left)/2
		val := (mid - a) * (mid - b)
		if val < S {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if (left-a)*(left-b) != S {
		fmt.Println(-1)
		return
	}

	fmt.Println(left)
}
