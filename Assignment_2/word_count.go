// Most repeated words in a sentence
package main

import (
	"fmt"
	"strings"
)

func main() {
	var sentence string = "Eating raw fish didn't sound like a good idea. 'It's a delicacy in Japan,' didn't seem to make it any more appetizing. Raw fish is raw fish, delicacy or not."
	var i, j int
	var count []string
	var repeated []string

	fmt.Println("SENTENCE: ", sentence)
	count = strings.Split(sentence, " ")

	for i = 0; i < len(count); i++ {
		for j = i + 1; j < len(count); j++ {
			if count[i] == count[j] {
				repeated = append(repeated, count[i])
			}
		}

	}
	fmt.Println("REPEATED WORDS: ", repeated)
	fmt.Println("COUNT OF REPEATED WORDS: ", len(repeated))
}
