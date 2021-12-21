package polymerization_test

import (
	"strconv"
	"testing"

	"github.com/tcc-sejohnson/advent-of-code-2021/14/polymerization"
)

var template string = "NNCB"
var rules map[string]rune = map[string]rune{
	"CH": 'B',
	"HH": 'N',
	"CB": 'H',
	"NH": 'C',
	"HB": 'C',
	"HC": 'B',
	"HN": 'C',
	"NN": 'C',
	"BH": 'H',
	"NC": 'B',
	"NB": 'B',
	"BN": 'B',
	"BB": 'N',
	"BC": 'B',
	"CC": 'N',
	"CN": 'C',
}

func makeTestTemplate() *polymerization.PolymerTemplate {
	return polymerization.New(template, rules)
}

func mapFromBasePolymer(base string) map[string]int {
	polymerMap := make(map[string]int)
	runes := []rune(base)
	for i, r := range runes {
		if i == 0 {
			continue
		}
		pair := string(runes[i-1]) + string(r)
		polymerMap[pair] += 1
	}
	return polymerMap
}

func TestSubstitutions(t *testing.T) {
	template := makeTestTemplate()
	testPolymers := []map[string]int{
		mapFromBasePolymer("NNCB"),
		mapFromBasePolymer("NCNBCHB"),
		mapFromBasePolymer("NBCCNBBBCBHCB"),
		mapFromBasePolymer("NBBBCNCCNBBNBNBBCHBHHBCHB"),
		mapFromBasePolymer("NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB"),
	}

	for i, want := range testPolymers {
		for pair, count := range template.Polymer {
			if _, exists := want[pair]; !exists {
				t.Errorf("substitution procuded an incorrect result on iteration %v. Result is missing pair %v", i, pair)
			} else if exists && want[pair] != count {
				t.Errorf("substitution procuded an incorrect result on iteration %v. Result pair count is off for pair %v. Want: %v, got: %v", i, pair, want[pair], count)
			}
		}
		template.Substitute()
	}
}

func TestManySubstitutions(t *testing.T) {
	template := makeTestTemplate()
	want := mapFromBasePolymer("NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB")
	template.SubstituteMany(4)
	for pair, count := range template.Polymer {
		if _, exists := want[pair]; !exists {
			t.Errorf("substitution procuded an incorrect result after 4 iterations. Result is missing pair %v", pair)
		} else if exists && want[pair] != count {
			t.Errorf("substitution procuded an incorrect result after 4 iterations. Result pair count is off for pair %v. Want: %v, got: %v", pair, want[pair], count)
		}
	}
}

func TestLeastFrequentElement(t *testing.T) {
	template := makeTestTemplate()
	template.SubstituteMany(10)
	wantRune := 'H'
	wantCount := 161
	gotRune, gotCount := template.LeastFrequentElement()
	if wantRune != gotRune {
		t.Errorf("substitution produced an incorrect LeastFrequentElement after 10 iterations. Want: %v, got: %v", wantRune, gotRune)
	}
	if wantCount != gotCount {
		t.Errorf("substitution produced an incorrect LeastFrequentElement count after 10 iterations. Want: %v, got: %v", wantCount, gotCount)
	}

	template.SubstituteMany(30)
	wantRune = 'H'
	wantCount = 3849876073
	gotRune, gotCount = template.LeastFrequentElement()
	if wantRune != gotRune {
		t.Errorf("substitution produced an incorrect LeastFrequentElement after 10 iterations. Want: %v, got: %v", wantRune, gotRune)
	}
	if wantCount != gotCount {
		t.Errorf("substitution produced an incorrect LeastFrequentElement count after 10 iterations. Want: %v, got: %v", wantCount, gotCount)
	}
}

func TestMostFrequentElement(t *testing.T) {
	template := makeTestTemplate()
	wantRune := 'B'
	wantCount := 1749
	template.SubstituteMany(10)
	gotRune, gotCount := template.MostFrequentElement()
	if wantRune != gotRune {
		t.Errorf("substitution produced an incorrect MostFrequentElement after 10 iterations. Want: %v, got: %v", wantRune, gotRune)
	}
	if wantCount != gotCount {
		t.Errorf("substitution produced an incorrect MostFrequentElement count after 10 iterations. Want: %v, got: %v", wantCount, gotCount)
	}

	template.SubstituteMany(30)
	wantRune = 'B'
	wantCount = 2192039569602
	gotRune, gotCount = template.MostFrequentElement()
	if wantRune != gotRune {
		t.Errorf("substitution produced an incorrect MostFrequentElement after 10 iterations. Want: %v, got: %v", wantRune, gotRune)
	}
	if wantCount != gotCount {
		t.Errorf("substitution produced an incorrect MostFrequentElement count after 10 iterations. Want: %v, got: %v", wantCount, gotCount)
	}
}

func BenchmarkSubstitutions(b *testing.B) {
	for i := 10; i <= 200; i += 10 {
		template := makeTestTemplate()
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				template.SubstituteMany(i)
			}
		})
	}
}
