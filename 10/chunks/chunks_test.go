package chunks_test

import (
	"errors"
	"testing"

	"github.com/tcc-sejohnson/advent-of-code-2021/10/chunks"
)

func TestIncompleteScoring(t *testing.T) {
	input := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}
	want := 288957

	logDump, err := chunks.ParseLogDumpFromString(input)
	if err != nil {
		t.Errorf("failed to parse LogDump with error %v", err)
	}

	got := logDump.IncompleteScore()
	if want != got {
		t.Errorf("incorrect score reported by LogDump. Want: %v, got: %v", want, got)
	}
}

func TestCorruptedScoring(t *testing.T) {
	input := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}
	want := 26397

	logDump, err := chunks.ParseLogDumpFromString(input)
	if err != nil {
		t.Errorf("failed to parse LogDump with error %v", err)
	}

	got := logDump.CorruptedScore()
	if want != got {
		t.Errorf("incorrect score reported by LogDump. Want: %v, got: %v", want, got)
	}
}

func TestSingleLineParse(t *testing.T) {
	// assertSuccessfulParse := func(t *testing.T, input string, chunks []*chunks.Chunk, err error) {
	// 	if err != nil {
	// 		t.Errorf("expected line %v to successfully parse, but got error %v", input, err)
	// 	}
	// }
	assertLineCorruptedError := func(t *testing.T, input string, chks []*chunks.Chunk, err error) {
		if err == nil {
			t.Errorf("expected line %v to fail parsing with a line corrupted error, but parsing succeeded", input)
		}
		var lineCorruptedError *chunks.LineCorruptedError
		if !errors.As(err, &lineCorruptedError) {
			t.Errorf("expected line %v to fail parsing with a line corrupted error, but it failed with a different error", input)
		}
	}
	assertLineIncompleteError := func(t *testing.T, input string, chks []*chunks.Chunk, err error) {
		if err == nil {
			t.Errorf("expected line %v to fail parsing with a line incomplete error, but parsing succeeded", input)
		}
		var lineIncompleteError *chunks.LineIncompleteError
		if !errors.As(err, &lineIncompleteError) {
			t.Errorf("expected line %v to fail parsing with a line incomplete error, but it failed with a different error", input)
		}
	}
	input := []struct {
		Line       string
		AssertFunc func(t *testing.T, input string, chks []*chunks.Chunk, err error)
	}{
		{
			Line:       "[({(<(())[]>[[{[]{<()<>>",
			AssertFunc: assertLineIncompleteError,
		},
		{
			Line:       "[(()[<>])]({[<{<<[]>>(",
			AssertFunc: assertLineIncompleteError,
		},
		{
			Line:       "{([(<{}[<>[]}>{[]{[(<()>",
			AssertFunc: assertLineCorruptedError,
		},
		{
			Line:       "(((({<>}<{<{<>}{[]{[]{}",
			AssertFunc: assertLineIncompleteError,
		},
		{
			Line:       "[[<[([]))<([[{}[[()]]]",
			AssertFunc: assertLineCorruptedError,
		},
		{
			Line:       "[{[{({}]{}}([{[{{{}}([]",
			AssertFunc: assertLineCorruptedError,
		},
		{
			Line:       "{<[[]]>}<{[{[{[]{()[[[]",
			AssertFunc: assertLineIncompleteError,
		},
		{
			Line:       "[<(<(<(<{}))><([]([]()",
			AssertFunc: assertLineCorruptedError,
		},
		{
			Line:       "<{([([[(<>()){}]>(<<{{",
			AssertFunc: assertLineCorruptedError,
		},
		{
			Line:       "<{([{{}}[<[[[<>{}]]]>[]]",
			AssertFunc: assertLineIncompleteError,
		},
	}

	for _, test := range input {
		chks, err := chunks.ParseChunksFromString(test.Line)
		test.AssertFunc(t, test.Line, chks, err)
	}
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chunks.PartOne(chunks.ChallengeInput)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chunks.PartTwo(chunks.ChallengeInput)
	}
}
