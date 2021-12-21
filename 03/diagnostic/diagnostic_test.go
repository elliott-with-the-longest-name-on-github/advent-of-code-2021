package diagnostic_test

import (
	"testing"

	"github.com/tcc-sejohnson/advent-of-code-2021/03/diagnostic"
)

func TestDiagnosticByteToInt(t *testing.T) {
	tests := []struct {
		Byte diagnostic.DiagnosticByte
		Want int
	}{
		{Byte: []bool{false, false, false, false, false, false, false, false, false, false}, Want: 0},
		{Byte: []bool{false, false, false, false, false, false, false, false, false, true}, Want: 1},
		{Byte: []bool{false, false, false, false, false, false, false, false, true, true}, Want: 3},
		{Byte: []bool{false, false, false, false, false, false, false, false, true, false}, Want: 2},
		{Byte: []bool{false, false, true, false, false, false, false, false, false, false}, Want: 128},
		{Byte: []bool{false, true, false, false, false, false, false, false, false, false}, Want: 256},
		{Byte: []bool{true, false, false, false, false, false, false, false, false, false}, Want: 512},
		{Byte: []bool{true, true, true, false, false, false, false, false, false, true}, Want: 897},
		{Byte: []bool{false, false, false, false, false, false, true, false, false, false}, Want: 8},
	}

	for _, v := range tests {
		got := v.Byte.ToInteger()
		if v.Want != got {
			t.Errorf("value from DiagnosticByte.ToInteger incorrect. want: %v, got: %v", v.Want, got)
		}
	}
}

func TestDiagnosticLogPowerConsumption(t *testing.T) {
	inputBinary := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	log, err := diagnostic.FromBinaryStrings(inputBinary, 5)
	if err != nil {
		t.Errorf("diagnostic package failed to parse binary with error %s", err)
	}

	pc := log.PowerConsumption()
	if pc != 198 {
		t.Errorf("incorrect value for Power Consumption. Want: %v, got: %v", 198, pc)
	}
}

func TestLifeSupportRating(t *testing.T) {
	inputBinary := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	log, err := diagnostic.FromBinaryStrings(inputBinary, 5)
	if err != nil {
		t.Errorf("diagnostic package failed to parse binary with error %s", err)
	}

	lsr := log.LifeSupportRating()
	if lsr != 230 {
		t.Errorf("incorrect value for Life Support Rating. Want: %v, got: %v", 230, lsr)
	}
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		diagnostic.PartOne()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		diagnostic.PartTwo()
	}
}
