package aoc2022

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type sensor struct {
	pos    point
	radius int
}

func (p_1 *point) Equals(p_2 point) bool {
	return p_1.x == p_2.x && p_1.y == p_2.y
}

func Day15(data string) {
	const y int = 2000000

	y_min := math.MaxInt32
	y_max := math.MinInt32
	x_min := math.MaxInt32
	x_max := math.MinInt32

	sensors := []sensor{}
	beacons := []point{}
	sensorDatas := strings.Split(data, "\r\n")
	for _, s := range sensorDatas {
		r, _ := regexp.Compile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
		matches := r.FindAllSubmatch([]byte(s), -1)
		s_x, _ := strconv.Atoi(string(matches[0][1]))
		s_y, _ := strconv.Atoi(string(matches[0][2]))
		b_x, _ := strconv.Atoi(string(matches[0][3]))
		b_y, _ := strconv.Atoi(string(matches[0][4]))
		dist := manDist(s_x, s_y, b_x, b_y)
		beacon := point{b_x, b_y}
		sensor := sensor{point{s_x, s_y}, dist}
		if s_x-dist < x_min {
			x_min = s_x - dist
		}
		if s_x+dist > x_max {
			x_max = s_x + dist
		}
		if s_y-dist < y_min {
			y_min = s_y - dist
		}
		if s_y+dist > y_max {
			y_max = s_y + dist
		}

		sensors = append(sensors, sensor)
		beacons = append(beacons, beacon)
	}

	count := 0
	for x := x_min; x <= x_max; x++ {
		for _, s := range sensors {
			if manDist(x, y, s.pos.x, s.pos.y) <= s.radius {
				count++
				break
			}
		}
	}

	beaconsOnY := []point{}
	for _, b := range beacons {
		if b.y == y {
			contains := false
			for _, beaconOnY := range beaconsOnY {
				if beaconOnY.Equals(b) {
					contains = true
				}
			}
			if !contains {
				beaconsOnY = append(beaconsOnY, b)
			}
		}
	}

	fmt.Println("Answer 1:", count-len(beaconsOnY), count, len(beaconsOnY))

	// for _, s := range sensors {

	// 	fmt.Println(s)
	// }

	// for _, b := range beacons {
	// 	fmt.Println(b)
	// }
}

func manDist(x_1 int, y_1 int, x_2 int, y_2 int) int {
	return int(math.Abs(float64(x_1-x_2))) + int(math.Abs(float64(y_1-y_2)))
}
