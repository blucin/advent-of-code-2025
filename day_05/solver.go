package day_05

import (
	"aoc-2025/utils"
	"log"
	"slices"
	"strconv"
	"strings"
)

func Solve() (string, string) {
	input, err := utils.ReadFile("inputs/day_05.txt")
	if err != nil {
		return "", ""
	}
	part1, part2 := 0, 0
	blocks := strings.Split(strings.TrimSpace(input), "\n\n")
	freshIdRanges := make([][2]int, 0)
	unknownIds := make([]int, 0)

	// parse fresh ids
	for _, idRange := range strings.Fields(blocks[0]) {
		lower, err := strconv.Atoi(strings.Split(idRange, "-")[0])
		if err != nil {
			log.Fatalln("Error convert lower limit of fresh id", err)
		}
		upper, err := strconv.Atoi(strings.Split(idRange, "-")[1])
		if err != nil {
			log.Fatalln("Error convert higher limit of fresh id", err)
		}
		freshIdRanges = append(freshIdRanges, [2]int{lower, upper})
	}

	// parse unknown ids
	for _, unknownId := range strings.Fields(blocks[1]) {
		id, err := strconv.Atoi(unknownId)
		if err != nil {
			log.Fatalln("Error convert unkown id", err)
		}
		unknownIds = append(unknownIds, id)
	}

	// part 1
	for _, unknownId := range unknownIds {
		for _, freshIdRange := range freshIdRanges {
			if freshIdRange[0] <= unknownId && unknownId <= freshIdRange[1] {
				part1 += 1
				break
			}
		}
	}

	// part 2
	slices.SortFunc(freshIdRanges, func(a, b [2]int) int {
		return a[0] - b[0]
	})
	mergedFreshIds := make([][2]int, 0)
	curr := freshIdRanges[0]
	for _, idRange := range freshIdRanges[1:] {
		// overlaps? merge
		if idRange[0] <= curr[1]+1 {
			if idRange[1] > curr[1] {
				curr[1] = idRange[1]
			}
		} else {
			mergedFreshIds = append(mergedFreshIds, curr)
			curr = idRange
		}
	}
	mergedFreshIds = append(mergedFreshIds, curr)

	for _, idRange := range mergedFreshIds {
		part2 += idRange[1] - idRange[0] + 1
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
