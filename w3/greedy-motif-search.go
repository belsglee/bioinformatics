// core is set benchmark >> 1st string of DNA

package main

import (
	"fmt"
	"biolib"
	"strings"
	"strconv"
)

var p = fmt.Println

func main() {

	f := biolib.Readfile("sample")
	kt := strings.Split(f[0], " ")
	k, _ := strconv.Atoi(kt[0])
	t, _ := strconv.Atoi(kt[1])
	Dna := f[1:]

	p(Greedy_motif_search(k, t, Dna))

}
func Greedy_motif_search(k, t int, Dna []string) []string  {


	var bestmotif []string

	for _, m := range Dna {
		bestmotif = append(bestmotif, m[:k])
	}

	bestscore := 9999
	curr_score := 0

	var profile [][]float64

	motifs := make([]string, t)

	for i:=0; i<=len(Dna[0])-k; i++ {
		motifs[0] = Dna[0][i:i+k]
		for j:=1; j < t; j++ {
			profile = biolib.Form_profile(motifs[0:j])
			motif := biolib.Profile_most_probable_kmer(Dna[j], k, profile)
			if motif == "" {
				motif = Dna[j][:k]
			}
			motifs[j] = motif
		}
		curr_score = biolib.Score(motifs)

		if curr_score < bestscore {
			bestmotif = nil
			bestmotif = append(bestmotif, motifs...)
			bestscore = curr_score
		}	
		
	}

	return bestmotif
	
}
