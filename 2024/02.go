package aoc2024

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day2(data string) {
	reports := strings.Split(data, "\r\n")

	safe := 0

	for _, report := range reports {
		levels := []int{}
		for _, level := range strings.Split(report, " ") {
			l, _ := strconv.Atoi(level)
			levels = append(levels, l)
		}

		errorIndex := areLevelsSafe(levels)
		if errorIndex == -1 {
			safe++
			continue
		}

		corrSafe := false
		for i := 0; i < len(levels); i++ {
			corr := []int{}
			for il, l := range levels {
				if il != i {
					corr = append(corr, l)
				}
			}
			if areLevelsSafe(corr) == -1 {
				corrSafe = true
			}
		}
		if corrSafe {
			safe++
			continue
		}
	}

	fmt.Printf("%v\n", safe)
}

// 515 too low
// 526 too low

func areLevelsSafe(levels []int) int {
	if levels[0] == levels[1] {
		return 1
	}

	increasing := levels[0] < levels[1]
	for i := 1; i < len(levels); i++ {
		diff := levels[i-1] - levels[i]
		if increasing != (diff < 0) || utils.Abs(diff) < 1 || utils.Abs(diff) > 3 {
			return i
		}
	}

	return -1
}
