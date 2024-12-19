package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/herrnan/aoc-helper"
)

const Day = 6
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

type Location struct {
	Row    int
	Column int
}

func NewLocation(row, column int) Location {
	return Location{row, column}
}

func (l Location) String() string {
	return fmt.Sprintf("Row: %d Column: %d\n", l.Row, l.Column)
}

// true is passable false is not
type Map [][]bool

func (m Map) String() string {
	var b strings.Builder

	for _, row := range m {
		b.Write([]byte(fmt.Sprintf("%v\n", row)))
	}
	b.Write([]byte(fmt.Sprintf("Rows: %d Colums: %d", len(m), len(m[0]))))

	return b.String()
}

func NewMap(input string) (Map, Guard) {
	rows := strings.Split(input, "\n")
	m := make(Map, len(rows))
	var g Guard
	for r, row := range rows {
		cols := strings.Split(row, "")
		for c, col := range cols {
			if col == "^" {
				m[r] = append(m[r], true)
				g = NewGuard(r, c)
			} else if col == "#" {
				m[r] = append(m[r], false)
			} else {
				m[r] = append(m[r], true)
			}
		}
	}

	return m, g
}

const (
	UP    = "up"
	DOWN  = "down"
	LEFT  = "left"
	RIGHT = "right"
)

type Guard struct {
	position  Location
	facing    string
	stepCount int
	rotations map[string]string
}

func NewGuard(startRow int, startCol int) Guard {
	mr := map[string]string{
		UP:    RIGHT,
		RIGHT: DOWN,
		DOWN:  LEFT,
		LEFT:  UP,
	}
	return Guard{position: []int{startRow, startCol}, facing: UP, stepCount: 0, rotations: mr}
}

func (g Guard) String() string {
	return fmt.Sprintf("row:%d col:%d facing:%s", g.position[0], g.position[1], g.facing)
}

func (g *Guard) Walk(m Map) int {
	newLoc := NewLocation(0, 0)
	switch g.facing {
	case UP:
		newLoc.Row = g.position.Row - 1
		newLoc.Column = g.position.Column
	case DOWN:
		newLoc.Row = g.position.Row + 1
		newLoc.Column = g.position.Column
	case LEFT:
		newLoc.Row = g.position.Row
		newLoc.Column = g.position.Column - 1
	case RIGHT:
		newLoc.Row = g.position.Row
		newLoc.Column = g.position.Column + 1
	}

	if newLoc.Row >= len(m) || newLoc.Row < 0 {
		return g.stepCount
	} else if newLoc.Column >= len(m[g.position.Row]) || newLoc.Column < 0 {
		return g.stepCount
	}

	// not passable
	if !m[newLoc.Row][newLoc.Column] {
		g.facing = g.rotations[g.facing]
		fmt.Println("Rotating to %v", g.facing)
	} else {
		g.position = newLoc
		fmt.Println("Moving to %v", newLoc)
		g.stepCount += 1
	}

	return g.Walk(m)
}

func Part1(input string) int {
	fmt.Printf("%v\n", input)
	m, g := NewMap(input)
	fmt.Printf("%v\n", m)
	fmt.Printf("Guard at %s\n", g)
	return g.Walk(m)
}

func Part2(input string) int {
	return 0
}
