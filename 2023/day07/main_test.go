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
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  6440,
		},
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
	tests := []struct {
		name  string
		input string
		want  int
	}{
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
