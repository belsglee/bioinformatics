package main

import (
	"fmt"
)

var p = fmt.Println

func main () {
	
/*	var filename = "1.4cc4"
	f, _ := ioutil.ReadFile(filename)*/

	var f = "TGCAT"
	var d = 2

	p(len(Neighbors(f, d)))

	var remove_me = []int{1,2,3}
	p(len(remove_me))

	
	
}


func Neighbors (pattern string, d int) []string {

	if d == 0 {return nil}
	if len(pattern) == 1 {
		return []string{"A", "C", "G", "T"}
		}

	Neighborhood := []string{}
	Suffix := pattern[1:]
	SuffixNeighbors := Neighbors(Suffix, d)

	for _, t := range SuffixNeighbors {

		if hamming_distance(Suffix, t) == d {
			Neighborhood = append(Neighborhood, (string(pattern[0]) + t))
		}

		if hamming_distance(Suffix, t) < d {
			Neighborhood = append(Neighborhood, ("A" + t))
			Neighborhood = append(Neighborhood, ("C" + t))
			Neighborhood = append(Neighborhood, ("G" + t))
			Neighborhood = append(Neighborhood, ("T" + t))
		}
	}
	return Neighborhood
}

func hamming_distance(text, compare string) int {

	hamming := 0		

	for i, nucle := range text {
		if byte(nucle) != compare[i] { hamming++ }
	}

	return hamming

}

