package main

import (
	"fmt"
	"math/rand"
	"biolib"
)

func main(){

	f := biolib.Readfile("test")
	Dna := f[0:]

	var bestmotif []string

	for _, motif := range Dna {  
		r := rand.Intn(len(motif)-4)
		bestmotif = append(bestmotif, motif[r:r+4])
	}
	
	profile := biolib.Profile(bestmotif)
	motifs := biolib.Motifs(profile, 4, Dna) 
	fmt.Println(bestmotif, motifs)

}
