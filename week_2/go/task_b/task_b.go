package main

import "fmt"

func main() {
	var river string

	fmt.Scan(&river)

	var l_min, r_min int

	r_min = 1

	for _, c := range river {
		switch c {
		case 'L':
			r_min = min(r_min, l_min+1)
			l_min = min(l_min+1, r_min+1)
		case 'R':
			l_min = min(l_min, r_min+1)
			r_min = min(r_min+1, l_min+1)
		default:
			l_min++
			r_min++
		}
	}

	fmt.Println(min(l_min+1, r_min))
}
