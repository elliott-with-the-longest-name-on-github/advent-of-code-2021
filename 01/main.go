package main

import (
	"fmt"

	"example.com/aoc01/depth"
)

func main() {
	m := depth.ChallengeMeasurements
	series := depth.New(m)
	c := 0
	for _, measurement := range series.WindowedMeasurements {
		fmt.Println(measurement)
		if measurement.PriorWindowedDepth >= 0 && measurement.PriorWindowedDepth < measurement.WindowedDepth {
			c++
		}
	}
	fmt.Printf("Number of increasing measurements: %v\n", c)
}
