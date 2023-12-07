package main

import (
	"testing"
)

var example = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func Test_part1(t *testing.T) {
	tests := []testCase{
		{
			name:  "example",
			input: example,
			want:  4361,
		},
		{
			name: "number at the end of line",
			input: example + `
			......#...
			.......100`,
			want: 4461,
		},
		{
			name: "number from start to finish",
			input: example + `
			......#...
			1000000000`,
			want: 1000004361,
		},
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  522726,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []testCase{
		{
			name:  "example",
			input: example,
			want:  467835,
		},
		{
			name: "number at the end of line",
			input: example + `
			......#...
			.......100`,
			want: 467835,
		},
		{
			name: "number from start to finish",
			input: example + `
			......#...
			1000000000`,
			want: 467835,
		},
		{
			name: "with numbers in one line",
			input: example + `
			......#...
			10*10.....`,
			want: 467935,
		},
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  81721933,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testCase struct {
	name  string
	input string
	want  int
}
