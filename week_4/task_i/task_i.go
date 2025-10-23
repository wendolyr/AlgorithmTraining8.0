package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x int
	y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, d int

	fmt.Fscan(reader, &n, &d)

	trees := make(map[Point]bool, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		trees[Point{x, y}] = true
	}

	dx_dy_pairs := make(map[Point]struct{})
	sqrt_d := int(math.Sqrt(float64(d)))
	for dx := 0; dx <= sqrt_d; dx++ {

		rem := d - dx*dx
		if rem < 0 {
			continue
		}

		dy := int(math.Sqrt(float64(rem)))
		if dy*dy == rem {
			dx_dy_pairs[Point{dx, dy}] = struct{}{}
			dx_dy_pairs[Point{dx, -dy}] = struct{}{}
			dx_dy_pairs[Point{-dx, dy}] = struct{}{}
			dx_dy_pairs[Point{-dx, -dy}] = struct{}{}
		}
	}

	result := int64(0)
	for i := range trees {
		for coord := range dx_dy_pairs {
			another_tree := Point{i.x + coord.x, i.y + coord.y}
			if _, exists := trees[another_tree]; exists {
				result++
			}
		}
	}

	fmt.Println(result / 2)
}
