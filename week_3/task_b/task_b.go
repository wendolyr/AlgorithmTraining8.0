package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Scan(&n)

	life_situations := make([][]int, n+1)
	distance := make([]int, n+1)
	from := make([]int, n+1)
	var queue []int

	for i := 1; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		life_situations[a] = append(life_situations[a], b)
		life_situations[b] = append(life_situations[b], a)
	}

	for i := 1; i <= n; i++ {
		if len(life_situations[i]) == 1 {
			queue = append(queue, i)
			from[i] = i
			distance[i] = 0
		} else {
			distance[i] = -1
		}
	}

	result := math.MaxInt32
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		for _, val := range life_situations[top] {
			if distance[val] == -1 {
				distance[val] = distance[top] + 1
				from[val] = from[top]
				queue = append(queue, val)
			} else if from[val] != from[top] {
				temp := distance[val] + distance[top] + 1
				result = min(result, temp)
			}
		}
	}

	fmt.Println(result)

}
