package main

import "fmt"

type Node struct {
	start, finish int
}

func main() {
	var n, m, root int

	fmt.Scan(&n)

	parent := make([]int, n+1)
	children := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&parent[i])

		if parent[i] != 0 {
			children[parent[i]] = append(children[parent[i]], i)
		} else {
			root = i
		}
	}

	timer := 0
	nodes := make([]Node, n+1)

	var count func(int)
	count = func(cur int) {
		timer++
		nodes[cur].start = timer

		for _, child := range children[cur] {
			count(child)
		}

		nodes[cur].finish = timer
	}
	count(root)

	fmt.Scan(&m)

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		if isParent(nodes, a, b) {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}

}

func isParent(nodes []Node, a, b int) bool {
	return nodes[a].start <= nodes[b].start && nodes[b].finish <= nodes[a].finish
}
