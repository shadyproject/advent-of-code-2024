package main

import (
	"fmt"
	"strings"

	"github.com/herrnan/aoc-helper"
)

func main() {
	h, _ := aoc.NewHelper(4, 2024)
	input, _ := h.GetInput()

	fmt.Printf("Day 4\n")
	fmt.Printf("Part 1: %v\n", Part1(input))
	fmt.Printf("Part 2: %v\n", Part2(input))
}

type Grid [][]string

func MakeGrid(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := make(Grid, len(lines)-1)
	for x, line := range lines {
		if line == "" {
			continue
		}
		columns := strings.Split(line, "")
		grid[x] = columns
	}

	return grid
}

func (g Grid) ElementAt(x, y int) string {
	if x < 0 || y < 0 || x >= len(g) || y >= len(g[x]) {
		return ""
	}

	return g[x][y]
}

func (g Grid) CountAroundPoint(s string, x, y int) (count int) {
	for xDir := -1; xDir <= 1; xDir++ {
		for yDir := -1; yDir <= 1; yDir++ {
			var found string
			for idx := 0; idx < len(s); idx++ {
				found += g.ElementAt(x+xDir*idx, y+yDir*idx)
			}
			if found == s {
				count += 1
			}
		}
	}

	return count
}

func (g Grid) IsXMas(x, y int) bool {
	apex := g.ElementAt(x, y)
	if apex != "A" {
		return false
	}

	stroke1 := g.ElementAt(x-1, y-1) + apex + g.ElementAt(x+1, y+1)
	stroke2 := g.ElementAt(x-1, y+1) + apex + g.ElementAt(x+1, y-1)

	if stroke1 == "MAS" && stroke2 == "SAM" {
		return true
	}

	if stroke1 == "SAM" && stroke2 == "MAS" {
		return true
	}

	if stroke1 == "MAS" && stroke2 == "MAS" {
		return true
	}

	if stroke1 == "SAM" && stroke2 == "SAM" {
		return true
	}

	return false
}

func Part1(input string) (count int) {
	grid := MakeGrid(input)

	for x := range grid {
		for y := range grid[x] {
			count += grid.CountAroundPoint("XMAS", x, y)
		}
	}

	return count
}

func Part2(input string) (count int) {
	grid := MakeGrid(input)

	for x := range grid {
		for y := range grid[x] {
			if grid.IsXMas(x, y) {
				count += 1
			}
		}
	}

	return count
}
