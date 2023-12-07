package main

import (
	"crypto/md5"
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
	return calculate(input, "00000")
}

func part2(input string) int {
	return calculate(input, "000000")
}

func calculate(input, prefix string) (answer int) {
	for {
		answer++
		hash := md5.Sum([]byte(input + cast.ToString(answer)))

		if strings.HasPrefix(fmt.Sprintf("%x", hash), prefix) {
			break
		}
	}

	return answer
}
