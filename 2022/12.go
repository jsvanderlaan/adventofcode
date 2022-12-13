package aoc2022

import (
	"fmt"
	"strings"
)

type dir byte

const (
	nodir dir = 0
	east  dir = 1
	south dir = 2
	west  dir = 3
	north dir = 4
)

type place struct {
	x     int
	y     int
	steps int
	dir   dir
}

func Day12(data string) {
	rows := strings.Split(data, "\r\n")
	height := len(rows)
	width := len(rows[0])
	x_S := -1
	y_S := -1
	x_E := -1
	y_E := -1

	for y, row := range rows {
		for x, cell := range row {
			if cell == 'S' {
				x_S = x
				y_S = y
			}
			if cell == 'E' {
				x_E = x
				y_E = y
			}
		}
	}

	visited := map[string]struct{}{}
	minSteps := map[string]int{}
	stack := []*place{{x_S, y_S, 0, nodir}}

	for len(stack) > 0 {
		currPlace := stack[len(stack)-1]
		key := fmt.Sprintf("%d|%d", currPlace.x, currPlace.y)
		visited[key] = struct{}{}
		minSteps[key] = currPlace.steps

		// fmt.Println("currPlace", currPlace)

		if currPlace.x == x_E && currPlace.y == y_E {
			// fmt.Println("\tpop! This is the end with stack")
			// for _, s := range stack {
			// fmt.Println("\t\t", s)
			// }
			// fmt.Println("\t\t", minSteps)

			stack = stack[:len(stack)-1]
			delete(visited, key)

			continue
		}

		x_next := -1
		y_next := -1
		if currPlace.dir == nodir {
			// fmt.Println("\tnext east")
			currPlace.dir = east
			x_next = currPlace.x + 1
			y_next = currPlace.y
		} else if currPlace.dir == east {
			// fmt.Println("\tnext south")
			currPlace.dir = south
			x_next = currPlace.x
			y_next = currPlace.y + 1
		} else if currPlace.dir == south {
			// fmt.Println("\tnext west")
			currPlace.dir = west
			x_next = currPlace.x - 1
			y_next = currPlace.y
		} else if currPlace.dir == west {
			// fmt.Println("\tnext north")
			currPlace.dir = north
			x_next = currPlace.x
			y_next = currPlace.y - 1
		} else {
			// fmt.Println("\tpop!")
			stack = stack[:len(stack)-1]
			delete(visited, key)
			continue
		}
		// fmt.Println("\tnext", currPlace, x_next, y_next)

		if x_next >= 0 && x_next < width && y_next >= 0 && y_next < height {
			// fmt.Println("\t\twithin map", currPlace, x_next, y_next)

			key_next := fmt.Sprintf("%d|%d", x_next, y_next)
			_, visited_next := visited[key_next]
			if !visited_next {
				elevation_curr := rune(rows[currPlace.y][currPlace.x])
				if elevation_curr == 'S' {
					elevation_curr = 'a'
				}
				elevation_next := rune(rows[y_next][x_next])
				if elevation_next == 'E' {
					elevation_next = 'z'
				}
				diff := int(elevation_next) - int(elevation_curr)
				// fmt.Println("\t\tnot visited", elevation_curr, elevation_next, diff)
				if diff <= 1 {
					// fmt.Println("\t\tE or diff <= 1", elevation_curr, elevation_next, diff)
					min, ok := minSteps[key_next]
					if !ok || min > currPlace.steps+1 {
						stack = append(stack, &place{x_next, y_next, currPlace.steps + 1, nodir})
					}
				}
			}
		}

	}

	fmt.Println("Answer 1:", minSteps[fmt.Sprintf("%d|%d", x_E, y_E)])
}
