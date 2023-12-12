package util

import (
	"reflect"
	"slices"
	"testing"
)

func TestGetMapValues(t *testing.T) {
	t.Run("map with int values", func(t *testing.T) {
		input := map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
		}

		want := []int{1, 2, 3}
		got := GetMapValues(input)
		slices.Sort(got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetMapValues(%v) = %v, want %v", input, got, want)
		}
	})

	t.Run("map with string values", func(t *testing.T) {
		input := map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
		}

		want := []string{"1", "2", "3"}
		got := GetMapValues(input)
		slices.Sort(got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetMapValues(%v) = %v, want %v", input, got, want)
		}
	})
}
