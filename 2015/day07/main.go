package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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

	return resolve("a", parsed)
}

func part2(input string) int {
	parsed := parseInput(input)
	copy := make(map[string]value)
	for k, v := range parsed {
		copy[k] = v
	}

	a := resolve("a", parsed)
	copy["b"] = value{value: a, processed: true}

	return resolve("a", copy)
}

type value struct {
	raw       string
	value     int
	processed bool
}

// 123 -> x
// 456 -> y
// x AND y -> d
// x OR y -> e
// x LSHIFT 2 -> f
// y RSHIFT 2 -> g
// NOT x -> h
// NOT y -> i
func parseInput(input string) map[string]value {
	parsed := make(map[string]value)

	for _, line := range strings.Split(input, "\n") {
		segments := strings.Split(line, " -> ")
		parsed[segments[1]] = value{raw: segments[0]}
	}

	return parsed
}

func resolve(point string, instructions map[string]value) int {
	val, ok := instructions[point]

	if !ok {
		return cast.ToInt(point)
	}

	if val.processed {
		return val.value
	}

	segments := strings.Split(val.raw, " ")

	if len(segments) == 1 {
		if value, err := strconv.Atoi(segments[0]); err == nil {
			val.value = value
		} else {
			val.value = resolve(segments[0], instructions)
		}
	} else if len(segments) == 2 && segments[0] == "NOT" {
		val.value = ^resolve(segments[1], instructions)
	} else {
		a := resolve(segments[0], instructions)
		b := resolve(segments[2], instructions)

		switch segments[1] {
		case "AND":
			val.value = a & b
		case "OR":
			val.value = a | b
		case "LSHIFT":
			val.value = a << b
		case "RSHIFT":
			val.value = a >> b
		}
	}

	val.processed = true
	instructions[point] = val

	return val.value
}
