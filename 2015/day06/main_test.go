package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []testCase{
		{
			name: "example",
			input: `turn on 0,0 through 999,999
			toggle 0,0 through 999,0
			turn off 499,499 through 500,500`,
			want: 1000*1000 - 1000 - 4,
		},
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  400410,
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
			name: "example",
			input: `turn on 0,0 through 0,0
			toggle 0,0 through 999,999`,
			want: 1 + 2000000,
		},
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  15343601,
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
