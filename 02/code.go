package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/herrnan/aoc-helper"
)

const MAX_SAFE_DELTA = 3
const MIN_SAFE_DELTA = 1
const LEVEL_INCREASING = "increasing"
const LEVEL_DESCREASING = "decreasing"
const LEVEL_EQUAL = "equal"

func main() {
	// lol what's error handling
	h, _ := aoc.NewHelper(2, 2024)
	input, _ := h.GetInput()
	// input := "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9\n14 13 12 9 8"
	// input := "14 13 12 9 8\n"

	fmt.Print("Day 2\n")
	fmt.Printf("Part 1 Answer: %v\n", Part1(input))
	fmt.Printf("Part 2 Answer: %v\n", Part2(input))
}

func Part1(input string) int {
	numSafe, numUnsafe := 0, 0
	reports := strings.Split(input, "\n")
	fmt.Printf("Processing %v reports", len(reports))

	for _, report := range reports {
		if report == "" {
			continue
		}
		if IsReportSafe(report) {
			numSafe += 1
			fmt.Printf("SAFE: %v\n", report)
		} else {
			numUnsafe += 1
			fmt.Printf("UNSAFE: %v\n", report)
		}
	}

	fmt.Printf("Safe: %v Unsafe: %v Total: %v\n", numSafe, numUnsafe, len(reports))
	return numSafe
}

func IsReportSafe(report string) bool {
	safe := true
	levels := strings.Fields(report)
	trend := LevelTrend(levels[0], levels[1])

	for i := 0; i < len(levels)-1; i++ {
		if LevelTrend(levels[i], levels[i+1]) == LEVEL_EQUAL {
			safe = false
			break
		} else {
			if LevelTrend(levels[i], levels[i+1]) != trend {
				safe = false
				break
			}

			if !IsTrendSafe(levels[i], levels[i+1]) {
				safe = false
				break
			}
		}
	}

	return safe
}

func IsTrendSafe(level1 string, level2 string) bool {
	l1, _ := strconv.Atoi(level1)
	l2, _ := strconv.Atoi(level2)

	delta := Abs(l1 - l2)

	return MIN_SAFE_DELTA <= delta && delta <= MAX_SAFE_DELTA
}

func LevelTrend(level1 string, level2 string) string {
	l1, _ := strconv.Atoi(level1)
	l2, _ := strconv.Atoi(level2)
	if l1 < l2 {
		return LEVEL_INCREASING
	} else if l1 > l2 {
		return LEVEL_DESCREASING
	}

	return LEVEL_EQUAL
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Part2(input string) (answer string) {
	return ""
}
