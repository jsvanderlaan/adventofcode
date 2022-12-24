package aoc2022

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	none     choice = 0
	ore      choice = 1
	clay     choice = 2
	obsidian choice = 3
	geode    choice = 4
)

type plan []choice

type blueprint struct {
	id             int
	ore_ore        int
	clay_ore       int
	obsidian_ore   int
	obsidian_clay  int
	geode_ore      int
	geode_obsidian int
}

func Day19(data string) {
	const minutes int = 24
	blueprints := []*blueprint{}
	for i, b := range strings.Split(data, "\r\n") {

		r, _ := regexp.Compile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)
		matches := r.FindAllSubmatch([]byte(b), -1)
		ore_ore, _ := strconv.Atoi(string(matches[0][2]))
		clay_ore, _ := strconv.Atoi(string(matches[0][3]))
		obsidian_ore, _ := strconv.Atoi(string(matches[0][4]))
		obsidian_clay, _ := strconv.Atoi(string(matches[0][5]))
		geode_ore, _ := strconv.Atoi(string(matches[0][6]))
		geode_obsidian, _ := strconv.Atoi(string(matches[0][7]))

		blueprints = append(blueprints, &blueprint{i + 1, ore_ore, clay_ore, obsidian_ore, obsidian_clay, geode_ore, geode_obsidian})
	}

	quality_level := 0

	for _, b := range blueprints {
		max_ore := b.ore_ore
		if b.clay_ore > max_ore {
			max_ore = b.clay_ore
		}
		if b.obsidian_ore > max_ore {
			max_ore = b.obsidian_ore
		}
		if b.geode_ore > max_ore {
			max_ore = b.geode_ore
		}
		max_clay := b.obsidian_clay
		max_obsidian := b.geode_obsidian

		max_geodes := -1
		queue := []plan{{}}

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			geodes, executed_plan := geodes_found(b, p, minutes, false)

			// fmt.Println("Blueprint", b.id, "with plan", p, "has geodes", geodes, "and executed plan", executed_plan, "resources", ore, clay, obsidian, geode)

			if len(p) <= len(executed_plan) {
				if max_geodes < geodes {
					max_geodes = geodes
				}
				// for i := ore; i <= geode; i++ {
				// 	queue = append(queue, append(plan, i))
				// }
				ore_count := 0
				clay_count := 0
				obsidian_count := 0
				for _, c := range p {
					if c == ore {
						ore_count++
					} else if c == clay {
						clay_count++
					} else if c == obsidian {
						obsidian_count++
					}
				}

				if ore_count < max_ore {
					ore_plan := make(plan, len(p))
					copy(ore_plan, p)
					ore_plan = append(ore_plan, ore)
					queue = append(queue, ore_plan)
				}
				if clay_count < max_clay {
					clay_plan := make(plan, len(p))
					copy(clay_plan, p)
					clay_plan = append(clay_plan, clay)
					queue = append(queue, clay_plan)
				}
				if obsidian_count < max_obsidian {
					obsidian_plan := make(plan, len(p))
					copy(obsidian_plan, p)
					obsidian_plan = append(obsidian_plan, obsidian)
					queue = append(queue, obsidian_plan)
				}
				geode_plan := make(plan, len(p))
				copy(geode_plan, p)
				geode_plan = append(geode_plan, geode)
				queue = append(queue, geode_plan)

				// clay_plan := make(plan, len(p))
				// copy(clay_plan, p)
				// clay_plan = append(clay_plan, clay)
				// obsidian_plan := make(plan, len(p))
				// copy(obsidian_plan, p)
				// obsidian_plan = append(obsidian_plan, obsidian)
				// geode_plan := make(plan, len(p))
				// copy(geode_plan, p)
				// geode_plan = append(geode_plan, geode)
				// // fmt.Println(ore_plan, clay_plan, obsidian_plan, geode_plan)
				// queue = append(queue,
				// 	ore_plan,
				// 	clay_plan,
				// 	obsidian_plan,
				// 	geode_plan,
				// )
			}
			// fmt.Println(queue)
		}

		fmt.Println("Blueprint", b.id, "has max geodes", max_geodes)
		quality_level += b.id * max_geodes
	}
	fmt.Println("Answer 1:", quality_level)
}

func geodes_found(b *blueprint, p plan, minutes int, print bool) (int, plan) {
	if print {
		fmt.Print("Blueprint ", b)
	}

	executed_plan := plan{}

	ore_robots := 1
	clay_robots := 0
	obsidian_robots := 0
	geode_robots := 0

	ore_found := 0
	clay_found := 0
	obsidian_found := 0
	geode_found := 0

	for i := 0; i < minutes; i++ {
		if print {
			fmt.Println("Minute", i+1)
		}
		new_ore_robots := ore_robots
		new_clay_robots := clay_robots
		new_obsidian_robots := obsidian_robots
		new_geode_robots := geode_robots

		next_choice := none
		if len(p) > 0 {
			next_choice = p[0]
		}

		if next_choice == ore {
			if b.ore_ore <= ore_found {
				ore_found -= b.ore_ore
				new_ore_robots++
				executed_plan = append(executed_plan, ore)
				p = p[1:]
				if print {
					fmt.Println("Spend", b.ore_ore, "ore to start building a ore-collecting robot")
				}
			}
		} else if next_choice == clay {
			if b.clay_ore <= ore_found {
				ore_found -= b.clay_ore
				new_clay_robots++
				executed_plan = append(executed_plan, clay)
				p = p[1:]
				if print {
					fmt.Println("Spend", b.clay_ore, "ore to start building a clay-collecting robot")
				}
			}
		} else if next_choice == obsidian {
			if b.obsidian_ore <= ore_found && b.obsidian_clay <= clay_found {
				ore_found -= b.obsidian_ore
				clay_found -= b.obsidian_clay
				new_obsidian_robots++
				executed_plan = append(executed_plan, obsidian)
				p = p[1:]
				if print {
					fmt.Println("Spend", b.obsidian_ore, "ore and", b.obsidian_clay, "clay to start building a obsidian-collecting robot")
				}
			}
		} else {
			if b.geode_ore <= ore_found && b.geode_obsidian <= obsidian_found {
				ore_found -= b.geode_ore
				obsidian_found -= b.geode_obsidian
				new_geode_robots++
				if next_choice == geode {
					executed_plan = append(executed_plan, geode)
					p = p[1:]
				}
				if print {
					fmt.Println("Spend", b.geode_ore, "ore and", b.geode_obsidian, "obsidian to start building a geode-cracking robot")
				}
			}
		}

		ore_found += ore_robots
		clay_found += clay_robots
		obsidian_found += obsidian_robots
		geode_found += geode_robots
		if print {
			fmt.Println("ore robots", ore_robots, "ore found", ore_found)
			fmt.Println("clay robots", clay_robots, "clay found", clay_found)
			fmt.Println("obsidian robots", obsidian_robots, "obsidian found", obsidian_found)
			fmt.Println("geode robots", geode_robots, "geode found", geode_found)
		}

		ore_robots = new_ore_robots
		clay_robots = new_clay_robots
		obsidian_robots = new_obsidian_robots
		geode_robots = new_geode_robots
	}

	return geode_found, executed_plan
}
