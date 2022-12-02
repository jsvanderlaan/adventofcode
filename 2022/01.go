package aoc2022

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day1(data string) {
	var sliced utils.Slice[string, []string] = strings.Split(string(data), "\r\n\r\n")
	elves := sliced.Map(func(x string) []string { return strings.Split(x, "\r\n") })

	elveCals := []int{}
	for _, elve := range elves {
		sum := 0
		for _, food := range elve {
			cals, err := strconv.Atoi(food)
			utils.Check(err)
			sum += cals
		}
		elveCals = append(elveCals, sum)
	}

	fmt.Println("Answer 1", utils.Max(elveCals))

	sort.Ints(elveCals)
	topThree := elveCals[len(elveCals)-3:]
	sum := 0
	for _, topElve := range topThree {
		sum += topElve
	}

	fmt.Println("Answer 2", sum)
}
