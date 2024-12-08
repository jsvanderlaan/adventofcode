package aoc2024

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day5(data string) {
	splitData := strings.Split(data, "\r\n\r\n")
	rulesStr := splitData[0]
	rules := map[int][]int{}
	for _, rule := range strings.Split(rulesStr, "\r\n") {
		split := strings.Split(rule, "|")
		lower, _ := strconv.Atoi(split[0])
		upper, _ := strconv.Atoi(split[1])
		if _, ok := rules[lower]; !ok {
			rules[lower] = []int{upper}
		} else {
			rules[lower] = append(rules[lower], upper)
		}
	}

	updatesStr := splitData[1]
	updates := [][]int{}
	for i, update := range strings.Split(updatesStr, "\r\n") {
		updates = append(updates, []int{})
		for _, page := range strings.Split(update, ",") {
			p, _ := strconv.Atoi(page)
			updates[i] = append(updates[i], p)
		}
	}

	correctIndeces := []int{}
	sum2 := 0
	for index, update := range updates {
		c := make([]int, len(update))
		copy(c, update)
		slices.SortFunc(c, func(a, b int) int {
			if _, ok := rules[a]; ok && slices.Contains(rules[a], b) {
				return -1
			}

			if _, ok := rules[b]; ok && slices.Contains(rules[b], a) {
				return 1
			}
			return 0
		})
		// fmt.Printf("c %v\n", c)
		// fmt.Printf("update %v\n", update)
		correct := true
		for i, page := range update {
			if c[i] != page {
				correct = false
				sum2 += c[len(c)/2]
				break
			}
		}
		if correct {
			correctIndeces = append(correctIndeces, index)
		}
	}

	sum := 0
	for _, index := range correctIndeces {
		sum += updates[index][len(updates[index])/2]
	}

	fmt.Println(sum2)
}
