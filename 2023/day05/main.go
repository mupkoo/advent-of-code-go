package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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
	seeds, maps := parseInput(input)
	closest := math.MaxInt

	for _, seed := range seeds {
		next := seed

		for _, m := range maps {
			for _, row := range m {

				if next >= row[1] && next <= row[1]+row[2] {
					next = next - row[1] + row[0]
					break
				}
			}
		}

		if closest > next {
			closest = next
		}
	}

	return closest
}

func part2(input string) int {
	seeds, maps := parseInput(input)
	closest := math.MaxInt

	for _, group := range groupSeeds(seeds) {
		for seed := group[0]; seed <= group[0]+group[1]; seed += 1 {
			next := seed

			for _, m := range maps {
				for _, row := range m {

					if next >= row[1] && next <= row[1]+row[2] {
						next = next - row[1] + row[0]
						break
					}
				}
			}

			if closest > next {
				closest = next
			}
		}
	}

	return closest
}

func parseInput(input string) (seeds []int, maps [][][]int) {
	segments := strings.Split(input, "\n\n")
	maps = make([][][]int, len(segments)-1)

	for _, seed := range strings.Split(segments[0], " ")[1:] {
		seeds = append(seeds, cast.ToInt(seed))
	}

	for i, segment := range segments[1:] {
		lines := strings.Split(segment, "\n")

		for _, line := range lines[1:] {
			maps[i] = append(maps[i], cast.SplitToInts(line, " "))
		}
	}

	return seeds, maps
}

func groupSeeds(seeds []int) [][]int {
	groups := make([][]int, 0, (len(seeds)+1)/2)

	for i := 0; i < len(seeds); i += 2 {
		end := i + 2
		if end > len(seeds) {
			end = len(seeds)
		}

		groups = append(groups, seeds[i:end])
	}

	return groups
}
