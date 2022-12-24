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
	const decryption_key = 811589153
	const number_of_mixes = 10

	initialFiles := []*file{}
	var zero_file *file
	var first_file *file
	var last_file *file

	for i, f := range strings.Split(data, "\r\n") {
		file_value, _ := strconv.Atoi(f)
		new := file{file_value * decryption_key, nil, nil}
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

	// let files wrap
	first_file.left = last_file
	last_file.right = first_file

	for j := 0; j < number_of_mixes; j++ {
		for _, f := range initialFiles {

			if f.value == 0 {
				continue
			}

			// cut file out
			f.left.right = f.right
			f.right.left = f.left

			number_of_steps := int(math.Abs(float64(f.value))) % (len(initialFiles) - 1)
			next_pos := f

			go_left := f.value < 0
			for i := 0; i < number_of_steps; i++ {
				if go_left {
					next_pos = next_pos.left
				} else {
					next_pos = next_pos.right
				}
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
		}
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
