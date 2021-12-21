package enhance_test

import (
	"testing"

	"github.com/tcc-sejohnson/advent-of-code-2021/20/enhance"
)

// This is really gross, but it's how the input is.
var testAlg string = "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#"
var testImg []string = []string{
	"#..#.",
	"#....",
	"##..#",
	"..#..",
	"..###",
}

func TestStringRoundtrip(t *testing.T) {
	wantStr := "#..#.\n#....\n##..#\n..#..\n..###\n"
	im, err := enhance.FromStrings(testImg, testAlg)
	if err != nil {
		t.Errorf("failed to parse testImg with error %v", err)
	}
	gotStr := im.String()
	if wantStr != gotStr {
		t.Errorf("string roundtrip failed. Want: %v, got: %v", wantStr, gotStr)
	}
	gotAlg := im.EnhancementAlgorithmToString()
	if gotAlg != testAlg {
		t.Errorf("string roundtrip failed. Want: %v, got: %v", testAlg, gotAlg)
	}
}

func TestLitPixels2x(t *testing.T) {
	want := 35
	im, err := enhance.FromStrings(testImg, testAlg)
	if err != nil {
		t.Errorf("failed to parse testImg with error %v", err)
	}
	got := im.Enhance().Enhance().LitPixelCount()
	if want != got {
		t.Errorf("incorrect lit pixel count reported. Want: %v, got: %v", want, got)
	}
}

func TestLitPixels50x(t *testing.T) {
	want := 3351
	im, err := enhance.FromStrings(testImg, testAlg)
	if err != nil {
		t.Errorf("failed to parse testImg with error %v", err)
	}
	for i := 0; i < 50; i++ {
		im = im.Enhance()
	}
	got := im.LitPixelCount()
	if want != got {
		t.Errorf("incorrect lit pixel count reported. Want: %v, got: %v", want, got)
	}
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		enhance.PartOne()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		enhance.PartTwo()
	}
}
