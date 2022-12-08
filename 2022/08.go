package aoc2022

import (
	"fmt"
	"strconv"
	"strings"
)

func Day8(data string) {
	rows := strings.Split(data, "\r\n")
	size := len(rows)

	visible := map[int]struct{}{}

	for i := 0; i < size; i++ {
		max := -1
		for j := 0; j < size; j++ {
			if height, _ := strconv.Atoi(string(rows[i][j])); height > max {
				visible[toId(i, j, size)] = struct{}{}
				max = height
			}
		}
	}

	for i := 0; i < size; i++ {
		max := -1
		for j := size - 1; j >= 0; j-- {
			if height, _ := strconv.Atoi(string(rows[i][j])); height > max {
				visible[toId(i, j, size)] = struct{}{}
				max = height
			}
		}
	}

	for j := 0; j < size; j++ {
		max := -1
		for i := 0; i < size; i++ {
			if height, _ := strconv.Atoi(string(rows[i][j])); height > max {
				visible[toId(i, j, size)] = struct{}{}
				max = height
			}
		}
	}

	for j := 0; j < size; j++ {
		max := -1
		for i := size - 1; i >= 0; i-- {
			if height, _ := strconv.Atoi(string(rows[i][j])); height > max {
				visible[toId(i, j, size)] = struct{}{}
				max = height
			}
		}
	}

	fmt.Println("Answer 1:", len(visible))

	heights := [100][100]int{}
	score := map[int]int{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			height, _ := strconv.Atoi(string(rows[i][j]))
			heights[i][j] = height
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			height := heights[i][j]
			down := 0
			for x := i + 1; x < size; x++ {
				down++
				if height <= heights[x][j] {
					break
				}
			}
			up := 0
			for x := i - 1; x >= 0; x-- {
				up++
				if height <= heights[x][j] {
					break
				}
			}
			right := 0
			for x := j + 1; x < size; x++ {
				right++
				if height <= heights[i][x] {
					break
				}
			}
			left := 0
			for x := j - 1; x >= 0; x-- {
				left++
				if height <= heights[i][x] {
					break
				}
			}

			score[toId(i, j, size)] = up * down * right * left

		}
	}

	max := 0
	for _, s := range score {
		if s > max {
			max = s
		}
	}

	fmt.Println("Answer 2:", max)

}

func toId(row int, col int, size int) int {
	return row*size + col
}
