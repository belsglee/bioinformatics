package biolib

import "fmt"

var p = fmt.Println

func Form_profile(Dna []string) [][]float64 {
	
	l := len(Dna[0])

	counter := make([][]int, l)
	for i := range counter{
		counter[i] = make([]int, 4)
	}
	profile := make([][]float64, l)
	for i := range profile{
		profile[i] = make([]float64, 4)
	}

	// count(Motifs)
	for _, motif := range Dna {
		for j, nucle := range motif {
			if nucle == 65 {counter[j][0]++}
			if nucle == 67 {counter[j][1]++}
			if nucle == 71 {counter[j][2]++}
			if nucle == 87 {counter[j][3]++}
		}
	}

	for i, x := range counter {
		for j, value := range x {
			if value != 0 {
				profile[i][j] = float64(value)/float64(len(Dna))
			}
		}
	}

	return profile

}
