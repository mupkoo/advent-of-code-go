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
	return solve(input, 0)
}

func part2(input string) int {
	return solve(input, 1)
}

func solve(input string, allowedDiff int) (sum int) {
	blocks := parseInput(input)

	for _, block := range blocks {
		lines := strings.Split(block, "\n")
		cols := linesToCols(lines)

		for i := 1; i < len(lines); i++ {
			if currentDiff := calculateDiff(lines, i, allowedDiff); currentDiff == allowedDiff {
				sum += i * 100
				break
			}
		}

		for i := 1; i < len(cols); i++ {
			if currentDiff := calculateDiff(cols, i, allowedDiff); currentDiff == allowedDiff {
				sum += i
				break
			}
		}
	}

	return sum
}

func parseInput(input string) []string {
	return strings.Split(input, "\n\n")
}

func calculateDiff(lines []string, i int, allowedDiff int) int {
	length := len(lines)
	expandBy := 0
	currentDiff := 0

	for {
		prev := i - 1 - expandBy
		next := i + expandBy
		expandBy++

		if prev < 0 || next >= length {
			break
		}

		if lines[prev] != lines[next] {
			currentDiff += diff(lines[prev], lines[next])
		}

		if currentDiff > allowedDiff {
			break
		}
	}

	return currentDiff
}

func linesToCols(lines []string) []string {
	cols := make([]string, len(lines[0]))
	for i := range cols {
		for _, line := range lines {
			cols[i] += string(line[i])
		}
	}

	return cols
}

func diff(a, b string) (diff int) {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}

		if diff > 1 {
			return diff
		}
	}

	return diff
}
