package main

import (
	"fmt"

	"github.com/tcc-sejohnson/advent-of-code-2021/09/basin"
)

func main() {
	fmt.Println(basin.PartOne(basin.ChallengeInput))
	fmt.Println(basin.PartTwo(basin.ChallengeInput))
	// VisualizeSmall()
	// VisualizeBig()
}

func VisualizeSmall() {
	fmt.Println(basin.FunPartTwo([]string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}))
}

func VisualizeBig() {
	fmt.Println(basin.FunPartTwo(basin.ChallengeInput))
}
