package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	BeginRace = 1
	EndRace   = -1
)

type Time struct {
	hour   int
	minute int
}

type Race struct {
	t      Time
	status int
	office int
}

type ByTime []Race

func (a ByTime) Len() int {
	return len(a)
}
func (a ByTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByTime) Less(i, j int) bool {
	if a[i].t.hour == a[j].t.hour {
		if a[i].t.minute == a[j].t.minute {
			return a[i].status < a[j].status
		}
		return a[i].t.minute < a[j].t.minute
	}
	return a[i].t.hour < a[j].t.hour
}

func ReadTimeRange(r *bufio.Reader) (Time, Time) {
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	split := strings.Split(line, "-")

	begin := ReadTime(split[0])
	end := ReadTime(split[1])

	return begin, end
}

func ReadTime(time string) Time {
	split := strings.Split(time, ":")
	res := Time{}
	res.hour, _ = strconv.Atoi(split[0])
	res.minute, _ = strconv.Atoi(split[1])

	return res
}

func ReadInt(r *bufio.Reader) int {
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	res, _ := strconv.Atoi(line)

	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n := ReadInt(reader)
	first_office := make([]Race, 2*n)
	for i := 0; i < 2*n; i++ {
		begin, end := ReadTimeRange(reader)
		first_office[i] = Race{begin, BeginRace, 1}
		i++
		first_office[i] = Race{end, EndRace, 2}
	}

	m := ReadInt(reader)
	second_office := make([]Race, 2*m)
	for i := 0; i < 2*m; i++ {
		begin, end := ReadTimeRange(reader)
		second_office[i] = Race{begin, BeginRace, 2}
		i++
		second_office[i] = Race{end, EndRace, 1}
	}

	all_events := append(first_office, second_office...)
	sort.Sort(ByTime(all_events))

	busses_first, busses_second := 0, 0
	for _, event := range all_events {
		if event.status == BeginRace {
			if event.office == 1 {
				busses_first = max(0, busses_first-1)
			} else {
				busses_second = max(0, busses_second-1)
			}
		} else {
			if event.office == 1 {
				busses_first += 1
			} else {
				busses_second += 1
			}
		}
	}

	fmt.Fprintln(writer, busses_first+busses_second)

}
