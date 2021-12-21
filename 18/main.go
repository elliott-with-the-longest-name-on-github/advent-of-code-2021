package main

import (
	"fmt"

	"github.com/tcc-sejohnson/advent-of-code-2021/18/snailfish"
)

func PartOne() {
	input := snailfish.ChallengeInput
	num := snailfish.SumStrings(input)
	fmt.Printf("Finished with SnailfishNumber %v of magnitude %v\n", num, num.Magnitude())
}

func PartTwo() {
	input := snailfish.ChallengeInput
	num, mag := snailfish.GreatestMagnitudeOfTwoStrings(input)
	fmt.Printf("Finished with SnailfishNumber %v of magnitude %v\n", num, mag)
}

func main() {
	PartOne()
	PartTwo()
}
