package aoc2022

import (
	"fmt"
	"strings"
)

func Day3(data string) {
	rucksacks := strings.Split(data, "\r\n")
	items := []byte{}
	for _, sack := range rucksacks {
		l := len(sack)
		found := make(map[string]struct{})
		for i := 0; i < l/2; i++ {
			found[sack[i:i+1]] = struct{}{}
		}
		for i := l / 2; i < len(sack); i++ {
			item := sack[i : i+1]
			if _, a := found[item]; a {
				items = append(items, toValue(item))
				break
			}
		}
	}

	sum := 0
	for _, item := range items {
		sum += int(item)
	}

	fmt.Println("Answer 1:", sum)

	badges := []byte{}
	dedupped := []string{}

	for _, sack := range rucksacks {
		dedupped = append(dedupped, removeDuplicate(sack))
	}

	numGroups := len(dedupped) / 3

	for i := 0; i < numGroups; i++ {
		found := make(map[string]byte)
		for j := 0; j < 3; j++ {
			sack := dedupped[i*3+j]
			for i := 0; i < len(sack); i++ {
				found[sack[i:i+1]] += 1
			}
		}

		for k, v := range found {
			if v == 3 {
				badges = append(badges, toValue(k))
				break
			}
		}
	}

	sum = 0
	for _, item := range badges {
		sum += int(item)
	}

	fmt.Println("Answer 2:", sum)
}

func toValue(c string) byte {
	b := []byte(c)[0]
	if b >= 97 { // a = 97, A = 65
		return b - 96
	} else {
		return b - 65 + 27
	}
}

func removeDuplicate(sliceList string) string {
	allKeys := make(map[rune]bool)
	list := ""
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list += string(item)
		}
	}
	return list
}
