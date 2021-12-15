package main

import (
	"fmt"

	"example.com/aoc15/chiton"
)

func main() {
	path, err := chiton.FromLinesExtended(chiton.ChallengeLines, 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(path.Distance)
}
