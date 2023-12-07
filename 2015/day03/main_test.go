package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []testCase{
		{
			name:  "example one",
			input: ">",
			want:  2,
		}, {
			name:  "example two",
			input: "^>v<",
			want:  4,
		}, {
			name:  "example three",
			input: "^v^v^v^v^v",
			want:  2,
		},
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  2081,
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
			name:  "example one",
			input: "^v",
			want:  3,
		}, {
			name:  "example two",
			input: "^>v<",
			want:  3,
		}, {
			name:  "example three",
			input: "^v^v^v^v^v",
			want:  11,
		},
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  2341,
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
