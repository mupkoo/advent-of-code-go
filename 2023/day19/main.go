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
	workflows, ratings := parseInput(input)
	sum := 0

	for _, rating := range ratings {
		current := "in"

		for !strings.Contains("AR", current) {
			for _, w := range workflows[current] {
				if w.condition.left == "TRUE" {
					current = w.result
					break
				} else if w.condition.op == ">" && rating[w.condition.left] > w.condition.right {
					current = w.result
					break
				} else if w.condition.op == "<" && rating[w.condition.left] < w.condition.right {
					current = w.result
					break
				}
			}
		}

		if current == "A" {
			for _, v := range rating {
				sum += v
			}
		}
	}

	return sum
}

func part2(input string) int {
	workflows, _ := parseInput(input)
	ratings := RatingRanges{
		"x": {1, 4000}, "m": {1, 4000}, "a": {1, 4000}, "s": {1, 4000},
	}

	return calculate(&workflows, ratings, "in")
}

type Condition struct {
	left  string
	op    string
	right int
}

type Workflow struct {
	condition Condition
	result    string
}

type RatingRanges map[string][2]int

func calculate(workflows *map[string][]Workflow, ratings RatingRanges, key string) int {
	if key == "A" {
		result := 1
		for _, v := range ratings {
			result *= v[1] - v[0] + 1
		}

		return result
	}

	if key == "R" {
		return 0
	}

	currentWorkflow := (*workflows)[key]
	result := 0

	for _, w := range currentWorkflow {
		if w.condition.left == "TRUE" {
			result += calculate(workflows, ratings, w.result)
		}

		// s>1200 | {0, 4000} -> {1201, 4000}
		if w.condition.op == ">" {
			next := [2]int{w.condition.right + 1, ratings[w.condition.left][1]}

			if next[0] < next[1] {
				copy := dup(ratings)
				copy[w.condition.left] = next
				result += calculate(workflows, copy, w.result)
			}

			failRange := [2]int{ratings[w.condition.left][0], w.condition.right}

			if failRange[0] <= failRange[1] {
				ratings[w.condition.left] = failRange
			}
		}

		// s<1200 | {0, 4000} -> {0, 1199}
		if w.condition.op == "<" {
			passRange := [2]int{ratings[w.condition.left][0], w.condition.right - 1}

			if passRange[0] < passRange[1] {
				copy := dup(ratings)
				copy[w.condition.left] = passRange
				result += calculate(workflows, copy, w.result)
			}

			failRange := [2]int{w.condition.right, ratings[w.condition.left][1]}

			if failRange[0] <= failRange[1] {
				ratings[w.condition.left] = failRange
			}
		}
	}

	return result
}

func dup(rating RatingRanges) RatingRanges {
	r := make(RatingRanges)
	for k, v := range rating {
		r[k] = v
	}
	return r
}

// qqz{s>2770:qs,m<1801:hdj,R}
// gd{a>3333:R,R}
// hdj{m>838:A,pv}
//
// {x=787,m=2655,a=1222,s=2876}
// {x=1679,m=44,a=2067,s=496}
func parseInput(input string) (workflows map[string][]Workflow, ratings []map[string]int) {
	parts := strings.Split(input, "\n\n")

	workflows = make(map[string][]Workflow)
	for _, rule := range strings.Split(parts[0], "\n") {
		parts := strings.Split(rule, "{")
		items := []Workflow{}

		for _, workflow := range strings.Split(strings.TrimRight(parts[1], "}"), ",") {
			segments := strings.Split(workflow, ":")

			if len(segments) == 1 {
				items = append(items, Workflow{
					Condition{left: "TRUE"},
					segments[0],
				})
			} else {
				runes := []rune(segments[0])

				items = append(items, Workflow{
					Condition{
						left:  string(runes[0]),
						op:    string(runes[1]),
						right: cast.ToInt(string(runes[2:])),
					},
					segments[1],
				})
			}

			workflows[parts[0]] = items
		}
	}

	for _, rating := range strings.Split(parts[1], "\n") {
		x, m, a, s := 0, 0, 0, 0
		fmt.Sscanf(rating, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)
		ratings = append(ratings, map[string]int{
			"x": x,
			"m": m,
			"a": a,
			"s": s,
		})
	}

	return workflows, ratings
}
