package main

import (
	"cmp"
	"fmt"
	"slices"
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
	// input := "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9"
	// input := "14 13 12 9 8\n"
	// input := "34 33 34 35 38 41 42 45\n"
	// input := "82 84 85 87 90 92 93 91\n91 82 84 85 87 90 92 93"

	fmt.Print("Day 2\n")
	fmt.Printf("Part 1 Answer: %v\n", Part1(input))
	fmt.Printf("Part 2 Answer: %v\n", Part2(input))
}

func Part1(input string) int {
	numSafe, numUnsafe := 0, 0
	reports := strings.Split(input, "\n")
	// fmt.Printf("Processing %v reports", len(reports))

	for _, report := range reports {
		if report == "" {
			continue
		}
		if IsReportSafe(report) {
			numSafe += 1
			// fmt.Printf("SAFE: %v\n", report)
		} else {
			numUnsafe += 1
			// fmt.Printf("UNSAFE: %v\n", report)
		}
	}

	// fmt.Printf("Safe: %v Unsafe: %v Total: %v\n", numSafe, numUnsafe, len(reports))
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
	} else if l1 > l2 { // I had a _very_ stupid bug here wherein I was comparing the string versions and not the int versions
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

func ParseReports(input string) [][]int {
	var reports [][]int

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		chunks := strings.Split(line, " ")
		var levels []int
		for _, l := range chunks {
			level, _ := strconv.Atoi(l)
			levels = append(levels, level)
		}

		reports = append(reports, levels)
	}

	return reports
}

func Part2(input string) int {
	reports := ParseReports(input)
	numSafe := 0

	for _, report := range reports {
		idx := DampSafe(report)
		if idx == -1 {
			numSafe += 1
		} else {
			if DampSafe(RemoveElement(report, idx-1)) == -1 || DampSafe(RemoveElement(report, idx)) == -1 || DampSafe(RemoveElement(report, 0)) == -1 {
				numSafe += 1
			}
		}
	}

	return numSafe
}

func DampSafe(report []int) int {
	prev := 0
	trend := 0

	for i, level := range report {
		if prev == 0 {
			prev = level
		} else {
			currentTrend := cmp.Compare(prev, level)
			delta := Abs(prev - level)

			if delta > MAX_SAFE_DELTA || currentTrend == 0 {
				return i
			}

			if trend == 0 {
				trend = currentTrend
			} else if trend != currentTrend {
				return i
			}

			prev = level
		}
	}

	return -1
}

func RemoveElement[S ~[]E, E any](s S, i int) S {
	result := make([]E, len(s))
	copy(result, s)
	return slices.Delete(result, i, i+1)
}
