package main

import (
	"fmt"

	"github.com/tcc-sejohnson/advent-of-code-2021/03/diagnostic"
)

func main() {
	bin := diagnostic.ChallengeBinary
	log, err := diagnostic.FromBinaryStrings(bin, 12)
	if err != nil {
		panic("failed to parse binary input with error " + err.Error())
	}

	fmt.Printf("Power consumption: %v\n", log.PowerConsumption())
	fmt.Printf("Life support rating: %v\n", log.LifeSupportRating())
}
