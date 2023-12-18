package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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
	items := []Item{}
	for _, line := range parseInput(input) {
		item := Item{}
		fmt.Sscanf(line, "%s %d (#%s)", &item.dir, &item.move)
		items = append(items, item)
	}

	return area(items)
}

// R 2 (#144d42)
// U 2 (#383193)
func part2(input string) int {
	items := []Item{}
	for _, line := range parseInput(input) {
		parts := strings.Split(line, " (#")
		hex := strings.TrimRight(parts[1], ")")
		move, _ := strconv.ParseInt(hex[:len(hex)-1], 16, 0)

		items = append(items, Item{
			string(hex[len(hex)-1]),
			int(move),
		})
	}

	return area(items)
}

var directions = map[string]Point{
	"R": {0, 1},
	"D": {1, 0},
	"L": {0, -1},
	"U": {-1, 0},
	"0": {0, 1},
	"1": {1, 0},
	"2": {0, -1},
	"3": {-1, 0},
}

type Point struct {
	row, col int
}

type Item struct {
	dir  string
	move int
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func area(items []Item) int {
	path := []Point{{0, 0}}
	pathLen := 0

	for i, item := range items {
		dir := directions[item.dir]
		current := path[i]
		pathLen += item.move

		path = append(path, Point{
			current.row + dir.row*item.move,
			current.col + dir.col*item.move,
		})
	}

	area := 0
	for i, point := range path {
		j := (i + 1) % len(path)
		area += point.col * path[j].row
		area -= point.row * path[j].col
	}

	area = abs(area) / 2
	insideOfArea := area - int(pathLen/2) + 1

	return insideOfArea + pathLen
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
