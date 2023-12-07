package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []testCase{}

	for _, example := range []struct {
		want  int
		cases []string
	}{
		{0, []string{"(())", "()()"}},
		{3, []string{"(((", "(()(()(", "))((((("}},
		{-1, []string{"())", "))("}},
		{-3, []string{")))", ")())())"}},
	} {
		for _, input := range example.cases {
			tests = append(tests, testCase{
				name:  "example zero",
				input: input,
				want:  example.want,
			})
		}
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  232,
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
			name:  "example 1",
			input: ")",
			want:  1,
		},
		{
			name:  "example 2",
			input: "()())",
			want:  5,
		},
	}

	// Do not run this in CI
	if input != "blank" {
		tests = append(tests, testCase{
			name:  "input",
			input: input,
			want:  1783,
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
