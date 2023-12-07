package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []testCase{
		{
			name: "example",
			input: `ugknbfddgicrmopn
			aaa
			jchzalrnumimnmhp
			haegwjzuvuyypxyu
			dvszwmarrgswjxmb`,
			want: 2,
		},
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  258,
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
			input: `qjhvhtzxzqqjkmpb
			xxyxx
			xxxyx
			uurcxstgmygtbstg
			ieodomkazucvgmuy
			zexpwallpocrxpvp
			tjpkkmcqbbkxaiak
			tftftwedtecglavz`,
			want: 4,
		},
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  53,
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
