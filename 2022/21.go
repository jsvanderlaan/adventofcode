package aoc2022

import (
	"fmt"
	"strconv"
	"strings"
)

type operator string

const (
	times  operator = "*"
	plus   operator = "+"
	minues operator = "-"
	devide operator = "/"
	noop   operator = ""
)

type monkey struct {
	has_value bool
	value     int
	op        operator
	left      string
	right     string
}

func Day21(data string) {
	monkeys := map[string]*monkey{}
	for _, m := range strings.Split(data, "\r\n") {
		key, value, _ := strings.Cut(m, ": ")
		number, err := strconv.Atoi(value)
		if err == nil {
			monkeys[key] = &monkey{true, number, noop, "", ""}
		} else {
			values := strings.Split(value, " ")
			monkeys[key] = &monkey{false, 0, operator(values[1]), values[0], values[2]}
		}
	}

	// evaluate(monkeys, "root")

	// fmt.Println("Answer 1:", monkeys["root"].value)

	you := "humn"
	root := monkeys["root"]
	left := root.left
	right := root.right

	youVal := 0
	if branchHasYou(monkeys, left, you) {
		evaluate(monkeys, right)
		youVal = solveForYou(monkeys, you, left, monkeys[right].value)
	}
	if branchHasYou(monkeys, right, you) {
		evaluate(monkeys, left)
		youVal = solveForYou(monkeys, you, right, monkeys[left].value)
	}
	fmt.Println("Answer 2:", youVal)
}

func solveForYou(monkeys map[string]*monkey, you string, root string, other int) int {
	if root == you {
		return other
	}

	m := monkeys[root]

	if branchHasYou(monkeys, m.right, you) {
		evaluate(monkeys, m.left)
		leftVal := monkeys[m.left].value
		newOther := other
		switch m.op {
		case plus:
			newOther -= leftVal
		case devide:
			newOther = leftVal / newOther
		case times:
			newOther /= leftVal
		case minues:
			newOther = leftVal - newOther
		default:
			panic("Unknown operator")
		}

		return solveForYou(monkeys, you, m.right, newOther)
	} else {
		evaluate(monkeys, m.right)
		rightVal := monkeys[m.right].value
		newOther := other
		switch m.op {
		case plus:
			newOther -= rightVal
		case devide:
			newOther *= rightVal
		case times:
			newOther /= rightVal
		case minues:
			newOther += rightVal
		default:
			panic("Unknown operator")
		}

		return solveForYou(monkeys, you, m.left, newOther)
	}

}

func evaluate(monkeys map[string]*monkey, root string) {
	stack := []string{root}

	for len(stack) > 0 {
		next := stack[len(stack)-1]

		m := monkeys[next]

		if m.has_value {
			stack = stack[:len(stack)-1]
			continue
		}

		left := monkeys[m.left]
		if !left.has_value {
			stack = append(stack, m.left)
			continue
		}
		right := monkeys[m.right]
		if !right.has_value {
			stack = append(stack, m.right)
			continue
		}

		m.has_value = true
		switch m.op {
		case plus:
			m.value = left.value + right.value
		case devide:
			m.value = left.value / right.value
		case times:
			m.value = left.value * right.value
		case minues:
			m.value = left.value - right.value
		default:
			panic("Unknown operator")
		}
		stack = stack[:len(stack)-1]
	}
}

func branchHasYou(monkeys map[string]*monkey, root string, you string) bool {
	if root == you {
		return true
	}

	stack := []string{root}

	hasYou := false

	for len(stack) > 0 {
		next := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if next == you {
			hasYou = true
			break
		}
		m := monkeys[next]
		if m.has_value {
			continue
		}

		stack = append(stack, m.left, m.right)
	}

	return hasYou
}
