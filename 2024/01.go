package aoc2024

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day1(data string) {
	rows := strings.Split(data, "\r\n")
	list1 := make([]int, len(rows))
	list2 := make([]int, len(rows))

	for i, row := range rows {
		cols := strings.Split(row, "   ")
		item1, _ := strconv.Atoi(cols[0])
		item2, _ := strconv.Atoi(cols[1])
		list1[i] = item1
		list2[i] = item2
	}

	sort.Slice(list1, func(i, j int) bool { return list1[i] < list1[j] })
	sort.Slice(list2, func(i, j int) bool { return list2[i] < list2[j] })

	total := 0
	for i := 0; i < len(rows); i++ {
		total += utils.Abs(list1[i] - list2[i])
	}

	fmt.Printf("%v\n", total)

	// solution 2
	map2 := make(map[int]int)
	for _, item := range list2 {
		map2[item] += 1
	}

	total = 0
	for _, item := range list1 {
		total += (item * map2[item])
	}

	fmt.Printf("%v\n", total)
}
