package day_04

import (
	"aoc-2025/utils"
	"strconv"
	"strings"
)

func Solve() (string, string) {
	input, err := utils.ReadFile("inputs/day_04.txt")
	if err != nil {
		return "", ""
	}
	part1, part2 := 0, 0

	grid := [][]rune{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		grid = append(grid, []rune(line))
	}
	_, part1 = removePaper(cloneGrid(grid), 0, false)
	_, part2 = removePaper(cloneGrid(grid), 0, true)
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func cloneGrid(src [][]rune) [][]rune {
	dst := make([][]rune, len(src))
	for i := range src {
		dst[i] = make([]rune, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func removePaper(grid [][]rune, removedSoFar int, keepRemoving bool) ([][]rune, int) {
	removed := 0
	rows := len(grid)
	isValidPoint := func(x, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < len(grid[x])
	}

	toRemove := make([][2]int, 0)

	for x, row := range grid {
		for y, point := range row {
			if point != '@' {
				continue
			}

			directions := [][2]int{{0, 1}, {0, -1}, {1, 1}, {1, -1}, {1, 0}, {-1, 1}, {-1, -1}, {-1, 0}}

			if point == '@' {
				neighbourCnt := 0
				for _, dir := range directions {
					nx := x + dir[0]
					ny := y + dir[1]
					if !isValidPoint(nx, ny) {
						continue
					}
					if grid[nx][ny] == '@' {
						neighbourCnt += 1
					}
				}
				if neighbourCnt < 4 {
					toRemove = append(toRemove, [2]int{x, y})
					removed += 1
				}
			}
		}
	}

	for _, p := range toRemove {
		grid[p[0]][p[1]] = 'x'
	}

	if keepRemoving == false {
		return grid, removedSoFar + removed
	}
	if removed == 0 {
		return grid, removedSoFar
	}
	return removePaper(grid, removedSoFar+removed, true)
}
