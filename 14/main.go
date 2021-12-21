package main

import (
	"fmt"

	"github.com/tcc-sejohnson/advent-of-code-2021/14/polymerization"
)

func main() {
	rules := polymerization.ChallengeRules
	base := polymerization.ChallengeBase
	template := polymerization.New(base, rules)
	template.SubstituteMany(40)
	mostFrequentElement, mostFrequentCount := template.MostFrequentElement()
	leastFrequentElement, leastFrequentCount := template.LeastFrequentElement()
	fmt.Printf("Most frequent element, with %v occurences: %v\n", mostFrequentCount, string(mostFrequentElement))
	fmt.Printf("Least frequent element, with %v occurences: %v\n", leastFrequentCount, string(leastFrequentElement))
	fmt.Printf("%v - %v = %v", mostFrequentCount, leastFrequentCount, mostFrequentCount-leastFrequentCount)
}
