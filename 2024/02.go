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
		// fmt.Printf("%v\n", levels)
		if levels[0] == levels[1] {
			continue
		}
		safeLevel := true
		increasing := levels[0] < levels[1]
		// fmt.Printf("increasing %v\n", increasing)
		for i := 1; i < len(levels); i++ {
			diff := levels[i-1] - levels[i]
			// fmt.Printf("diff %v\n", diff)
			if increasing != (diff < 0) || utils.Abs(diff) < 1 || utils.Abs(diff) > 3 {
				safeLevel = false
			}
		}

		if safeLevel {
			safe++
		}
	}

	fmt.Printf("%v\n", safe)
}
