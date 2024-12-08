package aoc2024

import (
	"fmt"
	"strconv"
	"unicode"
)

type ServerState int

func Day3(data string) {
	sum := 0
	do := true

	for i := 0; i < len(data)-7; i++ {
		// fmt.Printf("data[i] %v\n", string(data[i]))
		if !do && data[i:i+4] == "do()" {
			// fmt.Printf("do()\n")
			do = true
			i += 3
			continue
		}
		if do && data[i:i+7] == "don't()" {
			// fmt.Printf("don't()\n")
			do = false
			i += 6
			continue
		}

		if do && data[i] == 'm' && data[i:i+4] == "mul(" {
			num1 := ""
			for j := 0; j < 3; j++ {
				next := rune(data[i+4+j])
				if unicode.IsDigit(next) {
					num1 += string(next)
				} else {
					break
				}
			}
			if len(num1) == 0 || data[i+4+len(num1)] != ',' {
				continue
			}
			num2 := ""
			for j := 0; j < 3; j++ {
				next := rune(data[i+4+len(num1)+1+j])
				if unicode.IsDigit(next) {
					num2 += string(next)
				} else {
					break
				}
			}
			if len(num2) == 0 || data[i+4+len(num1)+1+len(num2)] != ')' {
				continue
			}
			n1, _ := strconv.Atoi(num1)
			n2, _ := strconv.Atoi(num2)
			sum += n1 * n2
			i += 4 + len(num1) + 1 + len(num2)
			// fmt.Printf("mul(%v,%v)\n", num1, num2)
		}

	}

	fmt.Println(sum)
}
