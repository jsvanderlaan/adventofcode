package aoc2022

import (
	"fmt"
	"strconv"
	"strings"
)

func Day10(data string) {
	rows := strings.Split(data, "\r\n")
	cycle := 0
	x := 1
	sum := 0

	for _, row := range rows {
		cycle++
		if cycle > 220 {
			break
		}
		if cycle%40-20 == 0 {
			sum += x * cycle
		}

		if row == "noop" {
			continue
		}

		cycle++
		if cycle > 220 {
			break
		}
		if cycle%40-20 == 0 {
			sum += x * cycle
		}

		_, v, _ := strings.Cut(row, " ")
		val, _ := strconv.Atoi(v)
		x += val
	}

	fmt.Println("Answer 1:", sum)
}
