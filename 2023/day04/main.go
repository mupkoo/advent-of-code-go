package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/mupkoo/advent-of-code-go/cast"
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

	for _, line := range parsed {
		segments := strings.FieldsFunc(line, splitter)

		if len(segments) != 3 {
			continue
		}

		winning := parseNumbers(segments[1])
		picked := parseNumbers(segments[2])
		current := 0

		for _, num := range picked {
			if slices.Contains(winning, num) {
				if current == 0 {
					current = 1
				} else {
					current *= 2
				}
			}
		}

		answer += current
	}

	return answer
}

func part2(input string) int {
	parsed := parseInput(input)
	won := make(map[int]int)

	for i, line := range parsed {
		segments := strings.FieldsFunc(line, splitter)

		if len(segments) != 3 {
			continue
		}

		winning := parseNumbers(segments[1])
		picked := parseNumbers(segments[2])
		won_index := 0
		won_multiplier := won[i] + 1

		for _, num := range picked {
			if slices.Contains(winning, num) {
				won_index += 1
				won[i+won_index] += won_multiplier
			}
		}

		won[i] += 1
	}

	answer := 0

	for _, v := range won {
		answer += v
	}

	return answer
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func splitter(r rune) bool {
	return r == ':' || r == '|'
}

func parseNumbers(input string) []int {
	var numbers []int

	for _, line := range strings.Split(input, " ") {
		if len(line) == 0 {
			continue
		}

		numbers = append(numbers, cast.ToInt(line))
	}
	return numbers
}
