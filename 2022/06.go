package aoc2022

import "fmt"

func Day6(data string) {
	markerPos := -1
	i := 0
	for markerPos == -1 {
		if checkAllDifferent([]byte(data[i : i+14])) {
			markerPos = i
		}
		i++
	}
	fmt.Println("Answer 1:", markerPos+14)
}

func checkAllDifferent(chars []byte) bool {
	m := make(map[byte]struct{})
	for _, b := range chars {
		m[b] = struct{}{}
	}

	return len(chars) == len(m)
}
