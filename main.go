package main

import (
	aoc2022 "aoc/2022"
	"aoc/utils"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	year := args[0]
	day := args[1]
	test := ""
	if len(args) > 2 {
		test = args[2]
	}
	data := readData(year, day, test)

	var f func(string)
	switch year {
	case "2022":
		switch day {
		case "01":
			f = aoc2022.Day1
		case "02":
			f = aoc2022.Day2
		case "03":
			f = aoc2022.Day3
		case "04":
			f = aoc2022.Day4
		case "05":
			f = aoc2022.Day5
		case "06":
			f = aoc2022.Day6
		case "07":
			f = aoc2022.Day7
		case "08":
			f = aoc2022.Day8
		case "09":
			f = aoc2022.Day9
		case "10":
			f = aoc2022.Day10
		case "11":
			f = aoc2022.Day11
		case "12":
			f = aoc2022.Day12
		}
	}

	f(data)
}

func readData(year string, day string, test string) string {
	t := ""
	if test == "t" {
		t = ".t"
	}
	file := fmt.Sprintf("./%s/data/%s%s.data", year, day, t)
	data, err := os.ReadFile(file)
	utils.Check(err)
	return string(data)
}
