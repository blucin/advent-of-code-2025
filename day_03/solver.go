package day_03

import (
	"aoc-2025/utils"
	"log"
	"strconv"
	"strings"
)

func Solve() (string, string) {
	input, err := utils.ReadFile("inputs/day_03.txt")
	if err != nil {
		return "", ""
	}
	part1, part2 := 0, 0

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		batteryBank := strings.TrimSpace(line)
		max2DigitVolt, err := findLargestVoltage(batteryBank, 2)
		if err != nil {
			log.Fatalln("Error getting max2DigitVolt", err)
		}
		max12DigitVolt, err := findLargestVoltage(batteryBank, 12)
		if err != nil {
			log.Fatalln("Error getting max12DigitVolt", err)
		}
		part1 += max2DigitVolt
		part2 += max12DigitVolt
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

// To build the largest k-digit number from a string of length n,
// At each pick, we must leave enough space for remaining digits to be picked from
//
// For k = 2:
//
//	First digit: search in array[start : n-1]   // must leave 1 digit after it
//	Second digit: search in array[start : n]    // can use any remaining digit
func findLargestVoltage(batteryBank string, k int) (int, error) {
	// constants
	n := len(batteryBank)
	start := 0
	numFound := 0
	maxVoltStr := ""

	for numFound < k {
		remaining := k - numFound
		end := n - remaining + 1

		// find max digit and its index in [start, end] window
		// let bestIdx be the index of that max digit in the current window
		bestIdx := start

		for i := start; i < end; i++ {
			if int(batteryBank[i]-'0') > int(batteryBank[bestIdx]-'0') {
				bestIdx = i
			}
		}

		maxVoltStr += string(batteryBank[bestIdx])
		start = bestIdx + 1
		numFound++
	}

	maxVolt, err := strconv.Atoi(maxVoltStr)
	if err != nil {
		return 0, err
	}
	return maxVolt, nil
}
