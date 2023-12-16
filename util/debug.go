package util

import (
	"fmt"

	"github.com/mupkoo/advent-of-code-go/cast"
)

func DebugGrid[T any](grid [][]T) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(cast.ToString(col))
		}
		println()
	}
	println()
}
