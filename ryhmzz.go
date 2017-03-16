package phonetizz

import "math"

const (
	// DefaultVDecay describes the decay from right to the left
	DefaultVDecayStart = 1.0
	DefaultVDecay      = 0.3
	DefaultCDecayStart = 1.0
	DefaultCDecay      = 0.5
)

// Ryhmzz calculates the rhymzz score.
func Ryhmzz(source string, target string, vStart, cStart, vDecay, cDecay float64) float64 {
	vSource := getString(source, true)
	vTarget := getString(target, true)
	cSource := getString(source, false)
	cTarget := getString(target, false)
	similarityVocals := calculateSimilarity(vSource, vTarget, vStart, vDecay)
	similarityConsontants := calculateSimilarity(cSource, cTarget, cStart, cDecay)
	distance := math.Abs(float64(len(source) - len(target)))

	return (similarityVocals + similarityConsontants) / (distance + 1)
}

func calculateSimilarity(source, target string, start, decay float64) float64 {
	weight := start
	score := 0.0
	if len(source) > len(target) {
		source = source[len(source)-len(target):]
	} else {
		target = target[len(target)-len(source):]
	}
	l := len(source)
	for i := l; i > 0; i-- {
		if source[i-1] == target[i-1] {
			score += weight
		}
		if weight > 0.0 {
			weight -= decay
		}
	}

	return score
}
