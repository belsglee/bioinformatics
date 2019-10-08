package main

import (
	"biolib"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

var p = fmt.Println

func main() {

	// will be much more easier if type sturct is dfined...

	f := biolib.Readfile("Sample")
	kt := strings.Split(f[0], " ")
	k, _ := strconv.Atoi(kt[0])
	t, _ := strconv.Atoi(kt[1])
	Dna := f[1:]

	//	p(k, t, Dna)

	p(Randomized_motif_search(Dna, k, t))
}

func Repeat_Randomized_motif_search(Dna []string, k, t int) []string {

	var bestscore int
	var bestmotif []string
	var motifs []string
	var score int

	for i := 0; i < 1000; i++ {
		motifs = Randomized_motif_search(Dna, k, t)
		score = biolib.Score(motifs)

		if score < bestscore {
			bestscore = score
			bestmotif = nil
			bestmotif = append(bestmotif, motifs...)
		}
	}

	return bestmotif
}

func Randomized_motif_search(Dna []string, k, t int) []string {

	// init motif
	var motifs []string

	// randomly select
	for _, motif := range Dna {
		n := rand.Intn(len(motif) - k)
		motifs = append(motifs, motif[n:n+k])
	}

	// init bestmotif
	bestmotif := motifs

	// training loop

	var profile [][]float64

	for {
		profile = biolib.Form_profile(motifs)
		motifs = biolib.Profile_most_probable_kmer(k, t, Dna)

		if biolib.Score(motifs) < biolib.Score(bestmotif) {
			bestmotif = nil
			bestmotif = append(bestmotif, motifs...)
		} else {
			return bestmotif
		}
	}
}
