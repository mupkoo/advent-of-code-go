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
	answer := 0

	for i, line := range parsed {
		for _, number := range line.numbers {
			symbols := line.symbols
			if i > 0 {
				symbols = append(symbols, parsed[i-1].symbols...)
			}
			if i < len(parsed)-1 {
				symbols = append(symbols, parsed[i+1].symbols...)
			}

			if isAdjacent(number, symbols...) {
				answer += cast.ToInt(number.value)
			}
		}
	}

	return answer
}

func part2(input string) int {
	parsed := parseInput(input)
	answer := 0

	for i, line := range parsed {
		for _, s := range line.symbols {
			if s.value == "*" {
				numbers := line.numbers
				if i > 0 {
					numbers = append(numbers, parsed[i-1].numbers...)
				}
				if i < len(parsed)-1 {
					numbers = append(numbers, parsed[i+1].numbers...)
				}

				adjacentNumbers := findAdjacentNumbers(s, numbers)

				if len(adjacentNumbers) == 2 {
					answer += cast.ToInt(adjacentNumbers[0].value) * cast.ToInt(adjacentNumbers[1].value)
				}
			}
		}
	}

	return answer
}

type number struct {
	value      string
	start, end int
}

type symbol struct {
	value string
	index int
}

type line struct {
	numbers []number
	symbols []symbol
}

func parseInput(input string) (ans []line) {
	for _, l := range strings.Split(input, "\n") {
		cur := line{}
		cur_number := number{start: -1, end: -1}

		for i, char := range l {
			if unicode.IsDigit(char) {
				if cur_number.start == -1 {
					cur_number.start = i
				}

				cur_number.end = i
				cur_number.value += string(char)

				continue
			}

			if cur_number.start != -1 {
				cur.numbers = append(cur.numbers, cur_number)
				cur_number = number{start: -1, end: -1}
			}

			if char != '.' {
				cur.symbols = append(cur.symbols, symbol{value: string(char), index: i})
			}
		}

		if cur_number.start != -1 {
			cur.numbers = append(cur.numbers, cur_number)
		}

		ans = append(ans, cur)
	}

	return ans
}

func isAdjacent(a number, symbols ...symbol) bool {
	for _, s := range symbols {
		if s.index >= a.start-1 && s.index <= a.end+1 {
			return true
		}
	}

	return false
}

func findAdjacentNumbers(s symbol, numbers []number) (ans []number) {
	for _, n := range numbers {
		if isAdjacent(n, s) {
			ans = append(ans, n)
		}
	}

	return ans
}
