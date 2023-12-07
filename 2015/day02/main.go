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
	parsed := parseInput(input)
	answer := 0

	for _, d := range parsed {
		// the area of the smallest side.
		sides := []int{d.l * d.w, d.w * d.h, d.h * d.l}
		slices.Sort(sides)

		answer += 2*d.l*d.w + 2*d.w*d.h + 2*d.h*d.l + sides[0]
	}

	return answer
}

func part2(input string) int {
	parsed := parseInput(input)
	answer := 0

	for _, d := range parsed {
		sides := []int{d.l, d.w, d.h}
		slices.Sort(sides)

		answer += 2*sides[0] + 2*sides[1] + d.l*d.w*d.h
	}

	return answer
}

type dimension struct {
	l, w, h int
}

func parseInput(input string) (parsed []dimension) {
	for _, line := range strings.Split(input, "\n") {
		var d dimension
		fmt.Sscanf(line, "%dx%dx%d", &d.l, &d.w, &d.h)

		parsed = append(parsed, d)
	}

	return parsed
}
