package dirac_test

import (
	"testing"

	"github.com/tcc-sejohnson/advent-of-code-2021/21/dirac"
)

func TestQuantumSupremacy(t *testing.T) {
	dd, err := dirac.DiracDiceFromStrings([]string{"Player 1 starting position: 4", "Player 2 starting position: 8"}, &dirac.DeterministicDie{}, 21)
	if err != nil {
		t.Errorf("%v", err)
	}
	p1Wins, p2Wins := dd.PlayForQuantumSupremacy()
	if p1Wins != 444356092776315 || p2Wins != 341960390180808 {
		t.Errorf("scores are off. P1 want: 444356092776315, got: %v, P2 want: 341960390180808, got: %v", p1Wins, p2Wins)
	}
}

func TestGame(t *testing.T) {
	testPlayers := []string{"Player 1 starting position: 4", "Player 2 starting position: 8"}
	dd, err := dirac.DiracDiceFromStrings(testPlayers, &dirac.DeterministicDie{}, 1000)
	if err != nil {
		t.Errorf("failed to parse input strings with error %v", err)
	}
	winner, loser := dd.PlayUntilVictory()
	if winner.ID != 1 {
		t.Errorf("Player 2 won, but Player 1 should have")
	}
	if winner.Score != 1000 {
		t.Errorf("Player 1 score was incorrect. Wanted: 1000, got: %v", winner.Score)
	}
	if loser.Score != 745 {
		t.Errorf("Player 2 score was incorrect. Wanted: 745, got: %v", loser.Score)
	}
	if dd.Die.NumRolls() != 993 {
		t.Errorf("Number of die rolls was incorrect. Wanted: 993, got: %v", dd.Die.NumRolls())
	}
}

func TestFromStringsSuccess(t *testing.T) {
	tests := []struct {
		Input       []string
		WantPlayers int
	}{
		{Input: []string{"Player 1 starting position: 4", "Player 2 starting position: 8"}, WantPlayers: 2},
	}

	for _, test := range tests {
		dd, err := dirac.DiracDiceFromStrings(test.Input, &dirac.DeterministicDie{}, 1000)
		if err != nil {
			t.Errorf("failed to parse input strings with error %v", err)
		}
		if len(dd.Players) != test.WantPlayers {
			t.Errorf("number of players parsed did not match expected. Want: %v, got: %v", test.WantPlayers, len(dd.Players))
		}
	}
}

func TestFromStringsErrors(t *testing.T) {
	tests := [][]string{
		{"A string with > 5 fields"},
		{"Less than 5 fields"},
		{"Second field not a number"},
		{"Fifth 5 field not number"},
	}

	for _, test := range tests {
		_, err := dirac.DiracDiceFromStrings(test, &dirac.DeterministicDie{}, 1000)
		if err == nil {
			t.Errorf("FromStrings should have failed on input %v, but did not", err)
		}
	}
}

func TestDeterministicDieRoll(t *testing.T) {
	die := dirac.DeterministicDie{}
	for i := 1; i < 100000; i++ {
		gotVal := die.Roll(1)
		if gotVal != i {
			t.Errorf("deterministic die's rolls weren't deterministic. Want: %v, got: %v", i, gotVal)
		}
		gotRolls := die.TotalRolls
		if gotRolls != i {
			t.Errorf("deterministic die's roll count is incorrect. Want: %v, got: %v", i, gotRolls)
		}
	}
}

func TestDeterministicdieRollThrice(t *testing.T) {
	die := dirac.DeterministicDie{}
	want := 6
	for i := 1; i < 100000; i++ {
		gotVal := die.Roll(3)
		if gotVal != want {
			t.Errorf("deterministic die wasn't deterministic. Want: %v, got: %v", want, gotVal)
		}
		want += 9
		gotRolls := die.TotalRolls
		wantRolls := i * 3
		if gotRolls != wantRolls {
			t.Errorf("deterministic die's roll count is incorrect. Want: %v, got: %v", wantRolls, gotRolls)
		}
	}
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dirac.PartOne()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dirac.PartTwo()
	}
}
