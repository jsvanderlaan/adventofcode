package aoc2022

import (
	"fmt"
	"strconv"
	"strings"
)

func Day5(data string) {
	parts := strings.Split(data, "\r\n\r\n")
	crates := strings.Split(parts[0], "\r\n")
	moves := strings.Split(parts[1], "\r\n")

	crates = reverse(crates)
	numberOfStacks := len(crates[0])

	stacks := make([][]byte, numberOfStacks)
	for _, crate := range crates {
		for i := 0; i < len(crate); i++ {
			item := crate[i]
			if item != ' ' {
				stacks[i] = append(stacks[i], item)
			}
		}
	}

	for _, move := range moves {
		instructions := strings.Split(move, " ")
		numberOfCratesToMove, _ := strconv.Atoi(instructions[0])
		from, _ := strconv.Atoi(instructions[1])
		to, _ := strconv.Atoi(instructions[2])
		from--
		to--
		// for i := 0; i < numberOfCratesToMove; i++ {
		// 	n := len(stacks[from]) - 1
		// 	stacks[to] = append(stacks[to], stacks[from][n])
		// 	stacks[from] = stacks[from][:n]
		// }
		n := len(stacks[from])
		stacks[to] = append(stacks[to], stacks[from][n-numberOfCratesToMove:n]...)
		stacks[from] = stacks[from][:n-numberOfCratesToMove]
	}

	answer1 := ""
	for _, stack := range stacks {
		n := len(stack) - 1
		answer1 += string(stack[n])
	}
	fmt.Println(answer1)
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
