package main

import (
	"testing"
)

var example = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func Test_part1(t *testing.T) {
	tests := []testCase{
		{
			name:  "example",
			input: example,
			want:  6440,
		},
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  250951660,
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
			want:  5905,
		},
		{
			name: "example2",
			input: `QQQQ2 20
			JTTT2 10
			JTJT1 5`,
			want: 85,
		},
		{
			name: "example3",
			input: `QQQQQ 10
			QQQQ2 9
			QQQ22 8
			QQ322 7
			QQ321 6
			Q4321 5`,
			want: 175,
		},
		{
			name: "example4",
			input: `QQQQQ 11
			QQJQQ 10
			QQQQ2 9
			QQQ22 8
			QQ322 7
			QQ321 6
			Q4321 5`,
			want: 252,
		},
	}

	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  251481660,
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
