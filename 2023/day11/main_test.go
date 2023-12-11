package main

import (
	"testing"
)

var example = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func Test_solve(t *testing.T) {
	tests := []testCase{
		{
			name:      "example with 1 expansion",
			input:     example,
			expand_by: 1,
			want:      374,
		},
		{
			name:      "example with 10 expansions",
			input:     example,
			expand_by: 9,
			want:      1030,
		},
		{
			name:      "example with 100 expansions",
			input:     example,
			expand_by: 100 - 1,
			want:      8410,
		},
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests,
			testCase{
				name:      "input",
				input:     input,
				expand_by: 1,
				want:      9543156,
			},
			testCase{
				name:      "input part 2",
				input:     input,
				expand_by: 1_000_000 - 1,
				want:      625243292686,
			},
		)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.input, tt.expand_by); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testCase struct {
	name      string
	input     string
	expand_by int
	want      int
}
