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
	parsed := parseInput(input, false)
	slices.SortFunc(parsed, createSortFunc(cardRanks))

	return sum(parsed)
}

func part2(input string) int {
	cardRanksWithJokers := copyMap(cardRanks)
	cardRanksWithJokers["J"] = 0

	parsed := parseInput(input, true)
	slices.SortFunc(parsed, createSortFunc(cardRanksWithJokers))

	return sum(parsed)
}

type hand struct {
	cards []string
	bid   int
	score int
}

func parseInput(input string, withJokers bool) (parsed []hand) {
	for _, line := range strings.Split(input, "\n") {
		segments := strings.Fields(line)
		cards := strings.Split(segments[0], "")

		parsed = append(parsed, hand{
			strings.Split(segments[0], ""),
			cast.ToInt(segments[1]),
			score(cards, withJokers),
		})
	}

	return parsed
}

var cardRanks = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func score(cards []string, withJokers bool) int {
	groups := make(map[string]int)

	for _, card := range cards {
		groups[card] += 1
	}

	if withJokers && groups["J"] > 0 {
		group := ""
		count := 0

		// find the group with the most cards
		for k, v := range groups {
			if k != "J" && v > count {
				group = k
				count = v
			}
		}

		// add jokers to the group with the most cards
		groups[group] += groups["J"]
		groups["J"] = 0
	}

	sum := 0
	for _, v := range groups {
		sum += v * v
	}

	return sum
}

func createSortFunc(ranks map[string]int) func(a, b hand) int {
	return func(a, b hand) int {
		if a.score != b.score {
			return a.score - b.score
		} else {
			for i, card := range a.cards {
				if ranks[card] != ranks[b.cards[i]] {
					return ranks[card] - ranks[b.cards[i]]
				}
			}

			return 0
		}
	}
}

func sum(hands []hand) (sum int) {
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}

	return sum
}

func copyMap(source map[string]int) map[string]int {
	destination := make(map[string]int)

	for key, value := range source {
		destination[key] = value
	}

	return destination
}
