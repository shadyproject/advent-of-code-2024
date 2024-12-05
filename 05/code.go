package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/herrnan/aoc-helper"
)

const Day = 5
const Year = 2024

func GetInput(sample bool) string {
	var input string
	if sample {
		input = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
	} else {

		h, _ := aoc.NewHelper(Day, Year)
		input, _ = h.GetInput()
	}
	return input
}

func main() {
	input := GetInput(false)
	fmt.Printf("Day %v\n", Day)
	fmt.Printf("Part 1: %v\n", Part1(input))
	fmt.Printf("Part 2: %v\n", Part2(input))
}

func Part1(input string) (middlePageSum int) {
	parts := strings.Split(input, "\n\n")

	rules := strings.Split(parts[0], "\n")
	updates := strings.Split(parts[1], "\n")

	for _, u := range updates {
		updateRule := UpdateToRuleChain(u)
		updateInChain := RulesContainChain(rules, updateRule)

		if updateInChain {
			fmt.Printf("update %v found\n", u)
			page := FindMiddlePage(u)
			fmt.Printf("middle page: %v\n", page)

			middlePageSum += page
		}
	}

	return middlePageSum
}

func FindMiddlePage(update string) int {
	pages := strings.Split(update, ",")

	mid := len(pages) / 2

	page, _ := strconv.Atoi(pages[mid])
	return page
}

// create a rule representation of the update order
func UpdateToRuleChain(update string) []string {
	var ruleChain []string

	updates := strings.Split(update, ",")
	for i := 0; i < len(updates)-1; i++ {
		rule := updates[i] + "|" + updates[i+1]
		ruleChain = append(ruleChain, rule)
	}
	return ruleChain
}

func RulesContainChain(rules []string, ruleChain []string) bool {
	foundChain := false

	for _, entry := range ruleChain {
		foundRule := false
		for _, r := range rules {
			if entry == r {
				foundRule = true
				break
			}
		}
		if !foundRule {
			return false
		}
		foundChain = foundRule
	}

	return foundChain
}

func Part2(input string) int {
	return 0
}
