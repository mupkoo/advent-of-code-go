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
		segments := strings.Split(line, ":")
		day := cast.ToInt(strings.Replace(segments[0], "Game ", "", 1))
		sets := strings.FieldsFunc(segments[1], splitter)

		answer += day

		for _, set := range sets {
			count_color := strings.Split(strings.Trim(set, " "), " ")
			count := cast.ToInt(count_color[0])
			color := count_color[1]

			// only 12 red cubes, 13 green cubes, and 14 blue cubes
			if (color == "red" && count > 12) || (color == "green" && count > 13) || (color == "blue" && count > 14) {
				answer -= day
				break
			}
		}
	}

	return answer
}

func part2(input string) int {
	parsed := parseInput(input)
	answer := 0

	for _, line := range parsed {
		segments := strings.Split(line, ":")
		cubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, cube := range strings.FieldsFunc(segments[1], splitter) {
			count_color := strings.Split(strings.Trim(cube, " "), " ")
			count := cast.ToInt(count_color[0])
			color := count_color[1]

			if cubes[color] < count {
				cubes[color] = count
			}
		}

		answer += cubes["red"] * cubes["green"] * cubes["blue"]
	}

	return answer
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")
}

func splitter(r rune) bool {
	return r == ';' || r == ','
}
