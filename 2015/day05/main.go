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

line_loop:
	for _, line := range parsed {
		chars := strings.Split(strings.TrimSpace(line), "")
		vowels_count := 0
		has_double := false

		if vowels[chars[0]] {
			vowels_count++
		}

		for i := 1; i < len(chars); i++ {
			prev := chars[i-1]
			char := chars[i]

			if bad[prev+char] {
				continue line_loop
			}

			if vowels[char] {
				vowels_count++
			}

			if prev == char {
				has_double = true
			}
		}

		if vowels_count >= 3 && has_double {
			answer++
		}
	}

	return answer
}

func part2(input string) int {
	parsed := parseInput(input)
	answer := 0

	for _, line := range parsed {
		pairs := make(map[string]int)
		chars := strings.Split(strings.TrimSpace(line), "")
		has_pair := false
		has_double := false

		fmt.Printf("Checking %v\n", line)

		// It contains a pair of any two letters that appears at least twice in
		// the string without overlapping, like xyxy (xy) or aabcdefgaa (aa),
		// but not like aaa (aa, but it overlaps).
		// It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.

		for i := 0; i < len(chars)-1; i++ {
			pair := chars[i] + chars[i+1]

			matches_with_prev := i > 0 && chars[i-1] == chars[i] && chars[i-1] == chars[i+1]

			if !has_pair && (i == 0 || !matches_with_prev) {
				pairs[pair]++

				has_pair = pairs[pair] > 1
			}

			if i > 0 && chars[i-1] == chars[i+1] {
				has_double = true
			}

			if has_pair && has_double {
				answer++
				break
			}
		}

		fmt.Printf("has_pair: %v, has_double: %v, pairs: %v\n", has_pair, has_double, pairs)
	}

	return answer
}

var vowels = map[string]bool{
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true,
}

var bad = map[string]bool{
	"ab": true,
	"cd": true,
	"pq": true,
	"xy": true,
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}
