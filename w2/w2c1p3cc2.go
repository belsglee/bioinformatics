//Mismatch Proglem: Compare Two string, find out how many position is mismatched
//INPUT: Two DNA string Genome of equal length
//OUTPUT: The Hamming distance between these strings

package main

import (
	"fmt"
//	"io/ioutil"
)

var p = fmt.Println

func main () {
	
/*	var filename = "1.4cc1"
	f, _ := ioutil.ReadFile(filename)
	
	var filename2 = "1.4cc1_2"
	f2, _ := ioutil.ReadFile(filename2)

	sample := string(f)*/

	ori := "CTACAGCAATACGATCATATGCGGATCCGCAGTGGCCGGTAGACACACGT"
	compare := "CTACCCCGCTGCTCAATGACCGGGACTAAAGAGGCGAAGATTATGGTGTG"

	p(len(ori))

	p(hamming_distance(ori, compare))
	
}

func hamming_distance (text, compare string) int {
	
	hamming :=0

	for i, nucle := range text {
		if byte(nucle) != compare[i] { hamming++ }
	}

	return hamming
	

}
