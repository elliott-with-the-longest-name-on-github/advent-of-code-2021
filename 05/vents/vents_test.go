package vents_test

import (
	"fmt"
	"testing"

	"github.com/tcc-sejohnson/advent-of-code-2021/05/vents"
)

func TestCountFinder(t *testing.T) {
	testStrings := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	m, err := vents.NewVentMapFromStrings(testStrings, true)
	if err != nil {
		t.Errorf("failed to parse a string with error: %v", err)
	}
	fmt.Print(m.String())
	points, err := m.FindPointsWhereCountGreaterThan(1)
	if err != nil {
		t.Errorf("%v", err)
	}

	if len(points) != 12 {
		t.Errorf("number of points with a count greater than 1 was incorrect. Wanted: 12, got: %v", len(points))
	}
}

func TestFromStrings(t *testing.T) {
	testStrings := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	m, err := vents.NewVentMapFromStrings(testStrings, true)
	if err != nil {
		t.Errorf("failed to parse a string with error: %v", err)
	}

	testPoint1 := &vents.Point{X: 3, Y: 4}
	testPoint1Count := m.GetPointCount(testPoint1)
	if testPoint1Count != 2 {
		t.Errorf("got incorrect point count for point %v. Wanted: %v, got: %v", testPoint1, 2, testPoint1Count)
	}
	testPoint2 := &vents.Point{X: 0, Y: 9}
	testPoint2Count := m.GetPointCount(testPoint2)
	if testPoint2Count != 2 {
		t.Errorf("got incorrect point count for point %v. Wanted: %v, got: %v", testPoint2, 2, testPoint2Count)
	}
}

func TestAddPoint(t *testing.T) {
	testPoints := []*vents.Point{
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
	}
	m := vents.NewVentMap()

	for i := 0; i < 10; i++ {
		m.AddPoint(testPoints[0])
	}

	for i := 0; i < 5; i++ {
		m.AddPoint(testPoints[1])
	}

	for i := 0; i < 2; i++ {
		m.AddPoint(testPoints[2])
	}

	p1Count := m.GetPointCount(testPoints[0])
	p2Count := m.GetPointCount(testPoints[1])
	p3Count := m.GetPointCount(testPoints[2])

	if p1Count != 10 {
		t.Errorf("incorrect point count. Wanted: %v, got: %v for point %v", 10, p1Count, testPoints[0])
	}
	if p2Count != 5 {
		t.Errorf("incorrect point count. Wanted: %v, got: %v for point %v", 5, p2Count, testPoints[1])
	}
	if p3Count != 2 {
		t.Errorf("incorrect point count. Wanted: %v, got: %v for point %v", 2, p3Count, testPoints[2])
	}

	p4Count := m.GetPointCount(testPoints[3])
	if p4Count != 0 {
		t.Errorf("incorrect point count. Wanted: %v, got: %v for point %v", 0, p4Count, testPoints[3])
	}
}

func TestLineStringRoundtripping(t *testing.T) {
	testCases := []string{
		"1,2 -> 2,1",
		"1,3 -> 3,1",
		"1234,1234 -> 12,1",
	}

	for _, tC := range testCases {
		line, err := vents.LineFromString(tC)
		if err != nil {
			t.Errorf("failed to parse string into line: %s", tC)
		}
		str := line.String()
		if str != tC {
			t.Errorf("failed to roundtrip string. Started with: %s, finished with: %s", tC, str)
		}
	}
}

func TestLineStringRoundtrippingErrors(t *testing.T) {
	testCases := []string{
		",2 -> 1,2",
		"1,2->1,2",
		"1,2 - 1,2",
		"a,b -> 1,2",
		"1,2 -> a,2",
		"1,b -> 1,2",
		"1,2 -> ",
	}

	for _, tC := range testCases {
		line, err := vents.LineFromString(tC)
		if err == nil {
			t.Errorf("should have failed on input string %s, but succeeded by parsing line %v instead", tC, line)
		}
	}
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vents.PartOne()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vents.PartTwo()
	}
}
