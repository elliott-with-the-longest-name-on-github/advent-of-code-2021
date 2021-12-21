package bingo_test

import (
	"testing"

	"github.com/tcc-sejohnson/advent-of-code-2021/04/bingo"
)

func TestBingoSubsystemWinDetection(t *testing.T) {
	stringBoards := [][]string{
		{
			"22 13 17 11  0",
			"8  2 23  4 24",
			"21  9 14 16  7",
			"6 10  3 18  5",
			"1 12 20 15 19",
		},
		{
			"3 15  0  2 22",
			"9 18 13 17  5",
			"19  8  7 25 23",
			"20 11 10 24  4",
			"14 21 16 12  6",
		},
		{
			"14 21 17 24  4",
			"10 16 15  9 19",
			"18  8 23 26 20",
			"22 11 13  6  5",
			"2  0 12  3  7",
		},
	}
	plays := []int{
		7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1,
	}

	subsystem, err := bingo.SubsystemFromText(plays, stringBoards)
	if err != nil {
		t.Errorf("failed to parse boards with error %s", err)
	}

	subsystem.PlayUntilAllBoardsWin()

	winningScore := subsystem.WinningBoard.Score()
	if winningScore != 4512 {
		t.Errorf("winning board had an incorrect score. Wanted: %v, got: %v", 4512, winningScore)
	}

	lastFinishedScore := subsystem.LastWinningBoard.Score()
	if lastFinishedScore != 1924 {
		t.Errorf("last finished board had an incorrect score. Wanted: %v, got: %v", 1924, lastFinishedScore)
	}
}

func TestBingoBoardScoring(t *testing.T) {
	stringBoard := []string{
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		"2  0 12  3  7",
	}

	b, err := bingo.BoardFromText(stringBoard)
	if err != nil {
		t.Errorf("failed to parse bingo board with error %s", err)
	}

	numbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24}
	for i, number := range numbers {
		won := b.PlayNumber(number)
		if i != len(numbers)-1 && won {
			t.Errorf("board won before it should've on turn %v with number %v", i, number)
		}
	}
	if !b.Won {
		t.Errorf("board did not win after all numbers were played")
	}

	score := b.Score()
	if score != 4512 {
		t.Errorf("board score incorrect. Wanted %v, got %v", 4512, score)
	}
}
