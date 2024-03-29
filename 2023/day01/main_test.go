package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []testCase{
		{
			name: "example",
			input: `1asb2
			3asb4
			sa2b3ass4ss
			sda2das`,
			want: 12 + 34 + 24 + 22,
		},
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  55712,
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
			input: `two1nine
			eightwothree
			abcone2threexyz
			xtwone3four
			4nineeightseven2
			zoneight234
			7pqrstsixteen`,
			want: 281,
		},
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  55413,
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
