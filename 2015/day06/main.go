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
	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}

	for _, ins := range parsed {
		for x := ins.start.x; x <= ins.end.x; x++ {
			for y := ins.start.y; y <= ins.end.y; y++ {
				switch ins.op {
				case "on":
					grid[x][y] = true
				case "off":
					grid[x][y] = false
				case "toggle":
					grid[x][y] = !grid[x][y]
				}
			}
		}
	}

	answer := 0

	for _, row := range grid {
		for _, light := range row {
			if light {
				answer++
			}
		}
	}

	return answer
}

func part2(input string) int {
	parsed := parseInput(input)
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, ins := range parsed {
		for x := ins.start.x; x <= ins.end.x; x++ {
			for y := ins.start.y; y <= ins.end.y; y++ {
				switch ins.op {
				case "on":
					grid[x][y]++
				case "off":
					grid[x][y]--
				case "toggle":
					grid[x][y] += 2
				}

				if grid[x][y] < 0 {
					grid[x][y] = 0
				}
			}
		}
	}

	answer := 0

	for _, row := range grid {
		for _, brightness := range row {
			answer += brightness
		}
	}

	return answer
}

type instruction struct {
	op    string // "on", "off", "toggle"
	start point
	end   point
}
type point struct {
	x, y int
}

// turn off 199,133 through 461,193
// toggle 537,781 through 687,941
// turn on 226,196 through 599,390
func parseInput(input string) (ins []instruction) {
	for _, line := range strings.Split(input, "\n") {
		var op string
		var start, end point
		line = strings.TrimSpace(line)
		line = strings.TrimPrefix(line, "turn ")
		fmt.Sscanf(line, "%s %d,%d through %d,%d", &op, &start.x, &start.y, &end.x, &end.y)

		ins = append(ins, instruction{
			op:    op,
			start: start,
			end:   end,
		})
	}

	return ins
}
