package day_02

import (
	"aoc-2025/utils"
	"log"
	"strconv"
	"strings"
)

func Solve() (string, string) {
	input, err := utils.ReadFile("inputs/day_02.txt")
	if err != nil {
		return "", ""
	}
	part1, part2 := 0, 0
	parsedInput := strings.Split(strings.ReplaceAll(input, "\n", ""), ",")

	for _, id_range := range parsedInput {
		id_range_split := strings.Split(id_range, "-")
		low, err := strconv.Atoi(id_range_split[0])
		if err != nil {
			log.Fatalf("Error converting low range: %v", err)
		}
		high, err := strconv.Atoi(id_range_split[1])
		if err != nil {
			log.Fatalf("Error converting low range: %v", err)
		}

		for id := low; id <= high; id++ {
			id_str := strconv.Itoa(id)
			mid := len(id_str) / 2

			// part-2
			for size := 1; size <= len(id_str)/2; size++ {
				prefixStr := id_str[:size]
				if len(id_str)%size != 0 {
					continue
				}

				repeated := ""
				for i := 0; i < (len(id_str) / size); i++ {
					repeated += prefixStr
				}
				if repeated == id_str {
					part2 += id
					break
				}
			}

			// part-1
			if len(id_str)%2 != 0 {
				continue
			}
			if id_str[:mid] == id_str[mid:] {
				part1 += id
			}
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
