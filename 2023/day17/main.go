package main

import (
	"container/heap"
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/mupkoo/advent-of-code-go/cast"
	"github.com/mupkoo/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	var ans int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans = part1(input)
	} else {
		ans = part2(input)
	}

	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	grid := parseInput(input)
	return leastHeatLost(grid, 1, 3)
}

func part2(input string) int {
	grid := parseInput(input)
	return leastHeatLost(grid, 4, 10)
}

func parseInput(input string) (grid [][]int) {
	for row, line := range strings.Split(input, "\n") {
		grid = append(grid, make([]int, len(line)))
		for col, char := range line {
			grid[row][col] = cast.ToInt(char)
		}
	}

	return grid
}

func leastHeatLost(grid [][]int, minMovement, maxMovement int) int {
	initialDir := Dir{0, 0}
	pq := &PriorityQueue{
		{&Value{0, 0, initialDir, 0}, 0, 0},
	}
	heap.Init(pq)

	seen := make(map[Value]bool)

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)

		if item.value.row == len(grid)-1 && item.value.col == len(grid[0])-1 && item.value.dirCount >= minMovement {
			return item.heatLost
		}

		if seen[*item.value] {
			continue
		}
		seen[*item.value] = true

		if item.value.dirCount < maxMovement && item.value.dir != initialDir {
			nextRow := item.value.row + item.value.dir.row
			nextCol := item.value.col + item.value.dir.col

			if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) {
				heap.Push(pq, &Item{
					value:    &Value{nextRow, nextCol, item.value.dir, item.value.dirCount + 1},
					heatLost: item.heatLost + grid[nextRow][nextCol],
					index:    0,
				})
			}
		}

		// We need to move in a direction at least `minMovement` times before we can change
		if item.value.dirCount < minMovement && item.value.dir != initialDir {
			continue
		}

		for _, dir := range []Dir{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			if item.value.dir == dir {
				continue
			}

			if item.value.dir == (Dir{-dir.row, -dir.col}) {
				continue
			}

			nextRow := item.value.row + dir.row
			nextCol := item.value.col + dir.col

			if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) {
				heap.Push(pq, &Item{
					value:    &Value{nextRow, nextCol, dir, 1},
					heatLost: item.heatLost + grid[nextRow][nextCol],
					index:    0,
				})
			}
		}
	}

	return 0
}

type Value struct {
	row, col int
	dir      Dir
	dirCount int
}

type Dir struct {
	row, col int
}

type Item struct {
	value    *Value
	heatLost int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heatLost < pq[j].heatLost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	item.index = -1
	*pq = (*pq)[:n-1]

	return item
}
