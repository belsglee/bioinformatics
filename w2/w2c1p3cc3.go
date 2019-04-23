package main

import (
	"fmt"
//	"io/ioutil"
)

var p = fmt.Println

func main () {
	
/*	var filename = "1.4cc2"
	f, _ := ioutil.ReadFile(filename)
	
	var filename2 = "1.4cc2_2"
	f2, _ := ioutil.ReadFile(filename2)*/

	sample := "TACGCATTACAAAGCACA"
	match := "AA"

	p(app_pattern_match(match, sample, 1))
	
}

func app_pattern_match (pattern, text string, mismatches int) []int {
	
	occurs := []int{}

	plen := len(pattern)

	for i:=0; i <= len(text) - plen; i++{
		p(hamming_distance(pattern, text[i:i+plen]))
		p(pattern,text[i:i+plen-1])
		if hamming_distance(pattern, text[i:i+plen]) <= mismatches { 
			occurs = append(occurs, i)
		}
	}
	return occurs
}


func hamming_distance(text, compare string) int {

	hamming := 0		

	for i, nucle := range text {
		if byte(nucle) != compare[i] { hamming++ }
	}

	return hamming

}

