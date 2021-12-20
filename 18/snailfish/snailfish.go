package snailfish

import (
	"fmt"
	"strconv"
)

const SPLIT_CUTOFF = 10

type SnailfishNumber struct {
	Parent   *SnailfishNumber
	Left     *SnailfishNumber
	Right    *SnailfishNumber
	LeftNum  int
	RightNum int
}

func (num *SnailfishNumber) Add(right *SnailfishNumber) *SnailfishNumber {
	newNumber := &SnailfishNumber{
		Parent:   nil,
		Left:     num,
		Right:    right,
		LeftNum:  -1,
		RightNum: -1,
	}
	newNumber.Left.Parent = newNumber
	newNumber.Right.Parent = newNumber

	return newNumber.Reduce()
}

func Sum(numbers []*SnailfishNumber) *SnailfishNumber {
	if len(numbers) == 0 {
		panic("numbers must be a valid, non-empty array")
	}
	baseNumber := numbers[0]
	for i := 1; i < len(numbers); i++ {
		baseNumber = baseNumber.Add(numbers[i])
	}
	return baseNumber
}

func SumStrings(strings []string) *SnailfishNumber {
	if len(strings) == 0 {
		panic("strings must be a valid, non-empty array")
	}
	numbers := make([]*SnailfishNumber, len(strings))
	for i, str := range strings {
		num, err := FromString(str)
		if err != nil {
			panic(err)
		}
		numbers[i] = num
	}
	return Sum(numbers)
}

// GreatestMagnitudeOfTwo calculates the greatest magnitude achievable by adding only two
// numbers out of a list of numbers.
// This could be greatly optimized, but I'm really tired and it's after 1AM.
func GreatestMagnitudeOfTwo(numbers []*SnailfishNumber) (*SnailfishNumber, int) {
	if len(numbers) == 0 {
		panic("numbers must be a valid, non-empty array")
	}
	greatestMagnitude := 0
	var greatestNumber *SnailfishNumber
	for _, leftNum := range numbers {
		for _, rightNum := range numbers {
			if leftNum == rightNum {
				// I'm assuming we can't add a number to itself...?
				continue
			}
			// var leftCopy *SnailfishNumber
			// var rightCopy *SnailfishNumber
			// need to write a DeepCopy function for a SnailfishNumber, otherwise
			// this is never going to work.
			sum := leftNum.Add(rightNum)
			mag := sum.Magnitude()
			if mag > greatestMagnitude {
				greatestMagnitude = mag
				greatestNumber = sum
			}
		}
	}
	return greatestNumber, greatestMagnitude
}

func GreatestMagnitudeOfTwoStrings(strings []string) (*SnailfishNumber, int) {
	if len(strings) == 0 {
		panic("strings must be a valid, non-empty array")
	}
	numbers := make([]*SnailfishNumber, len(strings))
	for i, str := range strings {
		num, err := FromString(str)
		if err != nil {
			panic(err)
		}
		numbers[i] = num
	}
	return GreatestMagnitudeOfTwo(numbers)
}

// Reduce simplifies a SnailfishNumber by applying the following two rules in order from left to right:
// 1. If a node is at the fourth level, explode it by adding its left number to the next number to the left of it (if any).
//    then replace the node with the number 0.
// 2. If a regular number is >10, replace it with a pair. The left number in the pair is the number / 2 rounded down,
//    the right number in the pair is the number / 2 rounded up.
func (num *SnailfishNumber) Reduce() *SnailfishNumber {
	numToExplode := num.getFirstNumAtDepth(4)
	if numToExplode != nil {
		numToExplode.explode()
		num.Reduce()
	} else if num.split(SPLIT_CUTOFF) {
		num.Reduce()
	}
	return num
}

func (num *SnailfishNumber) Magnitude() int {
	leftIsLeaf := num.Left == nil
	rightIsLeaf := num.Right == nil
	if leftIsLeaf && rightIsLeaf {
		return (num.LeftNum * 3) + (num.RightNum * 2)
	}
	if leftIsLeaf {
		return (num.LeftNum * 3) + (num.Right.Magnitude() * 2)
	}
	if rightIsLeaf {
		return (num.Left.Magnitude() * 3) + (num.RightNum * 2)
	}
	return (num.Left.Magnitude() * 3) + (num.Right.Magnitude() * 2)
}

func (num *SnailfishNumber) String() string {
	leftIsLeaf := num.Left == nil
	rightIsLeaf := num.Right == nil
	templateString := "[%v,%v]"
	if leftIsLeaf && rightIsLeaf {
		return fmt.Sprintf(templateString, num.LeftNum, num.RightNum)
	}
	if leftIsLeaf {
		return fmt.Sprintf(templateString, num.LeftNum, num.Right.String())
	}
	if rightIsLeaf {
		return fmt.Sprintf(templateString, num.Left.String(), num.RightNum)
	}
	return fmt.Sprintf(templateString, num.Left.String(), num.Right.String())
}

// FromRunes parses a string and constructs a SnailfishNumber.
func FromString(str string) (*SnailfishNumber, error) {
	runes := []rune(str)
	return FromRunes(runes, nil)
}

// FromRunes parses a rune array and constructs a SnailfishNumber. Listen, I did my best on error handling here.
// I do not have infinite time. Just don't feed it bad numbers and we can all be happy.
func FromRunes(runes []rune, parent *SnailfishNumber) (*SnailfishNumber, error) {
	var num *SnailfishNumber
	if parent == nil {
		num = &SnailfishNumber{Parent: nil}
	} else {
		num = &SnailfishNumber{Parent: parent}
	}

	if runes[0] != '[' {
		return nil, fmt.Errorf("malformed string provided to FromString. Must start with an opening bracket ([). Got: %s", string(runes[0]))
	}

	runesLen := len(runes)

	_, _, err := findRightBracket(runes)
	if err != nil {
		return nil, err
	}

	leftRunesRightBracketIndex := -1
	if runes[1] != '[' {
		// Must be an integer (if the string is valid)
		leafNum, err := strconv.Atoi(string(runes[1]))
		if err != nil {
			return nil, fmt.Errorf("found a non-int character where an integer was expected. Got: %s", string(runes[1]))
		}
		num.Left = nil
		num.LeftNum = leafNum
		leftRunesRightBracketIndex = 1
	} else {
		// the left number must be another pair rather than a leaf, so find the end of the pair and recurse
		var leftRunes []rune
		var err error
		leftRunes, leftRunesRightBracketIndex, err = findRightBracket(runes[1:runesLen])
		if err != nil {
			return nil, err
		}
		childNum, err := FromRunes(leftRunes, num)
		if err != nil {
			return nil, err
		}
		num.Left = childNum
		num.LeftNum = -1
	}

	// It's possible that we're missing a bracket
	if leftRunesRightBracketIndex+1 >= len(runes) {
		return nil, fmt.Errorf("missing a bracket in %s", string(runes))
	}
	// Make sure there's a comma between left and right
	if runes[leftRunesRightBracketIndex+1] != ',' {
		return nil, fmt.Errorf("missing comma between left and right elements in SnailfishNumber %s", string(runes))
	}

	rightRunesStartIndex := leftRunesRightBracketIndex + 2
	if runes[rightRunesStartIndex] != '[' {
		// Must be an integer (if the string is valid)
		if runes[rightRunesStartIndex+1] != ']' {
			return nil, fmt.Errorf("missing a right bracket in number %s", string(runes))
		}
		leafNum, err := strconv.Atoi(string(runes[rightRunesStartIndex]))
		if err != nil {
			return nil, fmt.Errorf("found a non-int character where an integer was expected. Got: %s", string(runes[rightRunesStartIndex]))
		}
		num.Right = nil
		num.RightNum = leafNum
	} else {
		// the left number must be another pair rather than a leaf, so find the end of the pair and recurse
		var rightRunes []rune
		var err error
		rightRunes, _, err = findRightBracket(runes[rightRunesStartIndex:runesLen])
		if err != nil {
			return nil, err
		}
		childNum, err := FromRunes(rightRunes, num)
		if err != nil {
			return nil, err
		}
		num.Right = childNum
		num.RightNum = -1
	}

	return num, nil
}

// getFirstNumAtDepth finds the leftmost number with a depth equal to the specified depth.
func (num *SnailfishNumber) getFirstNumAtDepth(depth int) *SnailfishNumber {
	currentDepth := 0
	return num.getFirstNumAtDepthRecursiveHelper(depth, currentDepth)
}

func (num *SnailfishNumber) getFirstNumAtDepthRecursiveHelper(stopDepth int, currentDepth int) *SnailfishNumber {
	if currentDepth == stopDepth {
		return num
	}
	depth := currentDepth + 1
	var foundNumber *SnailfishNumber = nil
	if num.Left != nil {
		foundNumber = num.Left.getFirstNumAtDepthRecursiveHelper(stopDepth, depth)
	}
	if foundNumber == nil && num.Right != nil {
		foundNumber = num.Right.getFirstNumAtDepthRecursiveHelper(stopDepth, depth)
	}
	return foundNumber
}

// explode adds this SnailfishNumber's left number to the nearest left number above it in the tree,
// and its right number to the nearest right number above it in the tree. It then "self destructs" by replacing itself with a 0 in its parent.
func (num *SnailfishNumber) explode() {
	num.recursivelyAddNumToParentsLeft(num.LeftNum)
	num.recursivelyAddNumToParentsRight(num.RightNum)
	if num.Parent.Left == num {
		num.Parent.Left = nil
		num.Parent.LeftNum = 0
	} else {
		num.Parent.Right = nil
		num.Parent.RightNum = 0
	}
}

// split finds the first leaf greater than or equal to cutoff and replaces it with a pair of numbers where
// the left number is 1/2 the number being replaced rounded down, and the right number is 1/2 the number being replaced
// rounded up.
func (num *SnailfishNumber) split(cutoff int) (tookAction bool) {
	if num.Left == nil {
		if num.LeftNum >= cutoff {
			num.Left = num.calculateSplitValue(num.LeftNum)
			num.LeftNum = -1
			tookAction = true
		}
	}
	if num.Left != nil && !tookAction {
		tookAction = num.Left.split(cutoff)
	}
	if num.Right == nil && !tookAction {
		if num.RightNum >= cutoff {
			num.Right = num.calculateSplitValue(num.RightNum)
			num.RightNum = -1
			tookAction = true
		}
	}
	if num.Right != nil && !tookAction {
		tookAction = num.Right.split(cutoff)
	}
	return
}

func (num *SnailfishNumber) calculateSplitValue(number int) *SnailfishNumber {
	// integer truncation helps here
	halfRoundedDown := number / 2
	newNum := &SnailfishNumber{
		Parent: num,
		Left:   nil,
		Right:  nil,
	}
	if number%2 == 1 {
		// odd
		newNum.LeftNum = halfRoundedDown
		newNum.RightNum = halfRoundedDown + 1
	} else {
		// even
		newNum.LeftNum = halfRoundedDown
		newNum.RightNum = halfRoundedDown
	}
	return newNum
}

func (num *SnailfishNumber) recursivelyAddNumToParentsLeft(number int) {
	if num.Parent != nil {
		if num.Parent.Left == nil {
			// we're on the right and our parent's left element is a regular number
			num.Parent.LeftNum += number
		} else if num.Parent.Left != num {
			// we're on the right, and we can guarantee some child of our parent's left has a right number,
			// which will be the closest to our left number.
			num.Parent.Left.recursivelyAddNumToChildsRight(number)
		} else {
			// we're on the left and we need to look to our parent's parent (and so on)
			num.Parent.recursivelyAddNumToParentsLeft(number)
		}
	}
}

func (num *SnailfishNumber) recursivelyAddNumToChildsRight(number int) {
	if num.Right == nil {
		num.RightNum += number
	} else {
		num.Right.recursivelyAddNumToChildsRight(number)
	}
}

func (num *SnailfishNumber) recursivelyAddNumToParentsRight(number int) {
	if num.Parent != nil {
		if num.Parent.Right == nil {
			// we're on the left and our parent's right element is a regular number
			num.Parent.RightNum += number
		} else if num.Parent.Right != num {
			// we're on the left, and we can guarantee some child of our parent's right has a left number,
			// which will be the closest to our right number.
			num.Parent.Right.recursivelyAddNumToChildsLeft(number)
		} else {
			// we're on the left and we need to look to our parent's parent (and so on)
			num.Parent.recursivelyAddNumToParentsRight(number)
		}
	}
}

func (num *SnailfishNumber) recursivelyAddNumToChildsLeft(number int) {
	if num.Left == nil {
		num.LeftNum += number
	} else {
		num.Left.recursivelyAddNumToChildsLeft(number)
	}
}

// findRightBracket searches a slice of runes for a matching right bracket
// and returns the slice of runes representing the start of the slice to the
// enclosing bracket. It also returns the index at which the right bracket was found.
func findRightBracket(runes []rune) ([]rune, int, error) {
	leftBracketCount := 1
	rightBracketCount := 0
	var leftRunes []rune = nil
	rightBracketIndex := -1
	for i := 1; i < len(runes); i++ {
		if runes[i] == '[' {
			leftBracketCount += 1
		} else if runes[i] == ']' {
			rightBracketCount += 1
		}

		if leftBracketCount == rightBracketCount {
			leftRunes = runes[0 : i+1]
			rightBracketIndex = i + 1
			break
		}
	}
	return leftRunes, rightBracketIndex, nil
}

var ChallengeInput []string = []string{
	"[[6,[[9,4],[5,5]]],[[[0,7],[7,8]],[7,0]]]",
	"[[[[2,1],[8,6]],[2,[4,0]]],[9,[4,[0,6]]]]",
	"[[[[4,2],[7,7]],4],[3,5]]",
	"[8,[3,[[2,3],5]]]",
	"[[[[0,0],[4,7]],[[5,5],[8,5]]],[8,0]]",
	"[[[[5,2],[5,7]],[1,[5,3]]],[[4,[8,4]],2]]",
	"[[5,[[2,8],[9,3]]],[[7,[5,2]],[[9,0],[5,2]]]]",
	"[[9,[[4,3],1]],[[[9,0],[5,8]],[[2,6],1]]]",
	"[[0,6],[6,[[6,4],[7,0]]]]",
	"[[[9,[4,2]],[[6,0],[8,9]]],[[0,4],[3,[6,8]]]]",
	"[[[[3,2],0],[[9,6],[3,1]]],[[[3,6],[7,6]],[2,[6,4]]]]",
	"[5,[[[1,6],[7,8]],[[6,1],[3,0]]]]",
	"[2,[[6,[7,6]],[[8,6],3]]]",
	"[[[[0,9],1],[2,3]],[[[7,9],1],7]]",
	"[[[[1,8],3],[[8,8],[0,8]]],[[2,1],[8,0]]]",
	"[[2,9],[[5,1],[[9,3],[4,0]]]]",
	"[9,[8,4]]",
	"[[[3,3],[[6,2],8]],5]",
	"[[[9,[4,8]],[[1,3],[6,7]]],[9,[[4,4],2]]]",
	"[[[[1,3],6],[[5,6],[1,9]]],[9,[[0,2],9]]]",
	"[7,[[[0,6],[1,2]],4]]",
	"[[[[5,0],[8,7]],[[7,3],0]],[[6,7],[0,1]]]",
	"[[[[5,4],7],[[8,2],1]],[[[7,0],[6,9]],0]]",
	"[[[3,[5,6]],[[9,5],4]],[[[9,4],[8,1]],[5,[7,4]]]]",
	"[[[3,[7,5]],[[8,1],8]],[[[6,3],[9,2]],[[5,7],7]]]",
	"[8,[[2,0],[[2,6],8]]]",
	"[[[[5,8],9],1],[9,6]]",
	"[[[9,9],[8,8]],[[[3,5],[8,0]],[[4,6],[3,2]]]]",
	"[[5,[[5,1],6]],[[5,8],9]]",
	"[[7,[[1,6],6]],[[[8,6],7],[6,6]]]",
	"[[0,[[9,5],0]],[4,[[7,9],[4,9]]]]",
	"[[[[4,3],[3,5]],[[1,9],[7,6]]],[3,[[6,4],[6,0]]]]",
	"[[[2,6],6],[6,3]]",
	"[[[[1,5],[3,7]],0],[3,7]]",
	"[4,[[[5,5],4],[[5,5],[9,3]]]]",
	"[[3,[8,6]],[8,[7,7]]]",
	"[8,[9,5]]",
	"[[[6,3],[2,[3,6]]],[[[6,0],[0,2]],[[8,7],5]]]",
	"[[[8,[1,2]],2],7]",
	"[[[[8,4],[2,7]],[[3,9],7]],[[4,[8,8]],[[7,4],9]]]",
	"[[[8,[2,5]],[3,[1,2]]],[[4,[5,0]],3]]",
	"[[8,[0,3]],[[5,1],[1,1]]]",
	"[[[8,[3,6]],6],[[7,[1,5]],[[4,8],9]]]",
	"[[[5,0],[0,3]],[[2,[7,8]],[1,[4,8]]]]",
	"[9,[4,[9,4]]]",
	"[[[9,[0,4]],2],3]",
	"[[9,[7,[8,9]]],3]",
	"[[[8,6],[[3,5],[9,2]]],[[3,[9,7]],5]]",
	"[[6,[[7,4],2]],[2,[7,[6,0]]]]",
	"[1,[[[2,2],6],8]]",
	"[[[6,[1,8]],[[9,3],[1,8]]],[[[8,2],[9,3]],[[8,2],[9,9]]]]",
	"[[[[2,9],[1,7]],[[4,0],8]],[[8,9],[6,3]]]",
	"[[[[2,4],[6,1]],[[5,4],[2,8]]],[8,[1,[2,4]]]]",
	"[[[4,6],[1,6]],[3,[1,1]]]",
	"[[[[8,3],8],8],[1,[[4,2],3]]]",
	"[[[9,[8,7]],[5,9]],[8,[[5,6],[4,5]]]]",
	"[[[[4,1],2],[[7,8],4]],[0,6]]",
	"[[[9,7],[[8,6],[6,9]]],[[8,[8,4]],[[9,0],2]]]",
	"[[[8,5],[1,9]],[[[2,4],5],6]]",
	"[[[9,[9,3]],[9,[2,3]]],[7,7]]",
	"[[[8,[7,4]],[2,6]],[[[4,5],[9,9]],[0,[5,2]]]]",
	"[7,[2,2]]",
	"[[[[1,8],[5,2]],3],[0,[2,[4,5]]]]",
	"[[5,[[4,8],[5,5]]],[4,[[3,4],[6,0]]]]",
	"[[3,1],[4,[3,[8,2]]]]",
	"[[3,7],[3,[[6,1],[0,2]]]]",
	"[[4,[6,2]],[[3,9],8]]",
	"[[[[2,9],3],[[5,6],4]],[8,2]]",
	"[[4,[[7,9],[4,9]]],[[4,3],[7,[0,7]]]]",
	"[[[3,[8,9]],[[3,4],[9,5]]],3]",
	"[0,[[[3,0],[8,7]],[[0,9],[9,1]]]]",
	"[[[5,[9,9]],2],[4,8]]",
	"[[[[4,4],4],5],[3,4]]",
	"[[[3,[2,2]],7],[[3,2],0]]",
	"[[[[0,5],[5,2]],2],[2,[[1,2],2]]]",
	"[[[4,6],6],[[0,1],6]]",
	"[2,[[[3,9],7],[[9,8],8]]]",
	"[[7,9],[7,[[3,0],9]]]",
	"[[[1,[6,2]],[0,8]],[[[7,2],4],9]]",
	"[[[[4,7],[1,5]],[5,9]],[[2,[0,4]],[7,[7,0]]]]",
	"[[1,[[2,0],[0,4]]],[[[4,6],9],[[6,8],[0,1]]]]",
	"[[[[6,0],7],[7,[9,6]]],[[7,[4,9]],[9,4]]]",
	"[[[5,[4,6]],[[1,9],[5,8]]],[[[3,6],[2,6]],[[7,3],7]]]",
	"[[[6,0],[6,6]],[2,8]]",
	"[[[4,[7,2]],[[5,6],[2,4]]],[[[6,8],5],[4,6]]]",
	"[[[[9,0],9],[4,0]],[[[9,1],8],[6,4]]]",
	"[[6,3],[1,[[5,0],[9,9]]]]",
	"[[[2,7],[5,6]],[[6,[1,4]],[9,9]]]",
	"[[[[0,5],3],[8,7]],[[[9,9],[6,2]],[0,7]]]",
	"[[[5,6],[1,7]],[[[0,4],9],9]]",
	"[[[7,3],3],[6,[0,[8,9]]]]",
	"[[[0,6],[[8,5],[4,6]]],[[[2,7],[4,2]],[[8,7],[0,5]]]]",
	"[[[8,[7,3]],1],8]",
	"[[8,[8,[8,2]]],[[5,4],[1,[2,6]]]]",
	"[[[[1,1],[8,6]],5],9]",
	"[[[[2,4],[5,7]],[[5,8],[3,1]]],7]",
	"[[4,[[0,1],9]],[[3,8],[4,2]]]",
	"[3,2]",
	"[[3,4],[8,[[6,5],[6,6]]]]",
	"[[[[7,0],[3,8]],[[3,3],[2,6]]],[[8,0],9]]",
}
