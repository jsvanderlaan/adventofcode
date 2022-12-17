package aoc2022

import (
	"fmt"
	"math/bits"
)

type row uint8

type rock struct {
	relative_height int
	rows            [4]row
}

type cycle struct {
	i_rock     uint8
	rock_count int64
	height     int64
}

func Day17(data string) {
	// start := time.Now()

	rock_map := map[uint8][4]row{
		0: {15 << 1, 0, 0, 0},               // horizontal line
		1: {1 << 3, 7 << 2, 1 << 3, 0},      // plus
		2: {7 << 2, 1 << 2, 1 << 2, 0},      // mirrored L
		3: {1 << 4, 1 << 4, 1 << 4, 1 << 4}, // vertical line
		4: {3 << 3, 3 << 3, 0, 0},           // square
	}

	const rock_target int64 = 1_000_000_000_000

	cycles := []cycle{}

	floor := int64(0)
	rock_count := int64(0)
	data_len := len(data)
	i_move := 0
	i_rock := uint8(0)

	rows := []row{}
	is_falling := false
	var falling_rock rock

	remaining_rocks := 0
	cycle_found := false

	var rocks_at_start_of_cycle int64
	var rocks_per_cycle int64
	var height_at_start_of_cycle int64
	var height_per_cycle int64
	var rocks_remaining_after_cycles int64
	var number_of_cycles_needed int64

	for !cycle_found || remaining_rocks >= 0 {
		if !is_falling {
			// create new rock
			is_falling = true
			rock_count++
			if cycle_found {
				remaining_rocks--
			}
			new_rock := rock_map[i_rock]
			i_rock++
			i_rock %= 5
			falling_rock = rock{len(rows) + 3, new_rock}
			// fmt.Println(floor + int64(len(rows)))
		}

		// if rock_count%100_000_000 == 0 {
		// 	fmt.Printf("Rock %d Execution time: %v\n", rock_count, time.Since(start))
		// }

		// print

		// move rock
		moveLeft := data[i_move%len(data)] == '<'
		canMove := true
		new_pos := [4]row{}
		if moveLeft {
			// fmt.Println("Move left")
			for i := 0; i < 4; i++ {
				new_row := falling_rock.rows[i] << 1
				if new_row >= 128 ||
					(falling_rock.relative_height+i < len(rows) &&
						rows[falling_rock.relative_height+i]&new_row > 0) {
					canMove = false
					break
				}
				new_pos[i] = new_row
			}
		} else {
			// fmt.Println("Move right")
			for i := 0; i < 4; i++ {
				new_row := falling_rock.rows[i] >> 1
				if bits.TrailingZeros8(uint8(falling_rock.rows[i])) == 0 ||
					(falling_rock.relative_height+i < len(rows) &&
						rows[falling_rock.relative_height+i]&new_row > 0) {
					canMove = false
					break
				}
				new_pos[i] = new_row
			}
		}
		if canMove {
			// fmt.Println("Can move to", new_pos)
			falling_rock.rows = new_pos
		}

		// drop rock
		canDrop := true
		for i := 0; i < 4; i++ {
			new_height := falling_rock.relative_height + i - 1
			if new_height < 0 || (new_height < len(rows) &&
				rows[falling_rock.relative_height+i-1]&falling_rock.rows[i] > 0) {
				canDrop = false
				break
			}
		}
		if canDrop {
			// fmt.Println("Can drop to", falling_rock.relative_height-1)
			falling_rock.relative_height--
		} else {
			is_falling = false
			// add rock to rows
			for i := 0; i < 4; i++ {
				if falling_rock.rows[i] != 0 {
					if falling_rock.relative_height+i >= len(rows) {
						rows = append(rows, falling_rock.rows[i])
					} else {
						rows[falling_rock.relative_height+i] |= falling_rock.rows[i]
					}
				}
			}

			// set new floor
			for i := len(rows) - 2; i > 1; i-- {
				compressedRows := rows[i] | rows[i+1]
				// isNewFloor := bits.LeadingZeros8(uint8(compressedRows)) <= 0 && bits.TrailingZeros8(uint8(compressedRows)) <= 0
				if compressedRows == 127 {
					// fmt.Println("Found a new floor", floor, i)

					floor += int64(i)
					rows = rows[i:]
					// print(falling_rock, rows)
					break
				}
			}

			if rows[len(rows)-1] == 127 && !cycle_found {
				fmt.Println("De vloer is plat", floor, i_rock)
				cycles = append(cycles, cycle{i_rock, rock_count, floor + int64(len(rows))})
				// print(falling_rock, rows)
			}

			if len(cycles) == 2 && !cycle_found {
				cycle_found = true
				rocks_at_start_of_cycle = cycles[0].rock_count
				rocks_per_cycle = cycles[1].rock_count - cycles[0].rock_count
				height_at_start_of_cycle = cycles[0].height
				height_per_cycle = cycles[1].height - cycles[0].height

				rocks_remaining_after_cycles = (rock_target - rocks_at_start_of_cycle) % rocks_per_cycle
				number_of_cycles_needed = (rock_target - rocks_at_start_of_cycle) / rocks_per_cycle

				fmt.Println("rocks_at_start_of_cycle", rocks_at_start_of_cycle)
				fmt.Println("rocks_per_cycle", rocks_per_cycle)
				fmt.Println("height_at_start_of_cycle", height_at_start_of_cycle)
				fmt.Println("height_per_cycle", height_per_cycle)
				fmt.Println("rocks_remaining_after_cycles", rocks_remaining_after_cycles)
				fmt.Println("number_of_cycles_needed", number_of_cycles_needed)

				remaining_rocks = int(rocks_remaining_after_cycles)
			}
		}

		i_move++
		i_move %= data_len

	}
	curr_height := floor + int64(len(rows))
	remaining_height := curr_height - (height_per_cycle + height_at_start_of_cycle)
	fmt.Println("Answer 2:", number_of_cycles_needed*height_per_cycle+height_at_start_of_cycle+remaining_height)
	// fmt.Println(cycles)
}

func print(falling_rock rock, rows []row) {
	max := falling_rock.relative_height + 4
	if len(rows) > max {
		max = len(rows)
	}
	for i := max - 1; i >= 0; i-- {
		x := i - falling_rock.relative_height
		for j := 0; j < 7; j++ {
			if x >= 0 && x < 4 && falling_rock.rows[x]&(64>>j) != 0 {
				fmt.Printf("@")
				continue
			}
			if i < len(rows) && rows[i]&(64>>j) != 0 {
				fmt.Printf("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
}

// 1_575_811_209_485 too low
