package chiton_test

import (
	"testing"

	"github.com/tcc-sejohnson/advent-of-code-2021/15/chiton"
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
	c, err := chiton.FromLines(lines)
	if err != nil {
		t.Errorf("failed to initialize ChitonMapper")
	}
	_, cost, err := c.BestPath()
	if err != nil {
		t.Errorf("failed to determine best path with error %v", err)
	}
	if cost != 40 {
		t.Errorf("incorrect distance for best path. Want: %v, got: %v", 40, cost)
	}
}

func TestExtendedPathDistance(t *testing.T) {
	c, err := chiton.FromLinesExtended(lines, 5)
	if err != nil {
		t.Errorf("failed to initialize ChitonMapper")
	}
	_, cost, err := c.BestPath()
	if err != nil {
		t.Errorf("failed to determine best path with error %v", err)
	}
	if cost != 315 {
		t.Errorf("incorrect distance for best path. Want: %v, got: %v", 315, cost)
	}
}
