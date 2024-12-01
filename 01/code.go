package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/herrnan/aoc-helper"
)

func main() {
	h, err := aoc.NewHelper(1, 2024)
	exitIfError(err)

	input, err := h.GetInput()
	exitIfError(err)

	fmt.Printf("Day 1 answer: %v\n", Part1(input))
	fmt.Printf("Day 1 part 2 answer: %v", Part2(input))
}

func exitIfError(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Part1(input string) (totalDistance int) {
	left := make([]int, len(input))
	right := make([]int, len(input))

	lines := strings.Split(input, "\n")

	for i, line := range lines {
		fields := strings.Fields(line)
		if line == "" {
			continue
		}
		left[i], _ = strconv.Atoi(fields[0])
		right[i], _ = strconv.Atoi(fields[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	for i := 0; i < len(left); i++ {
		distance := Abs(left[i] - right[i])
		totalDistance += distance
	}

	return
}

func Part2(input string) (similarityScore int) {
	left := make([]int, len(input))
	right := make([]int, len(input))

	lines := strings.Split(input, "\n")

	for i, line := range lines {
		fields := strings.Fields(line)
		if line == "" {
			continue
		}
		left[i], _ = strconv.Atoi(fields[0])
		right[i], _ = strconv.Atoi(fields[1])
	}

	for _, l := range left {
		score := l * Count(l, right)
		similarityScore += score
	}

	return similarityScore
}

func Count(x int, array []int) (count int) {
	for _, n := range array {
		if n == x {
			count += 1
		}
	}
	return count
}
