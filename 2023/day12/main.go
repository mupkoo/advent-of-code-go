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
	parsed := parseInput(input, 1)
	return solve(parsed)
}

func part2(input string) int {
	parsed := parseInput(input, 5)
	return solve(parsed)
}

func solve(parsed []line) (sum int) {
	for _, l := range parsed {
		sum += agg(l.springs, l.damaged)
	}

	return sum
}

var cache = map[string]int{}

func agg(springs []string, damaged []int) int {
	if len(springs) == 0 {
		if len(damaged) == 0 {
			return 1
		}
		return 0
	}

	if len(damaged) == 0 {
		if slices.Contains(springs, "#") {
			return 0
		}
		return 1
	}

	key := strings.Join(springs, "") + " " + strings.Join(cast.ToStrings(damaged...), ",")
	if val, ok := cache[key]; ok {
		return val
	}

	result := 0

	if strings.Contains(".?", springs[0]) {
		result += agg(springs[1:], damaged)
	}

	if strings.Contains("#?", springs[0]) {
		if len(springs) >= damaged[0] && !slices.Contains(springs[:damaged[0]], ".") && (damaged[0] == len(springs) || springs[damaged[0]] != "#") {

			next_springs := []string{}
			if len(springs) > damaged[0]+1 {
				next_springs = springs[damaged[0]+1:]
			}

			next_damaged := []int{}
			if len(damaged) > 1 {
				next_damaged = damaged[1:]
			}

			result += agg(next_springs, next_damaged)
		}
	}

	cache[key] = result
	return result
}

type line struct {
	springs []string
	damaged []int
}

// .??..??...?##. 1,1,3
func parseInput(input string, repeat int) (parsed []line) {
	for _, l := range strings.Split(input, "\n") {
		segments := strings.Split(l, " ")
		parsed = append(parsed, line{
			springs: repeatSlice(strings.Split(segments[0], ""), repeat, true, "?"),
			damaged: repeatSlice(cast.ToInts(strings.Split(segments[1], ",")...), repeat, false, 0),
		})
	}

	return parsed
}

func repeatSlice[T any](slice []T, repeat int, shouldJoinBy bool, joinBy T) (result []T) {
	for i := 0; i < repeat; i++ {
		if shouldJoinBy && i > 0 {
			result = append(result, joinBy)
		}
		result = append(result, slice...)
	}
	return result
}

func debug(parsed []line) {
	for _, l := range parsed {
		fmt.Println(strings.Join(l.springs, ""), strings.Join(cast.ToStrings(l.damaged...), ","))
	}
	fmt.Println()
}
