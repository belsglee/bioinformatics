// This is code challenge as a part of coursera 'bioinformatics specialized course'
// created by Phillip, Pavel (UC San Diego)
// Week 1, Chapter 1, Part 1, Code Challenge 3

package main
import (
	"fmt"
)

func main() {

	text := "ACGTTGCATGTCGCATGATGCATGAGAGCT"

	tlen := len(text)
	reversed := make([]int32, tlen, tlen)

	for i, nucle := range text {
		if nucle == 65 {reversed[tlen-i-1] = 84}	// if nucleotide is "A" -> "T"
		if nucle == 67 {reversed[tlen-i-1] = 71}	// if nucleotide is "C" -> "G"
		if nucle == 71 {reversed[tlen-i-1] = 67}	// if nucleotide is "G" -> "C"
		if nucle == 84 {reversed[tlen-i-1] = 65}	// if nucleotide is "T" -> "A"
	}

	fmt.Println(string(reversed))
	
}

