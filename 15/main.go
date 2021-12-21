package main

import (
	"fmt"

	"github.com/tcc-sejohnson/advent-of-code-2021/15/chiton"
)

func main() {
	c, err := chiton.FromLinesExtended(chiton.ChallengeLines, 5)
	if err != nil {
		panic(err)
	}
	_, cost, err := c.BestPath()
	if err != nil {
		panic(err)
	}
	fmt.Println(cost)
}
