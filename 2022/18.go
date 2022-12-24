package aoc2022

import (
	"fmt"
	"strconv"
	"strings"
)

type point3d struct {
	x int
	y int
	z int
}

func (p_1 *point3d) Equals(p_2 point3d) bool {
	return p_1.x == p_2.x && p_1.y == p_2.y && p_1.z == p_2.z
}

func Day18(data string) {
	max_x := -1
	max_y := -1
	max_z := -1
	points := []point3d{}
	for _, p := range strings.Split(data, "\r\n") {
		coordinates := strings.Split(p, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		z, _ := strconv.Atoi(coordinates[2])
		points = append(points, point3d{x, y, z})
		if x > max_x {
			max_x = x
		}
		if y > max_y {
			max_y = y
		}
		if z > max_z {
			max_z = z
		}
	}

	unconnectedSides := calcUnconnectedSides(points)

	fmt.Println("Answer 1:", unconnectedSides)

	connectedGroups := map[int][]point3d{}
	index := 0
	for x := 0; x <= max_x; x++ {
		for y := 0; y <= max_y; y++ {
			for z := 0; z <= max_z; z++ {
				found := false
				for _, p := range points {
					if p.Equals(point3d{x, y, z}) {
						found = true
						break
					}
				}
				if !found {
					connectedGroups[index] = []point3d{{x, y, z}}
					index++
				}
			}
		}
	}

	mergedAGroup := true
	for mergedAGroup {
		mergedAGroup = false
		i_connected := -1
		j_connected := -1
		for i, group := range connectedGroups {
			if mergedAGroup {
				break
			}
			for j, group2 := range connectedGroups {
				if i == j || mergedAGroup {
					break
				}
				for _, p := range group {
					if mergedAGroup {
						break
					}
					for _, p2 := range group2 {
						if p.connects(p2) {
							mergedAGroup = true
							i_connected = i
							j_connected = j
							break
						}
					}
				}
			}
		}

		if mergedAGroup {
			connectedGroups[i_connected] = append(connectedGroups[i_connected], connectedGroups[j_connected]...)
			delete(connectedGroups, j_connected)
		}
	}

	/////////////// PRINT
	groupMap := map[int]int{}
	groupIndex := 1
	for i, _ := range connectedGroups {
		groupMap[i] = groupIndex
		groupIndex++
	}

	for x := 0; x <= max_x; x++ {
		fmt.Print("\n")
		for y := 0; y <= max_y; y++ {
			fmt.Print("\n")
			for z := 0; z <= max_z; z++ {
				found := false
				for i, g := range connectedGroups {
					for _, p := range g {
						if p.Equals(point3d{x, y, z}) {
							fmt.Print(groupMap[i], " ")
							found = true
						}
					}
				}
				if !found {
					fmt.Print(".. ")
				}
			}
		}
	}
	fmt.Print("\n")
	////////////

	for i, g := range connectedGroups {
		d := false
		for _, p := range g {
			if p.x == 0 || p.y == 0 || p.z == 0 || p.x == max_x || p.y == max_y || p.z == max_z {
				d = true
				break
			}
		}
		if d {
			delete(connectedGroups, i)
		}
	}

	for _, g := range connectedGroups {
		unconnectedSides -= calcUnconnectedSides(g)
	}

	fmt.Println("Answer 2:", unconnectedSides)
}

func calcUnconnectedSides(points []point3d) int {
	unconnectedSides := 0
	for _, p := range points {
		unconnected := 6
		for _, p2 := range points {
			if p.connects(p2) {
				unconnected--
			}
			if unconnected == 0 {
				break
			}
		}
		unconnectedSides += unconnected
	}
	return unconnectedSides
}

func (p *point3d) connects(p2 point3d) bool {
	return manDist3d(p.x, p.y, p.z, p2.x, p2.y, p2.z) == 1
}

func manDist3d(x_1 int, y_1 int, z_1 int, x_2 int, y_2 int, z_2 int) int {
	return abs(x_1-x_2) + abs(y_1-y_2) + abs(z_1-z_2)
}
