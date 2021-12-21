package snailfish_test

import (
	"testing"

	"github.com/tcc-sejohnson/advent-of-code-2021/18/snailfish"
)

func TestSnailfishStringification(t *testing.T) {
	tests := []string{
		"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		"[[[[3,0],[5,3]],[4,4]],[5,5]]",
	}

	for _, test := range tests {
		num, err := snailfish.FromString(test)
		if err != nil {
			t.Errorf("failed to parse string %s with error %v", test, err)
		}
		result := num.String()
		if result != test {
			t.Errorf("failed to roundtrip a string. Started with: %s, ended with: %s", test, result)
		}
	}
}

func TestSnailfishAddition(t *testing.T) {
	tests := []struct {
		NumToAdd string
		Result   string
	}{
		{
			NumToAdd: "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			Result:   "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		},
		{
			NumToAdd: "[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
			Result:   "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
		},
		{
			NumToAdd: "[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
			Result:   "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
		},
		{
			NumToAdd: "[7,[5,[[3,8],[1,4]]]]",
			Result:   "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
		},
		{
			NumToAdd: "[[2,[2,2]],[8,[8,1]]]",
			Result:   "[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
		},
		{
			NumToAdd: "[2,9]",
			Result:   "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
		},
		{
			NumToAdd: "[1,[[[9,3],9],[[9,0],[0,7]]]]",
			Result:   "[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]",
		},
		{
			NumToAdd: "[[[5,[7,4]],7],1]",
			Result:   "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]",
		},
		{
			NumToAdd: "[[[[4,2],2],6],[8,7]]",
			Result:   "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
	}

	baseStr := "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]"
	baseNum, err := snailfish.FromString(baseStr)
	if err != nil {
		t.Errorf("failed to parse string %s", baseStr)
	}
	for _, test := range tests {
		numToAdd, err := snailfish.FromString(test.NumToAdd)
		if err != nil {
			t.Errorf("failed to parse string %s", test.NumToAdd)
		}
		baseNum = baseNum.Add(numToAdd)
		result := baseNum.String()
		if result != test.Result {
			t.Errorf("incorrect addition result. Want: %s, got: %s", test.Result, result)
		}
	}
}

func TestSnailfishMagnitude(t *testing.T) {
	tests := []struct {
		InputString string
		Want        int
	}{
		{
			InputString: "[[1,2],[[3,4],5]]",
			Want:        143,
		},
		{
			InputString: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			Want:        1384,
		},
		{
			InputString: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
			Want:        445,
		},
		{
			InputString: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
			Want:        791,
		},
		{
			InputString: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
			Want:        1137,
		},
		{
			InputString: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			Want:        3488,
		},
	}

	for _, test := range tests {
		num, err := snailfish.FromString(test.InputString)
		if err != nil {
			t.Errorf("failed to parse input string %s with error %v", test.InputString, err)
		}
		got := num.Magnitude()
		if got != test.Want {
			t.Errorf("magnitude incorrect. Want: %v, got: %v", test.Want, got)
		}
	}
}

func TestGreatestSumOfTwo(t *testing.T) {
	numbers := []string{
		"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
		"[[[5,[2,8]],4],[5,[[9,9],0]]]",
		"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
		"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
		"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
		"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
		"[[[[5,4],[7,7]],8],[[8,3],8]]",
		"[[9,3],[[9,9],[6,[4,9]]]]",
		"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
		"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
	}

	wantNum := "[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]"
	wantMag := 3993

	num, mag := snailfish.GreatestMagnitudeOfTwoStrings(numbers)

	numResult := num.String()
	if wantNum != numResult {
		t.Errorf("incorrect number result. Want: %s, got: %s", wantNum, numResult)
	}

	if wantMag != mag {
		t.Errorf("incorrect magnitude result. Want: %v, got: %v", wantMag, mag)
	}
}
