package main

import (
	_ "embed"
	"flag"
	"fmt"
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
		last_items := calculate(line, true)

		for i := len(last_items) - 1; i >= 0; i-- {
			answer += last_items[i]
		}
	}

	return answer
}

func part2(input string) int {
	parsed := parseInput(input)
	answer := 0

	for _, line := range parsed {
		first_items := calculate(line, false)

		acc := 0

		for i := len(first_items) - 1; i >= 0; i-- {
			acc = first_items[i] - acc
		}

		answer += acc
	}

	return answer
}

func parseInput(input string) (parsed [][]int) {
	for _, line := range strings.Split(input, "\n") {
		items := []int{}
		for _, item := range strings.Split(line, " ") {
			items = append(items, cast.ToInt(item))
		}
		parsed = append(parsed, items)
	}

	return parsed
}

func calculate(line []int, use_last bool) []int {
	items := []int{}

	for {
		next_items := []int{}

		for i := 1; i < len(line); i++ {
			next_items = append(next_items, line[i]-line[i-1])
		}

		if use_last {
			items = append(items, line[len(line)-1])
		} else {
			items = append(items, line[0])
		}

		line = next_items

		if sum(next_items) == 0 {
			break
		}
	}

	return items
}

func sum(items []int) (sum int) {
	for _, item := range items {
		sum += item
	}
	return sum
}
