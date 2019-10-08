package main

import (
	"biolib"
	"fmt"
)

var p = fmt.Println

func main(){

	motifs := []string{"GGC", "CAG", "GAG", "CAC", "CAA"}
	p(Scoring_motif(motifs))

}

func Scoring_motif(motifs []string) int {

	counter := biolib.Count_motif(motifs)

	var score, max int

	for _, x := range counter {
		max = 0
		for _, count := range x {
			if max < count { 
				score = score + max
				max = count
			} else {score = score + count}
		}
	}

	return score

}
