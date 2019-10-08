package main

import (
	"fmt"
	"biolib"
)

var p = fmt.Println

func main(){

	Dna := biolib.Readfile("Readfile")
	p(Medianstring(Dna, 6))

}

func Medianstring(Dna []string, k int) string {

	distance := 9999
	var median string	

	for _, kmer := range Makeallkmer(k) {
		if distance > Distancebtwpns(kmer, Dna) {
			distance = Distancebtwpns(kmer, Dna)
			median = kmer}
	}
	
	return median

}

func Distancebtwpns(pattern string, Dna []string) int {

	k := len(pattern)
	distance := 0
	h_distance := 0

	for _, text := range Dna {
		h_distance = 99
		for _, kmer := range biolib.List_pattern(string(text), k){
			if h_distance > biolib.Hamming_distance(pattern, kmer){
				h_distance = biolib.Hamming_distance(pattern, kmer)}}
		distance += h_distance
	}

	return distance
}

func Makeallkmer(k int) []string {

	var set []string
	
	pn := 1
	for i:=0; i<k; i++ {pn = pn * 4}
	for i:=0; i<pn; i++ {
		set = append(set, biolib.NumberToPattern(i, k))
	}	

	return set
}
