package dirac

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MIN_BOARD_SPACE       int = 1
	MAX_BOARD_SPACE       int = 10
	NUM_ROLLS             int = 3
	NUM_QUANTUM_DIE_SIDES int = 3
)

type DeterministicDie struct {
	CurrentRoll int
	TotalRolls  int
}

type Die interface {
	Roll(times int) int
	NumRolls() int
}

type Player struct {
	ID               int
	StartingPosition int
	CurrentPosition  int
	Score            int
}

type DiracDice struct {
	Die           Die
	Players       []*Player
	Turn          int
	WinThreshhold int
}

func PartOne() string {
	dd, err := DiracDiceFromStrings(ChallengeInput, &DeterministicDie{}, 1000)
	if err != nil {
		panic(err)
	}
	winner, loser := dd.PlayUntilVictory()
	numRolls := dd.Die.NumRolls()
	return fmt.Sprintf("Part One: Player %v won after %v turns with %v points. Player %v lost with %v points.\nChallenge solution: %v * %v = %v", winner.ID, numRolls, winner.Score, loser.ID, loser.Score, loser.Score, numRolls, loser.Score*numRolls)
}

func PartTwo() string {
	dd, err := DiracDiceFromStrings(ChallengeInput, &DeterministicDie{}, 21)
	if err != nil {
		panic(err)
	}
	p1Wins, p2Wins := dd.PlayForQuantumSupremacy()
	return fmt.Sprintf("Part Two: Player 1 won %v times, Player 2 won %v times.", p1Wins, p2Wins)
}

func PlayersFromStrings(strs []string) ([]*Player, error) {
	players := make([]*Player, len(strs))
	for i, str := range strs {
		fields := strings.Fields(str)
		if len(fields) != 5 {
			return nil, fmt.Errorf("failed to parse input string. Should be in the format 'Player x starting position: x' but got %v", str)
		}
		playerId, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse input string. Second field in the input string should be an int, but got %v from string %v", fields[1], str)
		}
		playerStartingPosition, err := strconv.Atoi(fields[4])
		if err != nil {
			return nil, fmt.Errorf("failed to parse input string. Fifth field in the input string should be an int, but got %v from string %v", fields[4], str)
		}
		players[i] = &Player{ID: playerId, StartingPosition: playerStartingPosition - 1, CurrentPosition: playerStartingPosition - 1}
	}
	return players, nil
}

func DiracDiceFromStrings(strs []string, die Die, winThreshhold int) (*DiracDice, error) {
	players, err := PlayersFromStrings(strs)
	if err != nil {
		return nil, err
	}
	dd := &DiracDice{
		Die:           die,
		Players:       players,
		Turn:          0,
		WinThreshhold: winThreshhold,
	}
	return dd, nil
}

func (d *DiracDice) PlayUntilVictory() (winner *Player, loser *Player) {
	var won bool
	for !won {
		loser = winner
		won, winner = d.TakeTurn()
	}
	return
}

type quantumFrequencyMap []struct {
	rollSum     int
	occurrences int
}

type quantumRecursionMemo map[string]struct{ p1S, p2S int }

// PlayForQuantumSupremacy. I mean, yeah, I could create a new QuantumDie and figure
// out how to fit this into the neat framework I built for Part 1, but frankly, I have a life.
func (d *DiracDice) PlayForQuantumSupremacy() (playerOneWins, playerTwoWins int) {
	// There IS a way to do this recursively for arbitrary die sides/rolls, but... no.
	frequencies := quantumFrequencyMap{{3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1}}
	memo := make(quantumRecursionMemo)
	return playForQuantumRecursion(frequencies, memo, d.Players[0].CurrentPosition, d.Players[1].CurrentPosition, d.Players[0].Score, d.Players[1].Score)
}

func playForQuantumRecursion(m quantumFrequencyMap, memo quantumRecursionMemo, playerOnePos, playerTwoPos, playerOneScore, playerTwoScore int) (playerOneWins, playerTwoWins int) {
	if playerOneScore >= 21 {
		return 1, 0
	}
	if playerTwoScore >= 21 {
		return 0, 1
	}

	totalPlayerOneWins := 0
	totalPlayerTwoWins := 0

	for _, o := range m {
		newPos := (playerOnePos + o.rollSum) % 10
		newScore := playerOneScore + newPos + 1
		playerTwoWins, playerOneWins = playForQuantumRecursion(m, memo, playerTwoPos, newPos, playerTwoScore, newScore)

		totalPlayerOneWins += playerOneWins * o.occurrences
		totalPlayerTwoWins += playerTwoWins * o.occurrences
	}
	return totalPlayerOneWins, totalPlayerTwoWins
}

func (d *DiracDice) TakeTurn() (won bool, player *Player) {
	if d.Turn >= len(d.Players) {
		d.Turn = 0
	}
	defer func() { d.Turn += 1 }()
	return d.Players[d.Turn].TakeTurn(d.Die, MIN_BOARD_SPACE, MAX_BOARD_SPACE, d.WinThreshhold), d.Players[d.Turn]
}

func (p *Player) TakeTurn(die Die, minBoardSpace, maxBoardSpace, winThreshhold int) (won bool) {
	numBoardSpaces := (maxBoardSpace - minBoardSpace + 1)
	moveValue := die.Roll(NUM_ROLLS)
	p.CurrentPosition = (moveValue + p.CurrentPosition) % numBoardSpaces // don't ya love how 0-based indexing works??
	p.Score += p.CurrentPosition + 1
	return p.Score >= winThreshhold
}

func (d *DeterministicDie) Roll(times int) int {
	d.TotalRolls += times
	// hehe, math tricks
	// if it's not the first role, the current roll
	// is always times^2.
	d.CurrentRoll += times * times
	if d.TotalRolls == times {
		// The formula for 1 to n is n(n+1)/2
		d.CurrentRoll = (times * (times + 1)) / 2
	}
	return d.CurrentRoll
}

func (d *DeterministicDie) NumRolls() int {
	return d.TotalRolls
}

var ChallengeInput []string = []string{
	"Player 1 starting position: 9",
	"Player 2 starting position: 6",
}
