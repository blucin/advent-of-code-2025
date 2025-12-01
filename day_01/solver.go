package day_01

import (
	"aoc-2025/utils"
	"log"
	"strconv"
	"strings"
)

func Solve() (string, string) {
	input, err := utils.ReadFile("inputs/day_01.txt")
	if err != nil {
		return "", ""
	}
	part1, part2 := 0, 0
	init_state := 50
	crossed_zero_times := 0

	for _, line := range strings.Split(input, "\n") {
		init_state, crossed_zero_times = turnDial(init_state, line[0:])
		if init_state == 0 {
			part1 = part1 + 1
		}
		part2 = part2 + crossed_zero_times
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

// Returns new state after rotation and no. of times dial touched/crossed zero
func turnDial(state int, rotation string) (int, int) {
	dial_len := 100
	turnBy, err := strconv.Atoi(string(rotation[1:]))

	if err != nil {
		log.Fatalln(err)
	}

	if rotation[0] == 'L' {
		new_state := ((state-turnBy)%dial_len + dial_len) % dial_len
		crossed := 0
		for i := 1; i <= turnBy; i++ {
			pos := (state - i) % 100
			if pos == 0 {
				crossed++
			}
		}
		return new_state, crossed
	} else {
		new_state := ((state+turnBy)%dial_len + dial_len) % dial_len
		crossed := 0
		for i := 1; i <= turnBy; i++ {
			pos := (state + i) % 100
			if pos == 0 {
				crossed++
			}
		}
		return new_state, crossed
	}
}
