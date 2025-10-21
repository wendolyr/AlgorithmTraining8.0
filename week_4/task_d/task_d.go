package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Result struct {
	min_diff int64
	l        int
	r        int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	inp, _ := reader.ReadString('\n')
	inp = strings.TrimSpace(inp)
	n, _ := strconv.Atoi(inp)

	inp, _ = reader.ReadString('\n')
	inp = strings.TrimSpace(inp)
	str_slice := strings.Split(inp, " ")
	tables := make([]int, n+1)
	for i := 1; i <= n; i++ {
		tables[i], _ = strconv.Atoi(str_slice[i-1])
	}

	sdv, sdm := int64(tables[1]), int64(tables[n])
	left, right := 1, n

	res := Result{min_diff: Abs(sdv - sdm), l: left, r: right}

	for left < right-1 && sdm != sdv {

		if sdm < sdv {
			right--
			sdm += int64(tables[right])
		} else {
			left++
			sdv += int64(tables[left])
		}

		diff := Abs(sdv - sdm)
		if diff < res.min_diff {
			res.min_diff = diff
			res.l = left
			res.r = right
		}
	}

	fmt.Println(res.min_diff, res.l, res.r)
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}

	return x
}
