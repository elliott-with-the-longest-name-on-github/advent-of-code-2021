package main

import (
	"fmt"

	"example.com/aoc02/submarine"
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
