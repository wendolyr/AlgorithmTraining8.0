package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Car struct {
	x, y, vx, vy, index int
	is_active           bool
}

type Event struct {
	time          float64
	event_type    int
	car_one_index int
	car_two_index int
}

const (
	CarsCollision = iota
	WallCollision
	Finish
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, l, w int

	fmt.Fscan(reader, &n, &l, &w)

	cars := make([]Car, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &cars[i].x, &cars[i].y, &cars[i].vx, &cars[i].vy)
		cars[i].index = i
		cars[i].is_active = true
	}

	events := make([]Event, 0)
	for i := 0; i < n; i++ {
		var wall_collision_time float64
		if cars[i].vy > 0 {
			wall_collision_time = float64(w-cars[i].y) / float64(cars[i].vy)
		} else if cars[i].vy < 0 {
			wall_collision_time = float64(-cars[i].y) / float64(cars[i].vy)
		}

		if wall_collision_time > 0 {
			events = append(events, Event{wall_collision_time, WallCollision, i, -1})
		}

		for j := i + 1; j < n; j++ {
			if t, have_collision := CollisionTime(cars[i], cars[j]); have_collision {
				events = append(events, Event{t, CarsCollision, i, j})
			}
		}

		if cars[i].vx > 0 {
			finish_time := float64(l-cars[i].x) / float64(cars[i].vx)
			if finish_time > 1e-6 {
				events = append(events, Event{finish_time, Finish, i, -1})
			}
		}
	}

	sort.Slice(events, func(i, j int) bool {
		if math.Abs(events[i].time-events[j].time) < 1e-6 {
			return events[i].event_type < events[j].event_type
		}

		return events[i].time < events[j].time
	})

	winners := make([]int, 0)
	winner_time := float64(0)

	last_time_crashed := float64(0)
	cars_in_last_collission := make([]int, 0)
	for _, ev := range events {
		if ev.event_type == WallCollision {
			cars[ev.car_one_index].is_active = false
			cars_in_last_collission = cars_in_last_collission[:0]
		} else if ev.event_type == CarsCollision && cars[ev.car_one_index].is_active && cars[ev.car_two_index].is_active {
			cars[ev.car_one_index].is_active = false
			cars[ev.car_two_index].is_active = false
			last_time_crashed = ev.time
			cars_in_last_collission = append(cars_in_last_collission, ev.car_one_index, ev.car_two_index)
		} else if ev.event_type == CarsCollision && math.Abs(last_time_crashed-ev.time) < 1e-6 {
			car, in_collision := IsCarsInCurrentCollision(cars_in_last_collission, ev.car_one_index, ev.car_two_index)
			if in_collision && car == ev.car_one_index {
				cars[ev.car_two_index].is_active = false
				cars_in_last_collission = append(cars_in_last_collission, ev.car_two_index)
			} else if in_collision && car == ev.car_two_index {
				cars[ev.car_one_index].is_active = false
				cars_in_last_collission = append(cars_in_last_collission, ev.car_one_index)
			}
		} else if ev.event_type == CarsCollision {
			cars_in_last_collission = cars_in_last_collission[:0]
		} else if ev.event_type == Finish && len(winners) == 0 && cars[ev.car_one_index].is_active {
			winners = append(winners, ev.car_one_index+1)
			winner_time = ev.time
			cars_in_last_collission = cars_in_last_collission[:0]
		} else if ev.event_type == Finish && math.Abs(ev.time-winner_time) < 1e-6 && cars[ev.car_one_index].is_active {
			winners = append(winners, ev.car_one_index+1)
			cars_in_last_collission = cars_in_last_collission[:0]
		}
	}

	sort.Ints(winners)

	fmt.Println(len(winners))
	for _, v := range winners {
		fmt.Printf("%d ", v)
	}

}

func CollisionTime(c1, c2 Car) (float64, bool) {
	dx := float64(c1.x - c2.x)
	dy := float64(c1.y - c2.y)
	dvx := float64(c2.vx - c1.vx)
	dvy := float64(c2.vy - c1.vy)

	if math.Abs(dvx) > 1e-6 && math.Abs(dvy) > 1e-6 {
		tx := dx / dvx
		ty := dy / dvy
		if math.Abs(tx-ty) < 1e-6 && tx > 1e-6 {
			return tx, true
		}

		return 0, false
	}

	if math.Abs(dvx) < 1e-6 {
		if math.Abs(dx) > 1e-6 {
			return 0, false
		}
		if math.Abs(dvy) > 1e-6 {
			t := dy / dvy
			if t > 1e-6 {
				return t, true
			}
		}

		return 0, false
	}

	if math.Abs(dvy) < 1e-6 {
		if math.Abs(dy) > 1e-6 {
			return 0, false
		}

		if math.Abs(dvx) > 1e-6 {
			t := dx / dvx
			if t > 1e-6 {
				return t, true
			}
		}
		return 0, false
	}

	return 0, false
}

func IsCarsInCurrentCollision(cars []int, first, second int) (int, bool) {
	for _, ind := range cars {
		if ind == first || ind == second {
			return ind, true
		}
	}

	return -1, false
}
