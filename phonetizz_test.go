package phonetizz

import (
	"testing"

	"github.com/facebookgo/ensure"
)

func TestPhonetizz(t *testing.T) {
	word1 := "An und Pfirsich John"
	word2 := "Wer macht denn sowas"
	word3 := "An und für sich schon"
	score := Phonetizz(word1, word2, DefaultVWeight, DefaultCWeight)
	ensure.DeepEqual(t, score, 11.2)
	score = Phonetizz(word1, word3, DefaultVWeight, DefaultCWeight)
	ensure.DeepEqual(t, score, 2.5)
}
