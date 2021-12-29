package basin_test

import (
	"testing"

	"github.com/tcc-sejohnson/advent-of-code-2021/09/basin"
)

func TestPartOne(t *testing.T) {
	input := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}
	got := basin.PartOne(input)
	want := "Part One: Found 4 low points with a total risk of 15."
	if want != got {
		t.Errorf("Part One output incorrect.\nWant: %v\nGot: %v\n", want, got)
	}
}

func TestPartTwo(t *testing.T) {
	input := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}
	got := basin.PartTwo(input)
	want := "Part Two: Top 3 basins' total area multiplied: 1134"
	if want != got {
		t.Errorf("Part Two output incorrect.\nWant: %v\nGot: %v\n", want, got)
	}
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basin.PartOne(basin.ChallengeInput)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basin.PartTwo(basin.ChallengeInput)
	}
}
