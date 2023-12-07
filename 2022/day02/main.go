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
		fields := strings.Fields(line)
		a := aToType[fields[0]]
		b := bToType[fields[1]]

		answer += typeToScore[b]

		if a == b {
			answer += 3
		} else if typeWinsAgainst[b] == a {
			answer += 6
		}
	}

	return answer
}

func part2(input string) int {
	parsed := parseInput(input)
	answer := 0

	for _, line := range parsed {
		fields := strings.Fields(line)
		a := aToType[fields[0]]
		b := ""

		if fields[1] == "X" {
			b = typeWinsAgainst[a]
		} else if fields[1] == "Y" {
			b = a
		} else if fields[1] == "Z" {
			b = typeLosesAgainst[a]
		}

		answer += typeToScore[b]

		if a == b {
			answer += 3
		} else if typeWinsAgainst[b] == a {
			answer += 6
		}
	}

	return answer
}

var aToType = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
}

var bToType = map[string]string{
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

var typeToScore = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var typeWinsAgainst = map[string]string{
	"rock":     "scissors",
	"paper":    "rock",
	"scissors": "paper",
}

var typeLosesAgainst = map[string]string{
	"rock":     "paper",
	"paper":    "scissors",
	"scissors": "rock",
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}
