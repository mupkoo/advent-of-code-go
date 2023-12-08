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
	left_right, instructions := parseInput(input)
	current := "AAA"
	steps := 0
	dir_len := len(left_right)
	i := 0

	for {
		dir := left_right[i]

		if dir == "L" {
			current = instructions[current].left
		} else {
			current = instructions[current].right
		}

		steps++
		i = (i + 1) % dir_len

		if current == "ZZZ" {
			break
		}
	}

	return steps
}

func part2(input string) int {
	left_right, instructions := parseInput(input)
	currents := []string{}
	solutions := make(map[string]int)
	steps := 0
	dir_len := len(left_right)
	i := 0

	for k := range instructions {
		if strings.HasSuffix(k, "A") {
			currents = append(currents, k)
		}
	}

	for {
		dir := left_right[i]

		for j := 0; j < len(currents); j++ {
			if dir == "L" {
				currents[j] = instructions[currents[j]].left
			} else {
				currents[j] = instructions[currents[j]].right
			}
		}

		steps++
		i = (i + 1) % dir_len

		for _, current := range currents {
			if strings.HasSuffix(current, "Z") {
				solutions[current] = steps
			}
		}

		if len(solutions) == len(currents) {
			break
		}
	}

	return lcm(util.GetMapValues(solutions))
}

type value struct {
	left, right string
}

// RL

// AAA = (BBB, CCC)
// BBB = (DDD, EEE)
// CCC = (ZZZ, GGG)
// DDD = (DDD, DDD)
// EEE = (EEE, EEE)
// GGG = (GGG, GGG)
// ZZZ = (ZZZ, ZZZ)
func parseInput(input string) ([]string, map[string]value) {
	parsed := make(map[string]value)
	parts := strings.Split(input, "\n\n")

	for _, line := range strings.Split(parts[1], "\n") {
		parts := strings.Split(line, " = ")
		inst := strings.Split(strings.Trim(parts[1], "()"), ", ")

		parsed[parts[0]] = value{left: inst[0], right: inst[1]}
	}

	return strings.Split(parts[0], ""), parsed
}

// lcm (Least Common Multiple) of a slice of integers
func lcm(xs []int) int {
	ans := 1
	for _, x := range xs {
		ans = (x * ans) / gcd(x, ans)
	}
	return ans
}

// gcd (Greatest Common Divisor) using Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
