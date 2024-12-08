package aoc2024

import (
	"fmt"
	"strings"
)

func Day4(data string) {
	// fmt.Println(data)
	letters := [][]rune{}
	rows := strings.Split(data, "\r\n")
	for _, row := range rows {
		letters = append(letters, []rune(row))
	}

	directions := // [dx, dy]
		[][]int{
			{0, 1},
			{-1, 1},
			{-1, 0},
			{-1, -1},
			{0, -1},
			{1, -1},
			{1, 0},
			{1, 1},
		}

	normalizedRows := []string{}
	for _, direction := range directions {
		dx := direction[0]
		dy := direction[1]
		startingPositions := map[int]map[int]bool{}
		for i := 0; i < len(letters); i++ {
			startingPositions[i] = map[int]bool{}
		}

		if dx == 1 {
			for i := 0; i < len(letters); i++ {
				startingPositions[i][0] = true
			}
		}

		if dx == -1 {
			for i := 0; i < len(letters); i++ {
				startingPositions[i][len(letters[0])-1] = true
			}
		}

		if dy == 1 {
			for i := 0; i < len(letters[0]); i++ {
				startingPositions[0][i] = true
			}
		}

		if dy == -1 {
			for i := 0; i < len(letters[0]); i++ {
				startingPositions[len(letters)-1][i] = true
			}
		}

		// fmt.Printf("%v\n", startingPositions)
		for y := 0; y < len(letters); y++ {
			for x := 0; x < len(letters[0]); x++ {
				if startingPositions[y][x] {
					i := x
					j := y
					word := ""
					for i >= 0 && i < len(letters[0]) && j >= 0 && j < len(letters) {
						word += string(letters[j][i])
						j += dy
						i += dx
					}
					normalizedRows = append(normalizedRows, word)
				}
			}
		}
	}
	// fmt.Printf("%v\n", normalizedRows)
	// fmt.Printf("%v\n", len(normalizedRows))

	count := 0
	for _, word := range normalizedRows {
		for i := 0; i < len(word)-3; i++ {
			if word[i] == 'X' && word[i:i+4] == "XMAS" {
				count++
				i += 3
			}
		}
	}

	fmt.Printf("count1 %v\n", count)
	count = 0
	for y := 1; y < len(letters)-1; y++ {
		for x := 1; x < len(letters[0])-1; x++ {
			if letters[y][x] == 'A' &&
				((letters[y-1][x-1] == 'M' && letters[y+1][x-1] == 'M' && letters[y+1][x+1] == 'S' && letters[y-1][x+1] == 'S') || (letters[y-1][x-1] == 'S' && letters[y+1][x-1] == 'M' && letters[y+1][x+1] == 'M' && letters[y-1][x+1] == 'S') || (letters[y-1][x-1] == 'S' && letters[y+1][x-1] == 'S' && letters[y+1][x+1] == 'M' && letters[y-1][x+1] == 'M') || (letters[y-1][x-1] == 'M' && letters[y+1][x-1] == 'S' && letters[y+1][x+1] == 'S' && letters[y-1][x+1] == 'M')) {
				count++
			}
		}
	}
	fmt.Printf("count2 %v\n", count)

}
