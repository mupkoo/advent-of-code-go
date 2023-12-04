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
	max := 0
	cur := 0

	for _, line := range parsed {
		if line == "" {
			if cur > max {
				max = cur
			}
			cur = 0
			continue
		}

		cur += cast.ToInt(line)
	}

	return max
}

func part2(input string) int {
	parsed := parseInput(input)
	max := []int{0, 0, 0}
	cur := 0

	for _, line := range parsed {
		if line == "" {
			max = appendMax(max, cur)
			cur = 0
			continue
		}

		cur += cast.ToInt(line)
	}

	sum := 0

	for _, num := range appendMax(max, cur) {
		sum += num
	}

	return sum
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func appendMax(max []int, cur int) []int {
	for i := 0; i < 3; i++ {
		if cur > max[i] {
			for j := 2; j > i; j-- {
				max[j] = max[j-1]
			}
			max[i] = cur
			break
		}
	}

	return max
}
