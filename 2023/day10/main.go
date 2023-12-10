package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/mupkoo/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	from, grid := parseInput(input)
	loop := findLoop(from, grid)

	return len(loop) / 2
}

func part2(input string) int {
	from, grid := parseInput(input)
	loop := findLoop(from, grid)
	input = strings.Replace(input, "S", pointToChar(from, grid), 1)
	area := 0

	fmt.Println(input)

	for row, line := range strings.Split(input, "\n") {
		inside := false
		var riding rune

		for col, char := range strings.TrimSpace(line) {
			partOfLoop := loop[point{col, row}]

			if !partOfLoop && inside {
				area++
			}

			if partOfLoop {
				if char == '|' {
					inside = !inside
				} else if slices.Contains([]rune{'F', 'L'}, char) {
					riding = char
				} else if riding == 'F' && char == '7' {
					char = '7'
				} else if riding == 'F' && char == 'J' {
					char = 'J'
					inside = !inside
				} else if riding == 'L' && char == 'J' {
					char = 'J'
				} else if riding == 'L' && char == '7' {
					char = '7'
					inside = !inside
				}
			}
		}
	}

	return area
}

type point struct {
	x, y int
}

func (p *point) next(from point, dir string) point {
	switch dir {
	case "|": // | is a vertical pipe connecting north and south.
		if from.y < p.y {
			return point{p.x, p.y + 1}
		} else {
			return point{p.x, p.y - 1}
		}

	case "-": // - is a horizontal pipe connecting east and west.
		if from.x < p.x {
			return point{p.x + 1, p.y}
		} else {
			return point{p.x - 1, p.y}
		}

	case "L": // L is a 90-degree bend connecting north and east.
		if from.y == p.y-1 {
			return point{p.x + 1, p.y}
		} else {
			return point{p.x, p.y - 1}
		}

	case "J": // J is a 90-degree bend connecting north and west.
		if from.y == p.y-1 {
			return point{p.x - 1, p.y}
		} else {
			return point{p.x, p.y - 1}
		}

	case "7": // 7 is a 90-degree bend connecting south and west.
		if from.y == p.y+1 {
			return point{p.x - 1, p.y}
		} else {
			return point{p.x, p.y + 1}
		}

	case "F": // F is a 90-degree bend connecting south and east.
		if from.y == p.y+1 {
			return point{p.x + 1, p.y}
		} else {
			return point{p.x, p.y + 1}
		}
	}

	panic("invalid direction")
}

func parseInput(input string) (point, map[point]string) {
	grid := make(map[point]string)
	start := point{0, 0}

	for y, line := range strings.Split(input, "\n") {
		for x, char := range strings.TrimSpace(line) {
			if char == '.' {
				continue
			}

			p := point{x, y}

			if char == 'S' {
				start = p
			}

			grid[p] = string(char)
		}
	}

	return start, grid
}

func findFirst(start point, grid map[point]string) point {
	char := pointToChar(start, grid)

	if char == "-" || char == "F" || char == "L" {
		return point{start.x + 1, start.y}
	} else if char == "|" || char == "7" {
		return point{start.x, start.y + 1}
	} else if char == "J" {
		return point{start.x, start.y - 1}
	}

	panic("no first")
}

func findLoop(from point, grid map[point]string) map[point]bool {
	current := findFirst(from, grid)
	loop := make(map[point]bool)
	loop[from] = true
	loop[current] = true

	for {
		next := current.next(from, grid[current])

		from = current
		current = next
		loop[current] = true

		if grid[current] == "S" {
			break
		}
	}

	return loop
}

func pointToChar(p point, grid map[point]string) string {
	possible := []string{"-", "|", "L", "J", "7", "F"}

	if pointMatches(point{p.x - 1, p.y}, grid, "-FL") {
		possible = intersect(possible, []string{"-", "J", "7"})
	}

	if pointMatches(point{p.x + 1, p.y}, grid, "-J7") {
		possible = intersect(possible, []string{"-", "L", "F"})
	}

	if pointMatches(point{p.x, p.y - 1}, grid, "|F7") {
		possible = intersect(possible, []string{"|", "L", "J"})
	}

	if pointMatches(point{p.x, p.y + 1}, grid, "|LJ") {
		possible = intersect(possible, []string{"|", "F", "7"})
	}

	if len(possible) != 1 {
		panic("invalid S")
	}

	return possible[0]
}

func pointMatches(p point, grid map[point]string, haystack string) bool {
	char, ok := grid[p]
	return ok && strings.Contains(haystack, char)
}

func intersect(a, b []string) []string {
	m := make(map[string]bool)
	for _, s := range a {
		m[s] = true
	}

	var c []string
	for _, s := range b {
		if m[s] {
			c = append(c, s)
		}
	}

	return c
}
