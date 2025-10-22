package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadInt(r *bufio.Reader) int {
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	res, _ := strconv.Atoi(line)

	return res
}

func ReadSlice(r *bufio.Reader) []int {
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	sl := strings.Split(line, " ")
	res := make([]int, len(sl))

	for i := range sl {
		res[i], _ = strconv.Atoi(string(sl[i]))
	}

	return res
}

type Pair struct {
	first  int
	second int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := ReadInt(reader)
	s := ReadSlice(reader)
	a := ReadSlice(reader)
	p := make([]Pair, n)
	for i := range n {
		p[i].first = s[i]
		p[i].second = a[i]
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].first < p[j].first
	})

	var weight_sum, pref_sum int64
	var seasons int

	for i := range n {
		weight_sum += int64(p[i].second)
	}

	for i := range n {
		pref_sum += int64(p[i].second)
		if pref_sum*2 >= weight_sum {
			seasons = p[i].first
			break
		}
	}

	extra_salary := int64(0)

	for i := range n {
		extra_salary += int64(Abs(seasons-p[i].first)) * int64(p[i].second)
	}

	fmt.Println(seasons, extra_salary)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
