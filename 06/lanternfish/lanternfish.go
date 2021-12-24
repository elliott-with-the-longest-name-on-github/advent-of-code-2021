package lanternfish

import (
	"fmt"
	"strconv"
	"strings"
)

type School map[int]int

func PartOne(str string) string {
	s, err := SchoolFromString(str)
	if err != nil {
		panic(err)
	}
	s = s.IterateDays(80)
	return fmt.Sprintf("Part One: Fish multiplied to %v in 80 days!", s.NumFish())
}

func PartTwo(str string) string {
	s, err := SchoolFromString(str)
	if err != nil {
		panic(err)
	}
	s = s.IterateDays(256)
	return fmt.Sprintf("Part Two: Fish multiplied to %v in 256 days!", s.NumFish())
}

func SchoolFromString(str string) (School, error) {
	nums := strings.Split(str, ",")
	school := make(map[int]int, 9)
	for _, num := range nums {
		parsed, err := strconv.Atoi(num)
		if err != nil {
			return nil, fmt.Errorf("failed to parse input string. Wanted an int, got: %v", num)
		}
		school[parsed] += 1
	}
	return School(school), nil
}

func (s School) IterateDay() School {
	newSchool := make(School, 9)
	for k, v := range s {
		newKey := k - 1
		if newKey < 0 {
			newKey = 6
			newSchool[8] += v
		}
		newSchool[newKey] += v
	}
	return newSchool
}

func (s School) IterateDays(days int) School {
	newSchool := make(School, 9)
	for k, v := range s {
		newSchool[k] = v
	}
	for i := 0; i < days; i++ {
		newSchool = newSchool.IterateDay()
	}
	return newSchool
}

func (s School) NumFish() int {
	count := 0
	for _, v := range s {
		count += v
	}
	return count
}

var ChallengeInput string = "3,1,4,2,1,1,1,1,1,1,1,4,1,4,1,2,1,1,2,1,3,4,5,1,1,4,1,3,3,1,1,1,1,3,3,1,3,3,1,5,5,1,1,3,1,1,2,1,1,1,3,1,4,3,2,1,4,3,3,1,1,1,1,5,1,4,1,1,1,4,1,4,4,1,5,1,1,4,5,1,1,2,1,1,1,4,1,2,1,1,1,1,1,1,5,1,3,1,1,4,4,1,1,5,1,2,1,1,1,1,5,1,3,1,1,1,2,2,1,4,1,3,1,4,1,2,1,1,1,1,1,3,2,5,4,4,1,3,2,1,4,1,3,1,1,1,2,1,1,5,1,2,1,1,1,2,1,4,3,1,1,1,4,1,1,1,1,1,2,2,1,1,5,1,1,3,1,2,5,5,1,4,1,1,1,1,1,2,1,1,1,1,4,5,1,1,1,1,1,1,1,1,1,3,4,4,1,1,4,1,3,4,1,5,4,2,5,1,2,1,1,1,1,1,1,4,3,2,1,1,3,2,5,2,5,5,1,3,1,2,1,1,1,1,1,1,1,1,1,3,1,1,1,3,1,4,1,4,2,1,3,4,1,1,1,2,3,1,1,1,4,1,2,5,1,2,1,5,1,1,2,1,2,1,1,1,1,4,3,4,1,5,5,4,1,1,5,2,1,3"
