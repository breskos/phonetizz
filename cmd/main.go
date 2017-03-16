package main

import (
	"fmt"

	"bitbucket.com/flezzfx/phonetizz"
)

func main() {
	word1 := "An und Pfirsich John"
	word2 := "Wer macht denn sowas"
	word3 := "An und für sich schon"

	score := phonetizz.Phonetizz(word1, word2, phonetizz.DefaultVWeight, phonetizz.DefaultCWeight)
	fmt.Printf("'%v'  VS  ''%v' = %v \n", word1, word2, score)
	score = phonetizz.Phonetizz(word1, word3, phonetizz.DefaultVWeight, phonetizz.DefaultCWeight)
	fmt.Printf("'%v'  VS  ''%v' = %v \n", word1, word3, score)
	return
}
