package submarine_test

import (
	"testing"

	"example.com/aoc02/submarine"
)

func TestSubmarineMovement(t *testing.T) {
	instructions := []submarine.Instruction{
		{Direction: "forward", Value: 5},
		{Direction: "down", Value: 5},
		{Direction: "forward", Value: 8},
		{Direction: "up", Value: 3},
		{Direction: "down", Value: 8},
		{Direction: "forward", Value: 2},
	}

	sub := submarine.Submarine{}
	err := sub.Move(instructions...)
	if err != nil {
		t.Errorf("submarine movement failed with error: %e", err)
	}

	finalDistance := sub.FinalDistance()
	if finalDistance != 900 {
		t.Errorf("submarine final distance is incorrect. Want: %v, got: %v", 900, finalDistance)
	}
}