package polymerization

type PolymerTemplate struct {
	BasePolymer         string
	Polymer             map[string]int
	PolymerElementCount map[rune]int
	SubstitutionRules   map[string]rune
}

func New(base string, substitutionRules map[string]rune) *PolymerTemplate {
	template := &PolymerTemplate{
		BasePolymer:         base,
		Polymer:             make(map[string]int),
		PolymerElementCount: make(map[rune]int),
		SubstitutionRules:   substitutionRules,
	}
	runes := []rune(base)
	for i, r := range runes {
		template.incrementElementCount(r)
		if i == 0 {
			continue
		}
		template.incrementPairCount(string(runes[i-1]) + string(r))
	}
	return template
}

func (p *PolymerTemplate) SubstituteMany(iterations int) {
	for i := 0; i < iterations; i++ {
		p.Substitute()
	}
}

func (p *PolymerTemplate) Substitute() {
	newPolymer := make(map[string]int)
	for pair, count := range p.Polymer {
		runes := []rune(pair)
		newRune := p.SubstitutionRules[pair]
		p.incrementElementCountBy(newRune, count)
		newPolymer[string(runes[0])+string(newRune)] += count
		newPolymer[string(newRune)+string(runes[1])] += count
	}
	p.Polymer = newPolymer
}

func (p *PolymerTemplate) MostFrequentElement() (element rune, count int) {
	maxFreq := 0
	var maxFreqElement rune
	for k, v := range p.PolymerElementCount {
		if v > maxFreq {
			maxFreq = v
			maxFreqElement = k
		}
	}
	return maxFreqElement, maxFreq
}

func (p *PolymerTemplate) LeastFrequentElement() (element rune, count int) {
	minFreq := 0
	var minFreqElement rune
	for k, v := range p.PolymerElementCount {
		if v < minFreq || minFreq == 0 {
			minFreq = v
			minFreqElement = k
		}
	}
	return minFreqElement, minFreq
}

func (p *PolymerTemplate) incrementPairCount(pair string) (newCount int) {
	return p.incrementPairCountBy(pair, 1)
}

func (p *PolymerTemplate) incrementPairCountBy(pair string, times int) (newCount int) {
	p.Polymer[pair] += times
	newCount = p.Polymer[pair]
	if newCount == 0 {
		delete(p.Polymer, pair)
	}
	return newCount
}

func (p *PolymerTemplate) incrementElementCount(element rune) (newCount int) {
	return p.incrementElementCountBy(element, 1)
}

func (p *PolymerTemplate) incrementElementCountBy(element rune, times int) (newCount int) {
	p.PolymerElementCount[element] += times
	newCount = p.PolymerElementCount[element]
	if newCount == 0 {
		delete(p.PolymerElementCount, element)
	}
	return newCount
}

var ChallengeBase string = "CKFFSCFSCBCKBPBCSPKP"
var ChallengeRules map[string]rune = map[string]rune{
	"NS": 'P',
	"KV": 'B',
	"FV": 'S',
	"BB": 'V',
	"CF": 'O',
	"CK": 'N',
	"BC": 'B',
	"PV": 'N',
	"KO": 'C',
	"CO": 'O',
	"HP": 'P',
	"HO": 'P',
	"OV": 'O',
	"VO": 'C',
	"SP": 'P',
	"BV": 'H',
	"CB": 'F',
	"SF": 'H',
	"ON": 'O',
	"KK": 'V',
	"HC": 'N',
	"FH": 'P',
	"OO": 'P',
	"VC": 'F',
	"VP": 'N',
	"FO": 'F',
	"CP": 'C',
	"SV": 'S',
	"PF": 'O',
	"OF": 'H',
	"BN": 'V',
	"SC": 'V',
	"SB": 'O',
	"NC": 'P',
	"CN": 'K',
	"BP": 'O',
	"PC": 'H',
	"PS": 'C',
	"NB": 'K',
	"VB": 'P',
	"HS": 'V',
	"BO": 'K',
	"NV": 'B',
	"PK": 'K',
	"SN": 'H',
	"OB": 'C',
	"BK": 'S',
	"KH": 'P',
	"BS": 'S',
	"HV": 'O',
	"FN": 'F',
	"FS": 'N',
	"FP": 'F',
	"PO": 'B',
	"NP": 'O',
	"FF": 'H',
	"PN": 'K',
	"HF": 'H',
	"VK": 'K',
	"NF": 'K',
	"PP": 'H',
	"PH": 'B',
	"SK": 'P',
	"HN": 'B',
	"VS": 'V',
	"VN": 'N',
	"KB": 'O',
	"KC": 'O',
	"KP": 'C',
	"OS": 'O',
	"SO": 'O',
	"VH": 'C',
	"OK": 'B',
	"HH": 'B',
	"OC": 'P',
	"CV": 'N',
	"SH": 'O',
	"HK": 'N',
	"NO": 'F',
	"VF": 'S',
	"NN": 'O',
	"FK": 'V',
	"HB": 'O',
	"SS": 'O',
	"FB": 'B',
	"KS": 'O',
	"CC": 'S',
	"KF": 'V',
	"VV": 'S',
	"OP": 'H',
	"KN": 'F',
	"CS": 'H',
	"CH": 'P',
	"BF": 'F',
	"NH": 'O',
	"NK": 'C',
	"OH": 'C',
	"BH": 'O',
	"FC": 'V',
	"PB": 'B',
}
