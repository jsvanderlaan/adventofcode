package aoc2022

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	id           int
	items        []int
	operation    string
	testNumber   int
	succeedId    int
	failId       int
	inspectCount int
}

func (m *Monkey) inspect(item int, prod int) int {
	m.inspectCount++
	// item := m.items[len(m.items)-1]
	// m.items = m.items[:len(m.items)-1]
	op := strings.ReplaceAll(m.operation, "old", strconv.Itoa(item))
	ops := strings.Split(op, " ")
	left, _ := strconv.Atoi(ops[0])
	right, _ := strconv.Atoi(ops[2])
	var new int

	if ops[1] == "+" {
		new = left + right
	} else {
		new = left * right
	}

	return new % prod // / 3
}

func (m *Monkey) test(item int) bool {
	//m.items[len(m.items)-1]
	return item%m.testNumber == 0
}

func Day11(data string) {
	monkeysData := strings.Split(data, "\r\n\r\n")

	monkeys := map[int]*Monkey{}

	// parse
	for i, m := range monkeysData {
		lines := strings.Split(m, "\r\n")
		items := []int{}
		for _, item := range strings.Split(lines[1][18:], ", ") {
			i, _ := strconv.Atoi(item)
			items = append(items, i)
		}
		operation := lines[2][19:]
		testNumber, _ := strconv.Atoi(lines[3][21:])
		succeedId, _ := strconv.Atoi(lines[4][29:])
		failedId, _ := strconv.Atoi(lines[5][30:])

		monkeys[i] = &Monkey{i, items, operation, testNumber, succeedId, failedId, 0}
	}

	prod := 1
	for _, m := range monkeys {
		prod *= m.testNumber
	}

	// rounds
	for i := 0; i < 10000; i++ {
		if i%1000 == 0 {
			fmt.Println("Round", i)
			for _, m := range monkeys {
				fmt.Println(m)
			}
		}
		for j := 0; j < len(monkeys); j++ {
			m := monkeys[j]
			for len(m.items) > 0 {
				item := m.items[0]
				// fmt.Println("Monkey", j, "with item", item)
				m.items = m.items[1:]
				item = m.inspect(item, prod)
				// fmt.Println("After inspect", item)
				test := m.test(item)
				// fmt.Println("Test", test)
				id := -1
				if test {
					id = m.succeedId
				} else {
					id = m.failId
				}
				// fmt.Println("Id", id)
				monkeys[id].items = append(monkeys[id].items, item)
			}
		}
	}

	inspectTimes := []int{}
	for _, m := range monkeys {
		inspectTimes = append(inspectTimes, m.inspectCount)
	}

	sort.Ints(inspectTimes)

	fmt.Println("Answer 1:", inspectTimes[len(inspectTimes)-2]*inspectTimes[len(inspectTimes)-1])
}
