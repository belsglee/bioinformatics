//Minimum Skew Problem: Find a position in a genome where the skew diagram attains a minimum.
//INPUT: A DNA string Genome.
//OUTPUT: All integer(s) i minimizing Skewi (Genome) among all values of i (from 0 to |Genome|).

package main

import (
	"fmt"
//	"io/ioutil"
)

var p = fmt.Println

func main () {
	
/*	var filename = "1.3cc1"
	f, _ := ioutil.ReadFile(filename)
	sample := string(f)*/

	sample := "CATTCCAGTACTTCGATGATGGCGTGAAGA"

	p(skew(sample))
	
}

func skew (test string) []int {
	
	candidate := 0
	min := 0
	list := []int{}

	for i, nucle := range test {
		if nucle == 67 {candidate -= 1} // if nucleotide is "C" -1 
		if nucle == 71 {candidate += 1} // if nucleotide is "G" +1
		if candidate < min {
				min = candidate
				list = nil
		}

		if candidate == min {
	 			list = append (list, i+1)
		}
	}
	return list

}
