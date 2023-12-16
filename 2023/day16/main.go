package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	var ans int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans = part1(input)
	} else {
		ans = part2(input)
	}

	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	grid := parseInput(input)
	return energize(grid, point{0, 0, right})
}

func part2(input string) int {
	grid := parseInput(input)

	max := 0
	for row := 0; row < len(grid); row++ {
		max = maxInt(
			max,
			energize(grid, point{row, 0, right}),
			energize(grid, point{row, len(grid[0]) - 1, left}),
		)
	}

	for col := 0; col < len(grid[0]); col++ {
		max = maxInt(
			max,
			energize(grid, point{0, col, down}),
			energize(grid, point{len(grid) - 1, col, up}),
		)
	}

	return max
}

func energize(grid [][]string, startPoint point) int {
	tracking := []*point{&startPoint}
	maxRow := len(grid)
	maxCol := len(grid[0])
	points := map[point]bool{}
	iter := 0

	for {
		next := []*point{}

		for _, p := range tracking {
			iter++
			if p.outOfBounds(maxCol, maxRow) || points[*p] {
				continue
			}

			points[*p] = true
			next = append(next, p)
			char := grid[p.row][p.col]

			if char == "-" && (p.dir == up || p.dir == down) {
				p.dir = right
				p2 := point{p.row, p.col, left}
				p2.move()
				next = append(next, &p2)
			} else if char == "|" && (p.dir == right || p.dir == left) {
				p.dir = down
				p2 := point{p.row, p.col, up}
				p2.move()
				next = append(next, &p2)
			} else if char == "\\" || char == "/" {
				p.dir = dirSwitch[char][p.dir]
			}

			p.move()
		}

		if len(next) == 0 {
			break
		}

		tracking = next
	}

	set := map[string]bool{}
	for p := range points {
		set[fmt.Sprintf("%v-%v", p.row, p.col)] = true
	}

	return len(set)
}

type direction int

const (
	right direction = iota
	left
	up
	down
)

type point struct {
	row, col int
	dir      direction
}

var dirSwitch = map[string]map[direction]direction{
	"\\": {
		right: down,
		left:  up,
		up:    left,
		down:  right,
	},
	"/": {
		right: up,
		left:  down,
		up:    right,
		down:  left,
	},
}

func (p *point) move() {
	switch p.dir {
	case right:
		p.col++
	case left:
		p.col--
	case up:
		p.row--
	case down:
		p.row++
	}
}

func (p *point) outOfBounds(maxCol, maxRow int) bool {
	return p.row < 0 || p.row >= maxRow || p.col < 0 || p.col >= maxCol
}

func parseInput(input string) (grid [][]string) {
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func maxInt(ints ...int) int {
	max := ints[0]
	for _, n := range ints[1:] {
		if n > max {
			max = n
		}
	}
	return max
}
