package paper_test

import (
	"fmt"
	"testing"

	"github.com/tcc-sejohnson/advent-of-code-2021/13/paper"
)

func TestVisibleDotsAfterFold(t *testing.T) {
	paper := paper.New(
		map[string]*paper.Dot{
			"6,10":  {X: 6, Y: 10},
			"0,14":  {X: 0, Y: 14},
			"9,10":  {X: 9, Y: 10},
			"0,3":   {X: 0, Y: 3},
			"10,4":  {X: 10, Y: 4},
			"4,11":  {X: 4, Y: 11},
			"6,0":   {X: 6, Y: 0},
			"6,12":  {X: 6, Y: 12},
			"4,1":   {X: 4, Y: 1},
			"0,13":  {X: 0, Y: 13},
			"10,12": {X: 10, Y: 12},
			"3,4":   {X: 3, Y: 4},
			"3,0":   {X: 3, Y: 0},
			"8,4":   {X: 8, Y: 4},
			"1,10":  {X: 1, Y: 10},
			"2,14":  {X: 2, Y: 14},
			"8,10":  {X: 8, Y: 10},
			"9,0":   {X: 9, Y: 0},
		},
		0,
		0,
	)

	fmt.Println("Paper starting state:")
	paper.Print()
	fmt.Print("\n")

	numDotsStart := paper.VisibleDots()
	if numDotsStart != 18 {
		t.Errorf("paper reports the wrong number of dots at the start. Wanted %v, got %v", 18, numDotsStart)
	}

	firstFold, err := paper.Fold("up", 7)
	if err != nil {
		t.Errorf("paper failed on the first fold with error %s", err)
	}

	fmt.Println("Paper state after first fold:")
	firstFold.Print()
	fmt.Print("\n")

	numDotsAfterFirstFold := firstFold.VisibleDots()
	if numDotsAfterFirstFold != 17 {
		t.Errorf("paper reports the wrong number of dots after the first fold. Wanted %v, got %v", 17, numDotsAfterFirstFold)
	}

	secondFold, err := firstFold.Fold("left", 5)
	if err != nil {
		t.Errorf("paper failed on the second fold with error %s", err)
	}

	fmt.Println("Paper state after second fold:")
	secondFold.Print()
	fmt.Print("\n")

	numDotsAfterSecondFold := secondFold.VisibleDots()
	if numDotsAfterSecondFold != 16 {
		t.Errorf("paper reports the wrong number of dots after the second fold. Wanted %v, got %v", 18, numDotsAfterSecondFold)
	}
}
