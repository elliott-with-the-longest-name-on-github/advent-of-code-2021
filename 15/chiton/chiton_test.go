package chiton_test

import (
	"fmt"
	"testing"

	"example.com/aoc15/chiton"
)

var lines []string = []string{
	"1163751742",
	"1381373672",
	"2136511328",
	"3694931569",
	"7463417111",
	"1319128137",
	"1359912421",
	"3125421639",
	"1293138521",
	"2311944581",
}

func TestPathDistance(t *testing.T) {
	best, err := chiton.FromLines(lines)
	if err != nil {
		t.Errorf("failed to determine best path with error %v", err)
	}
	if best.Distance != 40 {
		t.Errorf("incorrect distance for best path. Want: %v, got: %v", 40, best.Distance)
	}
}

func TestExtendedPathDistance(t *testing.T) {
	best, err := chiton.FromLinesExtended(lines, 5)
	if err != nil {
		t.Errorf("failed to determine best path with error %v", err)
	}
	if best.Distance != 315 {
		fmt.Println(best.Path)
		t.Errorf("incorrect distance for best path. Want: %v, got: %v", 315, best.Distance)
	}
}
