package aoc2022

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day4(data string) {
	pairs := strings.Split(data, "\r\n")

	countContains := 0
	countOverlap := 0
	for _, pair := range pairs {
		r, _ := regexp.Compile(`\d+`)
		strings := r.FindAllString(pair, -1)
		ids := []int{}
		for _, str := range strings {
			id, _ := strconv.Atoi(str)
			ids = append(ids, id)
		}
		if checkContains(ids) {
			countContains++
		}
		if checkOverlap(ids) {
			countOverlap++
		}
	}

	fmt.Println("Answer 1:", countContains)
	fmt.Println("Answer 2:", countOverlap)
}

func checkContains(ids []int) bool {
	return (ids[0] >= ids[2] && ids[1] <= ids[3]) || (ids[0] <= ids[2] && ids[1] >= ids[3])
}

func checkOverlap(ids []int) bool {
	return checkContains(ids) || (ids[0] >= ids[2] && ids[0] <= ids[3]) || (ids[1] >= ids[2] && ids[1] <= ids[3])
}
