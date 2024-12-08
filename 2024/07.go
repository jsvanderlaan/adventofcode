package aoc2024

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day7(data string) {
	var sum int64 = 0
	equations := strings.Split(data, "\r\n")
	for _, eq := range equations {
		split := strings.Split(eq, ": ")
		test, _ := strconv.Atoi(split[0])
		parts := []int64{}
		for _, part := range strings.Split(split[1], " ") {
			p, _ := strconv.Atoi(part)
			parts = append(parts, int64(p))
		}

		var n int64 = 0
		for n < int64(math.Pow(2, float64(len(parts)-1))) {
			if int64(test) == Calc(n, parts) {
				sum += int64(test)
				break
			}
			n++
		}
	}

	fmt.Println(sum)
	//1584652836364 too low
}

func Calc(n int64, parts []int64) int64 {
	binairy := fmt.Sprintf("%0*b", len(parts)-1, n)
	// fmt.Println(binairy)
	res := parts[0]
	for i, operator := range binairy {
		if operator == '1' {
			res += parts[i+1]
		} else {
			res *= parts[i+1]
		}
	}
	return res
}
