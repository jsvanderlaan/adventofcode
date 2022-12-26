package aoc2022

import (
	"fmt"
	"strconv"
	"strings"
)

// type monkey struct {
// 	has_operator bool
// 	value rune
// 	left *monkey
// 	right *monkey
// }

// type monkey struct {
// 	key   string
// 	value int
// }

func Day21(data string) {
	monkeys := map[string]string{}
	monkeyValues := map[string]int{}
	for _, monkey := range strings.Split(data, "\r\n") {
		key, value, _ := strings.Cut(monkey, ": ")
		monkeys[key] = value
	}

	stack := []string{"root"}

	for len(stack) > 0 {
		next := stack[len(stack)-1]

		value := monkeys[next]
		number, err := strconv.Atoi(value)
		if err == nil {
			monkeyValues[next] = number
			stack = stack[:len(stack)-1]
			continue
		}

		values := strings.Split(value, " ")
		left := values[0]
		leftValue, leftCalculated := monkeyValues[left]
		if !leftCalculated {
			stack = append(stack, left)
			continue
		}

		operator := values[1]
		right := values[2]
		rightValue, rightCalculated := monkeyValues[right]
		if !rightCalculated {
			stack = append(stack, right)
			continue
		}

		switch operator {
		case "+":
			monkeyValues[next] = leftValue + rightValue
		case "/":
			monkeyValues[next] = leftValue / rightValue
		case "*":
			monkeyValues[next] = leftValue * rightValue
		case "-":
			monkeyValues[next] = leftValue - rightValue
		default:
			panic("Unknown operator")
		}
		stack = stack[:len(stack)-1]
	}

	fmt.Println("Answer 1:", monkeyValues["root"])
}
