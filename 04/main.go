package main

import (
	"fmt"

	"example.com/aoc04/bingo"
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
