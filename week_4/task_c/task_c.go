package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m, x int
	fmt.Fscan(reader, &n, &x)

	queue := make([]int, n)
	prefix := make([]int, n+1)
	prefix[0] = 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &queue[i])

		if queue[i] >= x {
			prefix[i+1] = prefix[i] + 1
		} else {
			prefix[i+1] = prefix[i]
		}
	}

	fmt.Fscan(reader, &m)
	var exited_prof int
	for i := 0; i < m; i++ {
		var command int
		fmt.Fscan(reader, &command)
		switch command {
		case 1:
			var new_p int
			fmt.Fscan(reader, &new_p)
			queue = append(queue, new_p)
			last := prefix[len(prefix)-1]
			if new_p >= x {
				prefix = append(prefix, last+1)
			} else {
				prefix = append(prefix, last)
			}
		case 2:
			if queue[0] >= x {
				exited_prof++
			}
			queue = queue[1:]
			prefix = prefix[1:]
		case 3:
			var k int
			fmt.Fscan(reader, &k)

			fmt.Println(prefix[k] - exited_prof)
		}
	}

}
