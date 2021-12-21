package main

import (
	"fmt"

	"github.com/tcc-sejohnson/advent-of-code-2021/04/bingo"
)

func main() {
	plays := bingo.ChallengePlays
	stringBoards := bingo.ChallengeBoards
	subsystem, err := bingo.SubsystemFromText(plays, stringBoards)
	if err != nil {
		panic("there was a problem parsing the challenge boards. Error: " + err.Error())
	}

	subsystem.PlayUntilAllBoardsWin()
	fmt.Printf("Final score: %v\n", subsystem.LastWinningBoard.Score())
}
