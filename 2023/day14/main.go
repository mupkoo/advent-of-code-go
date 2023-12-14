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
	var part, ans int
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
	grid = tiltNorth(grid)

	return sum(grid)
}

func part2(input string) int {
	grid := parseInput(input)

	loop := func() {
		for i := 0; i < 4; i++ {
			grid = tiltNorth(grid)
			grid = rotateRight(grid)
		}
	}

	cache := map[string]bool{gridToString(grid): true}
	store := []string{gridToString(grid)}
	counter := 0

	for {
		counter += 1
		loop()
		gridAsString := gridToString(grid)

		if _, ok := cache[gridAsString]; ok {
			break
		}

		cache[gridAsString] = true
		store = append(store, gridAsString)
	}

	index := slices.Index(store, gridToString(grid))
	gridAsString := store[(1000000000-index)%(counter-index)+index]

	return sum(parseInput(gridAsString))
}

func rotateRight(lines [][]string) [][]string {
	rotated := make([][]string, len(lines))
	for i := range rotated {
		rotated[i] = make([]string, len(lines))
	}

	for row, line := range lines {
		for col, char := range line {
			rotated[col][len(lines)-row-1] = char
		}
	}

	return rotated
}

func tiltNorth(grid [][]string) [][]string {
	for col := range grid[0] {
		for row := range grid {
			if grid[row][col] != "." {
				continue
			}

			for i := row + 1; i < len(grid); i++ {
				if grid[i][col] == "#" {
					break
				}

				if grid[i][col] == "O" {
					grid[row][col], grid[i][col] = "O", "." // swap
					break
				}
			}
		}
	}

	return grid
}

func gridToString(grid [][]string) string {
	segments := make([]string, 0, len(grid))
	for _, line := range grid {
		segments = append(segments, strings.Join(line, ""))
	}

	return strings.Join(segments, "\n")
}

func sum(lines [][]string) (sum int) {
	for i, line := range lines {
		count := 0
		for _, char := range line {
			if char == "O" {
				count++
			}
		}

		sum += count * (len(lines) - i)
	}

	return sum
}

// O....#....
// O.OO#....#
// .....##...
// OO.#O....O
// .O.....O#.
func parseInput(input string) (parsed [][]string) {
	for _, line := range strings.Split(input, "\n") {
		parsed = append(parsed, strings.Split(line, ""))
	}

	return parsed
}
