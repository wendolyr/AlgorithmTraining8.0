package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch n {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(4)
	default:
		a := make([]int, n+1)
		a[n] = 1
		a[n-1] = 2
		a[n-2] = 4

		for i := n - 3; i > 0; i-- {
			a[i] = a[i+1] + a[i+2] + a[i+3]
		}

		fmt.Println(a[1])
	}

}
