package biolib

import (
	"fmt"
	"strings"
	"strconv"
)

func Profile_most_probable_kmer() string {
	
	f := Readfile("Readfile")
	Dna := f[0]
	k, _ := strconv.Atoi(f[1])
	profile := make([][]float64, len(f[2:]))

	for i, l := range f[2:]{
		profile[i] =  make([]float64, k)
		for j, float := range strings.Split(l, " "){
			profile[i][j], _ = strconv.ParseFloat(float, 64)
		}
	}
	
	fmt.Println(profile)	

	var best_pt string
	var best_pb float64
	best_pb = 0.0

	for i:=0; i<=len(Dna)-k; i++{
		kmer := Dna[i:i+k]
		if Calculate_probability(kmer, profile) > best_pb{
			best_pt = kmer
			best_pb = Calculate_probability(kmer, profile)}
	}
	return best_pt
}

func Calculate_probability(kmer string, profile [][]float64) float64 {
	t := []float64{}
	for i, nucle := range kmer{
		if nucle == 65 {t = append(t, profile[0][i])}
		if nucle == 67 {t = append(t, profile[1][i])}
		if nucle == 71 {t = append(t, profile[2][i])}
		if nucle == 84 {t = append(t, profile[3][i])}
	}

	return Multiple(t)
}

func Multiple(set []float64) float64 {

	var m float64
	m = 1.0
	
	for _, item := range set {
		m = m * item
	}

	return m
}
