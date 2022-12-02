package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./data/01.data")
	check(err)
	var sliced slice[string, []string] = strings.Split(string(data), "\r\n\r\n")
	elves := sliced.Map(func(x string) []string { return strings.Split(x, "\r\n") })

	elveCals := []int{}
	for _, elve := range elves {
		sum := 0
		for _, food := range elve {
			cals, err := strconv.Atoi(food)
			check(err)
			sum += cals
		}
		elveCals = append(elveCals, sum)
	}

	fmt.Println("Answer 1", max(elveCals))

	sort.Ints(elveCals)
	topThree := elveCals[len(elveCals)-3:]
	sum := 0
	for _, topElve := range topThree {
		sum += topElve
	}

	fmt.Println("Answer 2", sum)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func max(arr []int) int {
	max := math.MinInt32
	for _, item := range arr {
		if item > max {
			max = item
		}
	}
	return max
}
