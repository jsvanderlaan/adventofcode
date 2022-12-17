package aoc2022

import "fmt"

const max_height = 2022*4 + 3 + 4
const chamber_width = 9

var b2i = map[bool]int{false: 1, true: -1}

type column [max_height]bool

type position struct {
	x int
	y int
}

func Day17(data string) {
	columns := [chamber_width]*column{{}, {}, {}, {}, {}, {}, {}, {}, {}}
	for i := 0; i < max_height; i++ {
		columns[0][i] = true
		columns[chamber_width-1][i] = true
	}
	for i := 0; i < chamber_width-1; i++ {
		columns[i][0] = true
	}

	rockCounter := -1
	rockIsFalling := false
	rockPosition := []position{}
	i := -1

	for rockCounter < 2022 {
		i++

		// print
		// if i < 50 {
		// 	fmt.Println("\nMovement", i, ". Rock count", rockCounter)
		// 	for x, col := range columns {
		// 		for y := 0; y < 30; y++ {
		// 			isRock := false
		// 			for _, r := range rockPosition {
		// 				if r.x == x && r.y == y {
		// 					isRock = true
		// 				}
		// 			}
		// 			if isRock {
		// 				fmt.Print("@")
		// 			} else if col[y] {
		// 				fmt.Print("#")
		// 			} else {
		// 				fmt.Print(".")
		// 			}
		// 		}
		// 		fmt.Print("\n")
		// 	}
		// }

		if !rockIsFalling {
			// create new rock
			rockIsFalling = true
			rockCounter++
			maxHeightColumn := maxHeightColumn(columns)
			switch rockCounter % 5 {
			case 0:
				rockPosition = []position{{3, maxHeightColumn + 4}, {4, maxHeightColumn + 4}, {5, maxHeightColumn + 4}, {6, maxHeightColumn + 4}}
			case 1:
				rockPosition = []position{{3, maxHeightColumn + 5}, {4, maxHeightColumn + 6}, {4, maxHeightColumn + 5}, {4, maxHeightColumn + 4}, {5, maxHeightColumn + 5}}
			case 2:
				rockPosition = []position{{3, maxHeightColumn + 4}, {4, maxHeightColumn + 4}, {5, maxHeightColumn + 4}, {5, maxHeightColumn + 5}, {5, maxHeightColumn + 6}}
			case 3:
				rockPosition = []position{{3, maxHeightColumn + 4}, {3, maxHeightColumn + 5}, {3, maxHeightColumn + 6}, {3, maxHeightColumn + 7}}
			case 4:
				rockPosition = []position{{3, maxHeightColumn + 4}, {4, maxHeightColumn + 4}, {3, maxHeightColumn + 5}, {4, maxHeightColumn + 5}}
			}
		}

		// move falling rock
		moveLeft := data[i%len(data)] == '<'
		canMove := true
		movedRockPosition := []position{}
		for _, pos := range rockPosition {
			newPos := position{pos.x + b2i[moveLeft], pos.y}
			if columns[newPos.x][newPos.y] {
				canMove = false
				break
			}
			movedRockPosition = append(movedRockPosition, newPos)
		}
		if canMove {
			rockPosition = movedRockPosition
		}

		// drop rock
		canDrop := true
		droppedRockPosition := []position{}
		for _, pos := range rockPosition {
			newPos := position{pos.x, pos.y - 1}
			if columns[newPos.x][newPos.y] {
				canDrop = false
				break
			}
			droppedRockPosition = append(droppedRockPosition, newPos)
		}
		if canDrop {
			rockPosition = droppedRockPosition
		} else {
			rockIsFalling = false
			for _, pos := range rockPosition {
				columns[pos.x][pos.y] = true
			}
		}
	}

	fmt.Println("Answer 1:", maxHeightColumn(columns))
}

func maxHeightColumn(columns [chamber_width]*column) int {
	max := 0
	for i := 1; i < chamber_width-1; i++ {
		for j := max_height - 1; j >= 0; j-- {
			if columns[i][j] && j > max {
				max = j
			}
		}
	}
	return max
}
