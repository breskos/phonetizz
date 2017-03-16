# PHONETIZZ
Levenshtein based and phrase accepting phonetic algorithm written in Golang.

## Need
Most phonetic algorithms are using a code or symbols that are not easily comparable.
This motivates the development of a phonetic algorithm that uses the simplest forms of text comparison.
With Phonetizz the Levenshtein distance metric was utilized to perform such calculations for such a phonetic measure.
Another important point is that all known algorithms are just working on words.
With Phonetizz it is possible to work also on phrases instead of just one word.

## The algorithm
The idea behind Phonetizz is to use simple text metrics plus some adjustments.
Below I've listed the hypotheses that strengthen the use of Phonetizz:

* At least for spoken languages (Chinese token language counts as literary language) the vocals are more important than the rest of the tokens, because they have the most influence in the phonetic part.
* The vocal and consonant signature of the word should therefore considered separately.

Based on these hypotheses the algorithm first separates the vocals and consonants.
After that the Levenshtein distance is calculated for just vocals and just consonants.
These two distances are later combined with a different weight, where the vocal distance is higher weighted.

## Example

```go
word1 := "An und Pfirsich John"
word2 := "Wer macht denn sowas"
word3 := "An und für sich schon"

score := phonetizz.Phonetizz(word1, word2, phonetizz.DefaultVWeight, phonetizz.DefaultCWeight)
// score equals 11.2
score = phonetizz.Phonetizz(word1, word3, phonetizz.DefaultVWeight, phonetizz.DefaultCWeight)
// score equals 2.5
```

## Experimentation
At the moment, no representative evaluation of the algorithm exist.
Let me know if you are trying or want to do this in collaboration with me.



## Use Case: Rhymzz
Ryhmzz is a modified version of Phonetizz that applies another layer of different weights on top of Phonetizz.
Are you looking for words that are completing your rhyme then Rhymzz helps you with that.
With the use of Phonetizz (that can also used for that case) comes the possibility to get phonetic similarities of (so called) punch lines for a battle rap.
Maybe the future battle rap competitions are ruled by the guys with the fastest computer.

In contrast to Phonetizz the experiments (also not representative) here are performed just on words.
```
"mähen" -> winner word score: gehen
"probieren" -> winner word: verlieren
"Studenten" -> winner word: Enten
```
