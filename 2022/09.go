package aoc2022

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day9(data string) {
	rows := strings.Split(data, "\r\n")
	x_h := 0
	y_h := 0

	x_t := 0
	y_t := 0

	visited := map[string]struct{}{}

	for _, row := range rows {
		dir, val, _ := strings.Cut(row, " ")
		times, _ := strconv.Atoi(val)

		for i := 0; i < times; i++ {
			if dir == "R" {
				x_h++
			} else if dir == "U" {
				y_h++
			} else if dir == "L" {
				x_h--
			} else if dir == "D" {
				y_h--
			}

			dist_y := y_h - y_t
			dist_x := x_h - x_t

			if math.Abs(float64(dist_y)) > 1 || math.Abs(float64(dist_x)) > 1 {
				if dist_x > 0 {
					x_t++
				} else if dist_x < 0 {
					x_t--
				}
				if dist_y > 0 {
					y_t++
				} else if dist_y < 0 {
					y_t--
				}
			}

			// fmt.Println(
			// 	"x_h",
			// 	x_h,
			// 	"y_h",
			// 	y_h,
			// 	"x_t",
			// 	x_t,
			// 	"y_t",
			// 	y_t,
			// 	visited,
			// )

			key := fmt.Sprintf("%d|%d", x_t, y_t)
			visited[key] = struct{}{}
		}
	}

	fmt.Println("Answer 1:", len(visited))

	x := make([]int, 10)
	y := make([]int, 10)

	visited = map[string]struct{}{}

	for _, row := range rows {
		dir, val, _ := strings.Cut(row, " ")
		times, _ := strconv.Atoi(val)

		for i := 0; i < times; i++ {
			for j := 0; j < 10; j++ {
				if j == 0 {
					if dir == "R" {
						x[j]++
					} else if dir == "U" {
						y[j]++
					} else if dir == "L" {
						x[j]--
					} else if dir == "D" {
						y[j]--
					}
					continue
				}

				dist_y := y[j-1] - y[j]
				dist_x := x[j-1] - x[j]

				if math.Abs(float64(dist_y)) > 1 || math.Abs(float64(dist_x)) > 1 {
					if dist_x > 0 {
						x[j]++
					} else if dist_x < 0 {
						x[j]--
					}
					if dist_y > 0 {
						y[j]++
					} else if dist_y < 0 {
						y[j]--
					}
				}
			}

			key := fmt.Sprintf("%d|%d", x[9], y[9])
			visited[key] = struct{}{}
		}
	}

	fmt.Println("Answer 2:", len(visited))
}
