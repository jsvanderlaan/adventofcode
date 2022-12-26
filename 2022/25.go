package aoc2022

import (
	"fmt"
	"strconv"
	"strings"
)

type snafu string
type decimal int64

func (number *snafu) ToDecimal() decimal {
	runes := []rune(*number)

	extra := 0
	for i := len(runes) - 1; i >= 0; i-- {
		var value int
		switch runes[i] {
		case '=':
			value = -2
		case '-':
			value = -1
		default:
			value = int(runes[i] - '0')
		}

		actual := value - extra

		runes[i] = rune('0' + (actual+5)%5)

		if actual < 0 {
			extra = 1
		} else {
			extra = 0
		}
	}
	quinary, _ := strconv.ParseInt(string(runes), 5, 64)
	return decimal(quinary)
}

func (number *decimal) ToSnafu() snafu {
	quinary := strconv.FormatInt(int64(*number), 5)
	quinaryStr := []rune(quinary)

	extra := 0
	for i := len(quinaryStr) - 1; i >= 0; i-- {
		value := int(quinaryStr[i] - rune('0'))
		actual := value + extra

		if actual < 3 {
			quinaryStr[i] = rune('0') + rune(actual)
			extra = 0
			continue
		}

		extra = 1

		switch actual {
		case 3:
			quinaryStr[i] = '='
		case 4:
			quinaryStr[i] = '-'
		case 5:
			quinaryStr[i] = '0'
		}

		if i == 0 && extra == 1 {
			quinaryStr = append([]rune{'1'}, quinaryStr...)
		}
	}

	return snafu(quinaryStr)
}

func Day25(data string) {
	values := strings.Split(data, "\r\n")
	sum := int64(0)
	for _, v := range values {
		x := snafu(v)
		sum += int64(x.ToDecimal())
	}
	decimal := decimal(sum)
	fmt.Println(decimal.ToSnafu())
}
