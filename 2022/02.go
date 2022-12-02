package aoc2022

import (
	"fmt"
	"strings"
)

type choice int
type dual struct {
	you  choice
	them choice
}

const (
	Rock     choice = 1
	Paper    choice = 2
	Scissors choice = 3
)

type outcome int

const (
	loose outcome = 1
	draw  outcome = 2
	win   outcome = 3
)

func Day2(data string) {
	rows := strings.Split(data, "\r\n")
	duals := []dual{}
	for _, row := range rows {
		dual := dual{them: stoc(row[0:1]), you: stoc(row[2:3])}
		duals = append(duals, dual)
	}
	score := 0
	for _, dual := range duals {
		score += int(dual.you)
		if dual.you == dual.them {
			score += 3
		} else if (dual.you == Rock && dual.them == Scissors) || (dual.you == Paper && dual.them == Rock) || (dual.you == Scissors && dual.them == Paper) {
			score += 6
		}
	}

	fmt.Println("Answer 1", score)

	score = 0
	for _, dual := range duals {
		outc := outcome(dual.you)
		if outc == win {
			score += 6
			switch dual.them {
			case Rock: // Paper to win
				score += int(Paper)
			case Paper: // Scissors to win
				score += int(Scissors)
			case Scissors: // Rock to win
				score += int(Rock)
			}
		} else if outc == draw {
			score += 3
			score += int(dual.them)
		} else {
			switch dual.them {
			case Rock: // Scissors to loose
				score += int(Scissors)
			case Paper: // Rock to loose
				score += int(Rock)
			case Scissors: // Paper to loose
				score += int(Paper)
			}
		}
	}

	fmt.Println("Answer 2", score)
}

func stoc(str string) choice {
	switch str {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}
	panic("KAPOT")
}
