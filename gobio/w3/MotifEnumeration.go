package main

import (
	"fmt"
	"biolib"
//	"os"
//	"bufio"
)

func main() {

	dna := biolib.Readfile("filename")
	fmt.Println(Motifenumeration(dna, 5, 2))

}

func Motifenumeration(dna []string, k, d int) []string{

	var candidates = []string{}

	for i, pattern := range dna {
		for _, candidate :=	range biolib.Find_Neighbor(pattern, k, d) {
			if Appear(candidate, dna, k, d, i) {
				candidates = append(candidates, candidate)}}}

	return biolib.Unique(candidates)

}

func Appear (kmer string, dna []string, k, d, index int) bool {

	count := 0

	for idx, text := range dna {
		for i:=0; i <= len(text)-k; i++ {
			if idx != index && biolib.Hamming_distance(kmer, text[i:i+k]) <= d {
					count++
					break}}
//		fmt.Println(index, kmer, text, count)
		if count == len(dna) - 1 {return true}
	}

	return false


}
