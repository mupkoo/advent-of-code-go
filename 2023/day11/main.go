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
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := solve(input, 1)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := solve(input, 1_000_000-1)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func solve(input string, expandedBy int) int {
	parsed, expandedRows, expandedCols := parseInput(input, 1)
	pairs := [][]point{}

	for row, cols := range parsed {
		for col, char := range cols {
			if char == "." {
				continue
			}

			for i := col + 1; i < len(cols); i++ {
				if cols[i] == "#" {
					pairs = append(pairs, []point{{row, col}, {row, i}})
				}
			}

			for i := row + 1; i < len(parsed); i++ {
				for j, c := range parsed[i] {
					if c == "#" {
						pairs = append(pairs, []point{{row, col}, {i, j}})
					}
				}
			}
		}
	}

	sum := 0

	for _, pair := range pairs {
		sum += abs(pair[0].x-pair[1].x) + abs(pair[0].y-pair[1].y)
		sum += getExpansion(pair[0].x, pair[1].x, expandedRows, expandedBy)
		sum += getExpansion(pair[0].y, pair[1].y, expandedCols, expandedBy)
	}

	return sum
}

type point struct {
	x, y int
}

// ...#......
// .......#..
// #.........
// ..........
// ......#...
// .#........
// .........#
// ..........
// .......#..
// #...#.....
func parseInput(input string, expand int) (grid [][]string, expandedRows map[int]bool, expendedCols map[int]bool) {
	expandedRows = make(map[int]bool)
	expendedCols = make(map[int]bool)

	for i, line := range strings.Split(input, "\n") {
		cols := []string{}
		for _, col := range line {
			cols = append(cols, string(col))
		}

		grid = append(grid, cols)

		if len(line) == strings.Count(line, ".") {
			expandedRows[i] = true
		}
	}

	for colIndex := 0; colIndex < len(grid[0]); colIndex++ {
		col := []string{}
		for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
			col = append(col, grid[rowIndex][colIndex])
		}

		if len(col) == strings.Count(strings.Join(col, ""), ".") {
			expendedCols[colIndex] = true
		}
	}

	return grid, expandedRows, expendedCols
}

func getExpansion(a, b int, expanded map[int]bool, expandedBy int) int {
	if a < b {
		return expandBy(a, b, expanded, expandedBy)
	} else {
		return expandBy(b, a, expanded, expandedBy)
	}
}

func expandBy(a, b int, expanded map[int]bool, expandedBy int) int {
	sum := 0

	for i := a + 1; i < b; i++ {
		if expanded[i] {
			sum += expandedBy
		}
	}

	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
