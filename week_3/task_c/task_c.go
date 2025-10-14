package main

import (
	"fmt"
	"math"
)

type Word struct {
	a, b int
}

func main() {
	var n, W, H int

	fmt.Scan(&n, &W, &H)

	words := make([]Word, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&words[i].a, &words[i].b)
	}

	var left, right float64
	right = 1e9
	for math.Abs(left-right) > 1e-6 {
		mid := (left + right) / 2
		if canPrint(mid, float64(W), float64(H), words) {
			left = mid
		} else {
			right = mid
		}
	}

	fmt.Println(left)

}

func canPrint(k, W, H float64, words []Word) bool {
	var (
		line_width   float64
		total_height float64
		line_height  float64
		first_word   bool
	)

	first_word = true

	for i := range words {
		w := k * float64(words[i].a)
		h := k * float64(words[i].b)

		if (w > W && math.Abs(w-W) > 1e-7) || (h > H && math.Abs(h-H) > 1e-7) {
			return false
		}

		if first_word {
			line_height = h
			line_width = w
			first_word = false
		} else {
			if (line_width+w < W || math.Abs(line_width+w-W) < 1e-7) && math.Abs(line_height-h) < 1e-7 {
				line_width += w
			} else {
				total_height += line_height
				if total_height > H {
					return false
				}
				line_height = h
				line_width = w
			}
		}
	}

	total_height += line_height

	return H > total_height || math.Abs(H-total_height) < 1e-7
}
