package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	n := ScanInt(scanner)

	a := make([]int, n+1)
	b := make([]int, n+1)
	pref_a := make([]int64, n+1)
	pref_b := make([]int64, n+1)

	for i := range n {
		a[i] = ScanInt(scanner)
	}
	for i := range n {
		b[i] = ScanInt(scanner)

	}

	for i := 1; i <= n; i++ {
		pref_a[i] = pref_a[i-1] + int64(a[i-1])
		pref_b[i] = pref_b[i-1] + int64(b[i-1])
	}

	if pref_a[n] > pref_b[n] {
		fmt.Println(-1)
		return
	}

	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if check(mid, pref_a, pref_b) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	fmt.Println(right)
}

func ScanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func check(k int, pref_a, pref_b []int64) bool {
	n := len(pref_a) - 1

	min_v := int64(1e18)
	for i := 1; i <= n; i++ {
		l := max(1, i-k)
		r := min(n, i+k)

		cur := pref_a[i-1] - pref_b[l-1]

		min_v = min(min_v, cur)

		rem := pref_a[i] - pref_b[r]

		if rem > min_v {
			return false
		}

	}

	return true
}
