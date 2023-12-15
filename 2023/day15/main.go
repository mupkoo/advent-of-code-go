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
	var ans int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans = part1(input)
	} else {
		ans = part2(input)
	}

	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	parsed := parseInput(input)
	sum := 0

	for _, part := range parsed {
		sum += hash(part)
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	store := make(map[string]int)
	boxes := make([][]string, 256)

	for _, part := range parsed {
		if strings.Contains(part, "-") {
			key := strings.Trim(part, "-")
			index := hash(key)
			keyIndex := slices.Index(boxes[index], key)

			if keyIndex > -1 {
				boxes[index] = append(boxes[index][:keyIndex], boxes[index][keyIndex+1:]...)
			}
		}

		if strings.Contains(part, "=") {
			parts := strings.Split(part, "=")
			key := parts[0]
			value := cast.ToInt(parts[1])
			index := hash(key)

			store[key] = value

			if slices.Index(boxes[index], key) == -1 {
				boxes[index] = append(boxes[index], key)
			}
		}
	}

	sum := 0

	// rn: 1 (box 0) * 1 (first slot) * 1 (focal length) = 1
	// cm: 1 (box 0) * 2 (second slot) * 2 (focal length) = 4
	// ot: 4 (box 3) * 1 (first slot) * 7 (focal length) = 28
	// ab: 4 (box 3) * 2 (second slot) * 5 (focal length) = 40
	// pc: 4 (box 3) * 3 (third slot) * 6 (focal length) = 72
	for i, box := range boxes {
		for j, key := range box {
			sum += (i + 1) * (j + 1) * store[key]
		}
	}

	return sum
}

func hash(input string) (result int) {
	for _, char := range input {
		result += cast.ToASCIICode(char)
		result *= 17
		result %= 256
	}

	return result
}

func parseInput(input string) []string {
	return strings.Split(input, ",")
}
