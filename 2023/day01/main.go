package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"unicode"

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
	total := 0

	for _, line := range parsed {
		var first string
		var last string

		for _, char := range line {
			if unicode.IsDigit(char) {
				if first == "" {
					first = string(char)
				}

				last = string(char)
			}
		}

		total = total + cast.ToInt(first+last)
	}

	return total
}

var digitMap = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"0":     "0",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func part2(input string) int {
	parsed := parseInput(input)
	total := 0

	for _, line := range parsed {
		firstIndex := 1000
		lastIndex := -1
		first := ""
		last := ""

		for k, v := range digitMap {
			tempFirstIndex := strings.Index(line, k)
			if tempFirstIndex != -1 && tempFirstIndex < firstIndex {
				firstIndex = tempFirstIndex
				first = v
			}

			tempLastIndex := strings.LastIndex(line, k)
			if tempLastIndex != -1 && tempLastIndex > lastIndex {
				lastIndex = tempLastIndex
				last = v
			}
		}

		total = total + cast.ToInt(first+last)
	}

	return total
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")
}
