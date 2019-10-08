// This is code challenge as a part of coursera 'bioinformatics specialized course'
// created by Phillip, Pavel (UC San Diego)
// Week 1, Chapter 1, Part 1, Code Challenge 3

package bioinformatics
import (
	"bioinformaics/lib"
)

func reversed(text string) string {

	tlen := len(text)
	reversed := make([]int32, tlen, tlen)

	for i, nucle := range text {
		if nucle == 65 {reversed[tlen-i-2] = 84}	// if nucleotide is "A" -> "T"
		if nucle == 67 {reversed[tlen-i-2] = 71}	// if nucleotide is "C" -> "G"
		if nucle == 71 {reversed[tlen-i-2] = 67}	// if nucleotide is "G" -> "C"
		if nucle == 84 {reversed[tlen-i-2] = 65}	// if nucleotide is "T" -> "A"
	}

	return string(reversed)

	
}

