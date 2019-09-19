package biolib

import (
	"fmt"
)

func Scoring_motif(motifs []string) int {

	counter := Count_motif(motifs)

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
