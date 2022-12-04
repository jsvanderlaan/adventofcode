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
