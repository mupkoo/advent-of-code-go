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
	answer := 1

	for _, set := range parsed {
		cur := 0
		time, distance := set[0], set[1]

		for i := 0; i < time; i++ {
			if i*(time-i) > distance {
				cur += 1
			}
		}

		answer *= cur
	}

	return answer
}

func part2(input string) int {
	parsed := parseInput(input)
	time_temp := ""
	distance_temp := ""
	answer := 0

	for _, set := range parsed {
		time_temp += cast.ToString(set[0])
		distance_temp += cast.ToString(set[1])
	}

	time := cast.ToInt(time_temp)
	distance := cast.ToInt(distance_temp)

	for i := 0; i < time; i++ {
		if i*(time-i) > distance {
			answer += 1
		} else if answer > 1 {
			break
		}
	}

	return answer
}

// Time:      7  15   30
// Distance:  9  40  200
func parseInput(input string) (parsed [][2]int) {
	lines := strings.Split(input, "\n")
	times := parseLine(lines[0])
	distances := parseLine(lines[1])

	for i, time := range times {
		parsed = append(parsed, [2]int{
			time,
			distances[i],
		})
	}

	return parsed
}

func parseLine(line string) (parsed []int) {
	part := strings.Split(line, ":")[1]
	part = strings.TrimSpace(part)

	for _, element := range strings.Fields(part) {
		parsed = append(parsed, cast.ToInt(strings.TrimSpace(element)))
	}

	return parsed
}
