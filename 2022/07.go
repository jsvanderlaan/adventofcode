package aoc2022

import (
	"fmt"
	"strconv"
	"strings"
)

type Dir struct {
	name   string
	size   int
	childs []int
	isDir  bool
	parent int
}

func Day7(data string) {
	output := strings.Split(data, "\r\n")
	dirs := map[int]*Dir{}
	counter := 0
	currDir := 0
	dirs[0] = &Dir{"/", 0, []int{}, true, 0}

	for _, o := range output {
		if strings.HasPrefix(o, "$ cd") {
			d := strings.Split(o, " ")[2]
			if d == ".." {
				currDir = dirs[currDir].parent
			} else {
				for _, c := range dirs[currDir].childs {
					if dirs[c].name == d {
						currDir = c
					}
				}
			}
			continue
		}

		if strings.HasPrefix(o, "$ ls") {
			continue
		}

		if strings.HasPrefix(o, "dir ") {
			counter++
			x := strings.Split(o, " ")[1]
			dirs[currDir].childs = append(dirs[currDir].childs, counter)
			dirs[counter] = &Dir{x, 0, []int{}, true, currDir}
			continue
		}

		// file
		counter++
		x := strings.Split(o, " ")
		size, _ := strconv.Atoi(x[0])
		dirs[currDir].childs = append(dirs[currDir].childs, counter)
		dirs[counter] = &Dir{x[1], size, []int{}, false, currDir}
	}

	setSize(0, dirs)

	sum := 0
	for _, d := range dirs {
		if d.size < 100_000 && d.isDir {
			sum += d.size
		}
	}

	// for k, d := range dirs {
	// 	fmt.Println(k, d)
	// }

	fmt.Println("Answer 1:", sum)

	totalSpace := 70000000
	totalUsed := dirs[0].size
	totalFree := totalSpace - int(totalUsed)
	chosenSize := 70000000

	for _, d := range dirs {
		if d.isDir && d.size <= chosenSize && d.size+totalFree >= 30000000 {
			chosenSize = d.size
		}
	}

	fmt.Println("Answer 2:", chosenSize)
}

func setSize(d int, dirs map[int]*Dir) int {
	dir := dirs[d]
	if !dir.isDir {
		return dir.size
	}

	sum := 0
	for _, c := range dir.childs {
		sum += setSize(c, dirs)
	}
	dir.size = sum
	return sum
}
