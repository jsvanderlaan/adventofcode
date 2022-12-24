package aoc2022

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type file struct {
	value int
	left  *file
	right *file
}

func Day20(data string) {
	initialFiles := []*file{}
	var zero_file *file
	var first_file *file
	var last_file *file
	// fileMap := map[int]*file{}
	for i, f := range strings.Split(data, "\r\n") {
		file_value, _ := strconv.Atoi(f)
		new := file{file_value, nil, nil}
		initialFiles = append(initialFiles, &new)

		if file_value == 0 {
			zero_file = &new
		}

		if i == 0 {
			first_file = &new
		}

		var left *file
		if i > 0 {
			left = last_file
			left.right = &new
			new.left = left
		}

		last_file = &new
	}

	// let file wrap

	first_file.left = last_file
	last_file.right = first_file

	for _, f := range initialFiles {
		// fmt.Println("Before moving", f)
		// printFiles(fileMap)

		if f.value == 0 {
			continue
		}

		// cut file out
		f.left.right = f.right
		f.right.left = f.left

		// fmt.Println("After cutting out", f, f.left, f.right)
		// printFiles(fileMap)

		number_of_steps := int(math.Abs(float64(f.value)))
		next_pos := f
		// fmt.Println("next_pos", next_pos)
		go_left := f.value < 0
		for i := 0; i < number_of_steps; i++ {
			if go_left {
				next_pos = next_pos.left
			} else {
				next_pos = next_pos.right
			}
			// if next_pos.value == 698 {
			// 	fmt.Println("next_pos", next_pos, next_pos.right)
			// }
		}

		if go_left { // next_pos is new right of file
			f.right = next_pos
			f.left = next_pos.left
			next_pos.left.right = f
			next_pos.left = f
		} else { // next_pos is new left oof file
			f.left = next_pos
			f.right = next_pos.right
			next_pos.right.left = f
			next_pos.right = f
		}

		// fmt.Println("After pasting", f, "where next_pos was", next_pos, "pasted file is", f)
		// printFiles(fileMap)
	}

	sum := 0
	pos := zero_file
	for i := 0; i < 3000; i++ {
		pos = pos.right
		if (i+1)%1000 == 0 {
			fmt.Println(i+1, pos.value)
			sum += pos.value
		}
	}
	fmt.Println("Answer 1:", sum)
}

func printFiles(fileMap map[int]*file) {
	start_file := fileMap[0]
	for i := 0; i < len(fileMap); i++ {
		fmt.Print(start_file.value, ", ")
		start_file = start_file.right
	}
	fmt.Print("\n")
}
