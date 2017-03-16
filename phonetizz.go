package phonetizz

const (
	// DefaultVWeight can be used as the default weight for vocals
	DefaultVWeight = 0.7
	// DefaultCWeight can be used as the default weight for consonants
	DefaultCWeight = 0.3
)

var (
	vocals = []string{"a", "e", "i", "o", "u"}
)

// Phonetizz calculates the phonetizz score.
func Phonetizz(source string, target string, vWeight float64, cWeight float64) float64 {
	vSource := getString(source, true)
	vTarget := getString(target, true)
	vDistance := float64(DistanceForStrings([]rune(vSource), []rune(vTarget), DefaultOptions))
	cSource := getString(source, false)
	cTarget := getString(target, false)
	cDistance := float64(DistanceForStrings([]rune(cSource), []rune(cTarget), DefaultOptions))
	return vWeight*vDistance + cWeight*cDistance
}

func getString(input string, isVocale bool) string {
	var output string
	for idx := range input {
		contains := containsString(string(input[idx]), vocals)
		if isVocale && contains {
			output += string(input[idx])
		} else if !isVocale && !contains {
			output += string(input[idx])
		}
	}
	return output
}

func containsString(needle string, haystack []string) bool {
	for _, e := range haystack {
		if e == needle {
			return true
		}
	}
	return false
}
