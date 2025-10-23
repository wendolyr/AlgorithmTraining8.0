package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int

	fmt.Fscan(reader, &n)

	stuff := make([]int, n)

	for i := range stuff {
		fmt.Fscan(reader, &stuff[i])
	}

	diff_arr := make([]int, 2*n)

	for i := 0; i < n; i++ {
		v := i + stuff[i]
		if v > i+1 {
			diff_arr[i+1] += 1
			diff_arr[v] -= 1
		}
	}

	pref_sum := make([]int, n)
	bonus := int64(0)

	for i := 1; i < n; i++ {
		pref_sum[i] = pref_sum[i-1] + diff_arr[i]
		bonus += int64(pref_sum[i] * stuff[i])
	}

	fmt.Println(bonus)

}
