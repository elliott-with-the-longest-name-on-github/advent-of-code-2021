package whales_test

import (
	"testing"

	"github.com/tcc-sejohnson/advent-of-code-2021/07/whales"
)

func TestMedian(t *testing.T) {
	inputEven := []int{1, 12, 123, 14, 15, 16, 9, 9}
	wantEven := []int{14, 15}
	inputOdd := []int{1, 12, 123, 15, 16, 9, 9}
	wantOdd := []int{15, 15}
	gotEvenOne, gotEvenTwo := whales.Median(inputEven)
	gotOddOne, gotOddTwo := whales.Median(inputOdd)
	if gotEvenOne != wantEven[0] || gotEvenTwo != wantEven[1] {
		t.Errorf("GetMedian failed on an even-length slice. Want: %v %v, got: %v %v", wantEven[0], wantEven[1], gotEvenOne, gotEvenTwo)
	}
	if gotOddOne != wantOdd[0] || gotOddTwo != wantOdd[1] {
		t.Errorf("GetMedian failed on an odd-length slice. Want: %v %v, got: %v %v", wantOdd[0], wantOdd[1], gotOddOne, gotOddTwo)
	}
}

func TestPartOne(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"
	want := "Part One: The crabs have aligned! It cost them... dearly. 37 fuel consumed."
	got := whales.PartOne(input)
	if want != got {
		t.Errorf("incorrect result from Part One.\nWant: %v\nGot: %v", want, got)
	}
}

func TestPartTwo(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"
	want := "Part Two: The crabs have aligned! It has cost them... most dearly. 168 fuel consumed."
	got := whales.PartTwo(input)
	if want != got {
		t.Errorf("incorrect result from Part Two.\nWant: %v\nGot:%v", want, got)
	}
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		whales.PartOne(whales.ChallengeInput)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		whales.PartTwo(whales.ChallengeInput)
	}
}
