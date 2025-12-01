package main

import (
	"fmt"
	"os"
	"strconv"

	"aoc-2025/day_01"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: aoc <day>")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid day: %v\n", err)
		os.Exit(1)
	}

	var part1, part2 string

	switch day {
	case 1:
		part1, part2 = day_01.Solve()
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
		return
	}

	fmt.Printf("Part 1: %s\n", part1)
	fmt.Printf("Part 2: %s\n", part2)
}
