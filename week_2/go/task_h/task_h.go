package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	prime := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		prime[i] = true
	}

	for i := 2; i <= n; i++ {
		if prime[i] {
			for j := 2 * i; j <= n; j = j + i {
				prime[j] = false
			}
		}
	}

	dp := make([]bool, n+1)

	for i := 1; i <= n; i++ {
		for j := 1; j <= 3; j++ {
			rem := i - j
			if rem == 0 || !prime[rem] {
				if !dp[rem] {
					dp[i] = true
					break
				}
			}
		}
	}

	if dp[n] {
		fmt.Println(1)
	} else {
		fmt.Println(2)
	}
}
