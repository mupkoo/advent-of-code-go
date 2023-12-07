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
	houses := make(map[cords]int)
	santa_cords := cords{0, 0}
	houses[santa_cords] = 1

	for _, c := range parsed {
		santa_cords.move(c)
		houses[santa_cords]++
	}

	return len(houses)
}

func part2(input string) int {
	parsed := parseInput(input)
	houses := make(map[cords]int)
	santa_cords := cords{0, 0}
	robot_cords := cords{0, 0}
	houses[santa_cords] = 1

	for i := 0; i < len(parsed); i += 2 {
		santa_cords.move(parsed[i])
		houses[santa_cords]++

		robot_cords.move(parsed[i+1])
		houses[robot_cords]++
	}

	return len(houses)
}

type cords struct {
	x, y int
}

func (c *cords) move(dir string) {
	switch dir {
	case "^":
		c.y++
	case "v":
		c.y--
	case ">":
		c.x++
	case "<":
		c.x--
	}
}

func parseInput(input string) []string {
	return strings.Split(input, "")
}
