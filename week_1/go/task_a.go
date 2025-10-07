package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		vasya_sum int
		masha_sum int
		masha_max int
		vasya_min int
	)

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	a := make([]int, n)

	line, _ = reader.ReadString('\n')
	numbers := strings.Fields(line)

	vasya_min = 1 << 30

	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(numbers[i])
	}

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			vasya_sum += a[i]
			if a[i] < vasya_min {
				vasya_min = a[i]
			}
		} else {
			masha_sum += a[i]
			if a[i] > masha_max {
				masha_max = a[i]
			}
		}
	}

	if vasya_min < masha_max {
		vasya_sum += (masha_max - vasya_min)
		masha_sum += (vasya_min - masha_max)
	}

	result := vasya_sum - masha_sum

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')

}
