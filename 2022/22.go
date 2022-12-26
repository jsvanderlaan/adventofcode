package aoc2022

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	right dir = 0
	down  dir = 1
	left  dir = 2
	up    dir = 3
)

// 151012 too high
// 143396 too low

func Day22(data string) {
	boardString, commandString, _ := strings.Cut(data, "\r\n\r\n")
	board := [][]rune{}

	for _, row := range strings.Split(boardString, "\r\n") {
		board = append(board, []rune(row))
	}

	r, _ := regexp.Compile(`(\d{1,2}[RL]?)`)
	matches := r.FindAllString(commandString, -1)
	commands := []string{}
	for _, command := range matches {
		lastChar := command[len(command)-1]
		if lastChar == 'R' || lastChar == 'L' {
			commands = append(commands, command[:len(command)-1], string(lastChar))
		} else {
			commands = append(commands, command)
		}
	}

	direction := right
	pos := point{strings.Index(string(board[0]), "."), 0}

	// num := 0

	for _, command := range commands {
		// num++
		// if num > 20 {
		// 	break
		// }
		steps, err := strconv.Atoi(command)
		if err != nil {
			switch command {
			case "R":
				direction = (direction + 1) % 4
			case "L":
				direction = (direction - 1 + 4) % 4
			default:
				panic("Invalid command")
			}
			continue
		}
		next := pos
		next_dir := direction
		for i := 0; i < steps; i++ {
			next.x, next.y, next_dir = nextPos(board, next.x, next.y, next_dir)
			board[next.y][next.x] = getDirectionCharacter(next_dir)
		}
		pos = next
		direction = next_dir
		// fmt.Println(direction, pos)
	}

	for _, row := range board {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")

	fmt.Println("Answer 1:", 1000*(pos.y+1)+4*(pos.x+1)+int(direction))
}

// func findLastUsefulRow(board [][]rune, col int) int {
// 	for i := len(board) - 1; i >= 0; i-- {
// 		if len(board[i])-1 >= col {
// 			return i
// 		}
// 	}
// 	panic("No useful last row found")
// }

// func findFirstUsefulRow(board [][]rune, col int) int {
// 	for i := 0; i < len(board); i++ {
// 		if len(board[i])-1 >= col && board[i][col] != ' ' {
// 			return i
// 		}
// 	}
// 	panic("No useful first row found")
// }

// func findFirstUsefulIndex(str []rune) int {
// 	for i := 0; i < len(str); i++ {
// 		if str[i] != ' ' {
// 			return i
// 		}
// 	}
// 	panic("No useful first col found")
// }

func getDirectionCharacter(d dir) rune {
	switch d {
	case right:
		return '>'
	case up:
		return '^'
	case left:
		return '<'
	case down:
		return 'v'
	}
	panic("Direction unknown")
}

func findSide(x int, y int) int {
	if x >= 0 && x < 50 {
		if y >= 100 && y < 150 {
			return 4
		}
		if y >= 150 && y < 200 {
			return 6
		}
	}
	if x >= 50 && x < 100 {
		if y >= 0 && y < 50 {
			return 1
		}
		if y >= 50 && y < 100 {
			return 3
		}
		if y >= 100 && y < 150 {
			return 5
		}
	}

	if x >= 100 && x < 150 {
		if y >= 0 && y < 50 {
			return 2
		}
	}
	return -1
}

func nextPos(board [][]rune, x int, y int, d dir) (int, int, dir) {
	side := findSide(x, y)
	x_next, y_next, dir_next := x, y, d
	switch d {
	case right:
		new_x := x + 1
		new_y := y
		new_d := d
		if new_x >= len(board[y]) {
			switch side {
			case 2:
				new_x = 99
				new_y = 149 - y
				new_d = left
			case 3:
				new_x = 100 + (y - 50)
				new_y = 49
				new_d = up
			case 5:
				new_x = 149
				new_y = 49 - (y - 100)
				new_d = left
			case 6:
				new_x = 50 + (y - 150)
				new_y = 149
				new_d = up
			default:
				panic("OEPS")
			}
		}
		if board[new_y][new_x] != '#' {
			x_next = new_x
			y_next = new_y
			dir_next = new_d
		}
	case left:
		new_x := x - 1
		new_y := y
		new_d := d
		if new_x < 0 || board[y][new_x] == ' ' {
			switch side {
			case 1:
				new_x = 0
				new_y = 149 - y
				new_d = right
			case 3:
				new_x = (y - 50)
				new_y = 100
				new_d = down
			case 4:
				new_x = 50
				new_y = 49 - (y - 100)
				new_d = right
			case 6:
				new_x = 50 + (y - 150)
				new_y = 0
				new_d = down
			default:
				panic("OEPS")
			}
		}
		if board[new_y][new_x] != '#' {
			x_next = new_x
			y_next = new_y
			dir_next = new_d
		}
	case up:
		new_x := x
		new_y := y - 1
		new_d := d
		if new_y < 0 || board[new_y][x] == ' ' {
			switch side {
			case 4:
				new_x = 50
				new_y = 50 + x
				new_d = right // 3
			case 1:
				new_x = 0
				new_y = 150 + (x - 50)
				new_d = right // 6
			case 2:
				new_x = x - 100
				new_y = 199
				new_d = up // 6
			default:
				panic("OEPS")
			}
		}
		if board[new_y][new_x] != '#' {
			x_next = new_x
			y_next = new_y
			dir_next = new_d
		}
	case down:
		new_x := x
		new_y := y + 1
		new_d := d
		if new_y >= len(board) || x >= len(board[new_y]) || board[new_y][x] == ' ' {
			switch side {
			case 6:
				new_x = 100 + x
				new_y = 0
				new_d = down // 2
			case 5:
				new_x = 49
				new_y = 150 + (x - 50)
				new_d = left // 6
			case 2:
				new_x = 99
				new_y = 50 + (x - 100)
				new_d = left // 3
			default:
				panic("OEPS")
			}
		}
		if board[new_y][new_x] != '#' {
			x_next = new_x
			y_next = new_y
			dir_next = new_d
		}
	}
	return x_next, y_next, dir_next
}
