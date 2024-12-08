package aoc2024

import (
	"fmt"
	"strings"
)

func Next(dir [2]int) [2]int {
	if dir[0] == 1 {
		return [2]int{0, 1}
	} else if dir[0] == -1 {
		return [2]int{0, -1}
	} else if dir[1] == 1 {
		return [2]int{-1, 0}
	} else if dir[1] == -1 {
		return [2]int{1, 0}
	}
	panic("invalid dir")
}

func DirToInt(dir [2]int) int {
	return (dir[0] + dir[1]*2)
}

func Day6(data string) {
	visited := map[int]map[int]bool{}
	obstacles := map[int]map[int]bool{}
	start := [2]int{0, 0}
	startDir := [2]int{0, -1}

	rows := strings.Split(data, "\r\n")
	height := len(rows)
	width := len(rows[0])

	for y, row := range rows {
		visited[y] = map[int]bool{}
		obstacles[y] = map[int]bool{}
		for x, cell := range row {
			if cell == '#' {
				obstacles[y][x] = true
			}
			if cell == '^' {
				visited[y][x] = true
				start = [2]int{x, y}
			}
		}
	}

	// fmt.Println(obstacles)
	current := [2]int{start[0], start[1]}
	dir := [2]int{startDir[0], startDir[1]}

	for current[0] >= 0 && current[0] < width && current[1] >= 0 && current[1] < height {
		// fmt.Println(current)
		x := current[0]
		y := current[1]
		visited[y][x] = true

		next := [2]int{x + dir[0], y + dir[1]}
		if _, ok := obstacles[next[1]][next[0]]; ok {
			dir = Next(dir)
		}
		current = [2]int{x + dir[0], y + dir[1]}
	}

	sum := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if _, ok := visited[y][x]; ok {
				sum++
			}
		}
	}

	fmt.Println(sum)

	// part 2
	// 1599 too low

	sum = 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if obstacle, ok := obstacles[y][x]; ok && obstacle {
				continue
			}

			if x == start[0] && y == start[1] {
				continue
			}

			if _, ok := visited[y][x]; !ok {
				continue
			}

			cp := CopyMap(obstacles)
			cp[y][x] = true

			// fmt.Println(cp)

			if ContainsLoop(map[int]map[int]int{}, cp, height, width, [2]int{start[0], start[1]}, [2]int{startDir[0], startDir[1]}) {
				sum++
			}

		}
	}

	fmt.Println(sum)
}

func ContainsLoop(
	visitedDirs map[int]map[int]int,
	obstacles map[int]map[int]bool,
	height int,
	width int,
	current [2]int,
	dir [2]int) bool {

	for current[0] >= 0 && current[0] < width && current[1] >= 0 && current[1] < height {
		// fmt.Println(current)
		x := current[0]
		y := current[1]

		if _, ok := visitedDirs[y]; ok {
			if count, ok := visitedDirs[y][x]; ok {
				if count > 3 {
					return true
				} else {
					visitedDirs[y][x] += 1
				}
			} else {
				visitedDirs[y][x] = 1
			}
		} else {
			visitedDirs[y] = map[int]int{}
			visitedDirs[y][x] = 1
		}

		obstacle := true
		for obstacle {
			next := [2]int{x + dir[0], y + dir[1]}
			if _, ok := obstacles[next[1]][next[0]]; ok {
				dir = Next(dir)
			} else {
				obstacle = false
			}

		}

		current = [2]int{x + dir[0], y + dir[1]}
	}
	return false

}

func CopyMap(m map[int]map[int]bool) map[int]map[int]bool {
	cp := make(map[int]map[int]bool)
	for i, v := range m {
		cp[i] = map[int]bool{}
		for j, k := range v {
			cp[i][j] = k
		}
	}

	return cp
}
