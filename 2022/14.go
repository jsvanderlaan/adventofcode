package aoc2022

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day14(data string) {
	columns := map[int]map[int]byte{}

	rockSnakes := strings.Split(data, "\r\n")
	for _, rockSnake := range rockSnakes {
		corners := strings.Split(rockSnake, " -> ")
		for i := 0; i < len(corners)-1; i++ {
			sc_x, sc_y, _ := strings.Cut(corners[i], ",")
			sn_x, sn_y, _ := strings.Cut(corners[i+1], ",")

			c_x, _ := strconv.Atoi(sc_x)
			c_y, _ := strconv.Atoi(sc_y)
			n_x, _ := strconv.Atoi(sn_x)
			n_y, _ := strconv.Atoi(sn_y)

			if c_x == n_x {
				var r []int
				if c_y >= n_y {
					r = makeRange(n_y, c_y)
				} else {
					r = makeRange(c_y, n_y)
				}
				_, ok := columns[c_x]
				if !ok {
					columns[c_x] = map[int]byte{}
				}
				for _, y := range r {
					columns[c_x][y] = 0
				}
			} else if c_y == n_y {
				var r []int
				if c_x >= n_x {
					r = makeRange(n_x, c_x)
				} else {
					r = makeRange(c_x, n_x)
				}
				for _, x := range r {
					_, ok := columns[x]
					if !ok {
						columns[x] = map[int]byte{}
					}
					columns[x][c_y] = 0
				}
			}
		}
	}

	sand_counter := 0
	stop := false
	for !stop {
		sand_counter++
		curr_x := 500
		curr_y := 0
		can_move := true
		for can_move {
			// fmt.Println("NEXT MOVE: curr_x", curr_x, "curr_y", curr_y)
			_, col_has_stop := columns[curr_x]
			if !col_has_stop {
				// fmt.Println("\tSTOP: no stops in column", curr_x)
				stop = true
				break
			}
			stop_point := getStopPointOfColumn(columns[curr_x], curr_y)
			// fmt.Println("\tstop_point", stop_point)

			if stop_point == -2 {
				// fmt.Println("\tSTOP: no stop point for curr_x", curr_x, "curr_y", curr_y)
				stop = true
				break
			}

			// check if can go left
			_, next_left_has_stop := columns[curr_x-1]
			if !next_left_has_stop {
				stop = true
				break
			}
			_, next_left_occupied := columns[curr_x-1][stop_point+1]
			if !next_left_occupied {
				// fmt.Println("\tgo left")
				curr_x -= 1
				curr_y = stop_point + 1
				continue
			}

			// check if can go right
			_, next_right_has_stop := columns[curr_x+1]
			if !next_right_has_stop {
				stop = true
				break
			}
			_, next_right_occupied := columns[curr_x+1][stop_point+1]
			if !next_right_occupied {
				// fmt.Println("\tgo right")
				curr_x += 1
				curr_y = stop_point + 1
				continue
			}

			// stop the sand!
			// fmt.Println("\tplace sand at x", curr_x, "y", stop_point)
			columns[curr_x][stop_point] = 1
			can_move = false
		}
	}

	// fmt.Println(columns)
	fmt.Println("Answer 1:", sand_counter-1)
}

func getStopPointOfColumn(m map[int]byte, start int) int {
	min := math.MaxInt32
	set := false
	for k := range m {
		if k > start && k < min {
			min = k
			set = true
		}
	}
	if !set {
		return -2
	}
	return min - 1
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
