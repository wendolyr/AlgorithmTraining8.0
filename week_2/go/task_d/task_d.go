package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	var n int
	fmt.Fscan(reader, &n)

	dict := make(map[string]int)
	for i := 0; i < n; i++ {
		var temp string
		fmt.Fscan(reader, &temp)
		dict[temp]++
	}

	l := len(s)
	dp := make([]bool, l+1)
	prev := make([]int, l+1)
	dp[0] = true

	for i := 1; i <= l; i++ {
		prev[i] = -1
		for j := 0; j < i; j++ {
			if dp[j] {
				cur := s[j:i]
				_, exists := dict[cur]
				if exists {
					dp[i] = true
					prev[i] = j
					break
				}

			}
		}
	}

	pos := l
	var res []string

	for pos > 0 {
		beg := prev[pos]
		if beg >= 0 {
			res = append(res, s[beg:pos])
		}
		pos = beg
	}

	for i := range res {
		r_ind := len(res) - 1 - i
		fmt.Fprint(writer, res[r_ind], " ")
	}

}
