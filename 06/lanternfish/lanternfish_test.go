package lanternfish_test

import (
	"testing"

	"github.com/tcc-sejohnson/advent-of-code-2021/06/lanternfish"
)

func TestPartOne(t *testing.T) {
	got := lanternfish.PartOne("3,4,3,1,2")
	want := "Part One: Fish multiplied to 5934 in 80 days!"
	if want != got {
		t.Errorf("Incorrect number of fish for Part One.\nWant: %v\nGot: %v", want, got)
	}
}

func TestPartTwo(t *testing.T) {
	got := lanternfish.PartTwo("3,4,3,1,2")
	want := "Part Two: Fish multiplied to 26984457539 in 256 days!"
	if want != got {
		t.Errorf("Incorrect number of fish for Part Two.\nWant: %v\nGot: %v", want, got)
	}
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lanternfish.PartOne(lanternfish.ChallengeInput)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lanternfish.PartTwo(lanternfish.ChallengeInput)
	}
}
