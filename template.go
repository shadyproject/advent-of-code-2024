package main

import (
	"fmt"
	"os"

	"github.com/herrnan/aoc-helper"
)

const Day = 0
const Year = 2024

func GetSampleData() string {
	bytes, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Printf("Error loading sample data: %v\n", err)
	}

	return string(bytes)
}

func GetInput(sample bool) string {
	var input string
	if sample {
		input = GetSampleData()
	} else {

		h, _ := aoc.NewHelper(Day, Year)
		input, _ = h.GetInput()
	}
	return input
}

func main() {

	input := GetInput(true)

	fmt.Printf("Day %v\n", Day)
	fmt.Printf("Part 1: %v\n", Part1(input))
	fmt.Printf("Part 2: %v\n", Part2(input))
}

func Part1(input string) int {
	return 0
}

func Part2(input string) int {
	return 0
}
