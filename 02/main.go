package main

import (
	"fmt"

	"github.com/tcc-sejohnson/advent-of-code-2021/02/submarine"
)

func main() {
	instructions := submarine.ChallengeInstructions
	sub := submarine.Submarine{}
	err := sub.Move(instructions...)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Submarine moved %v units.", sub.FinalDistance())
}
