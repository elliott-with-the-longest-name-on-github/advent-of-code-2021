package main

import (
	"fmt"

	"github.com/tcc-sejohnson/advent-of-code-2021/13/paper"
)

func main() {
	verbose := false

	steps := []struct {
		Direction string
		Line      int
	}{
		{Direction: "left", Line: 655},
		{Direction: "up", Line: 447},
		{Direction: "left", Line: 327},
		{Direction: "up", Line: 223},
		{Direction: "left", Line: 163},
		{Direction: "up", Line: 111},
		{Direction: "left", Line: 81},
		{Direction: "up", Line: 55},
		{Direction: "left", Line: 40},
		{Direction: "up", Line: 27},
		{Direction: "up", Line: 13},
		{Direction: "up", Line: 6},
	}

	p := &paper.ChallengePaper
	if verbose {
		fmt.Println("Starting state:")
		p.Print()
		fmt.Println()
	}

	var err error
	for i, step := range steps {
		p, err = p.Fold(step.Direction, step.Line)
		if err != nil {
			fmt.Errorf("fold failed at step %v", i+1)
		}
		fmt.Printf("Dots visible on stage %v: %v\n", i+1, p.VisibleDots())
		if verbose {
			p.Print()
		}
	}

	p.Print()
}
