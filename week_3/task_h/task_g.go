package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	population := 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		population += a[i]
	}

	paths := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		var v, u int
		fmt.Scan(&v, &u)

		paths[v] = append(paths[v], u)
		paths[u] = append(paths[u], v)
	}

	child_sum := make([]int, n+1)
	parent := make([]int, n+1)

	var dfs func(v, p int)

	dfs = func(v, p int) {
		parent[v] = p
		child_sum[v] = a[v]
		for _, dest := range paths[v] {
			if dest == p {
				continue
			}
			dfs(dest, v)
			child_sum[v] += child_sum[dest]
		}
	}
	dfs(1, -1)

	result := 1
	min_max_queue := population

	for i := 1; i <= n; i++ {
		max_queue := 0
		for _, dest := range paths[i] {
			var sum int
			if dest == parent[i] {
				sum = population - child_sum[i]
			} else {
				sum = child_sum[dest]
			}

			max_queue = max(max_queue, sum)
		}

		if max_queue < min_max_queue {
			min_max_queue = max_queue
			result = i
		}
	}

	fmt.Println(result)

}
