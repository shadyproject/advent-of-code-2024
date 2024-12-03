package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// h, _ := aoc.NewHelper(3, 2024)
	// input, _ := h.GetInput()
	// input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	fmt.Print("Day 3\n")
	fmt.Printf("Part 1 Answer: %v\n", Part1(input))
	fmt.Printf("Part 2 Answer: %v\n", Part2(input))
}

func Part1(input string) int {

	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

	sum := 0

	for _, inst := range re.FindAllString(input, -1) {
		// fmt.Printf("%v\n", inst)
		idx := strings.Index(inst, ",")
		end := strings.Index(inst, ")")
		lhs, _ := strconv.Atoi(inst[4:idx])
		rhs, _ := strconv.Atoi(inst[idx+1 : end])

		sum += (lhs * rhs)
	}

	return sum
}

func Part2(input string) int {
	return 0
}
