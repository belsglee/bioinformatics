package main

import (
	"biolib"
	"fmt"
)

var p = fmt.Println

func main(){

/*	var profile = [][]float64{
//		{0.0, 0.0, 1.0, 0.0},
//		{0.0, 0.0, 1.0, 0.0},
//		{0.0, 1.0, 0.0, 0.0}
		{0.0, 0.0, 0.0},
		{0.0, 0.0, 1.0},
		{1.0, 1.0, 0.0},
		{0.0, 0.0, 0.0}}				// should i change this like anyone else?	
*/

	var profile = [][]float64{}

	for i, kmer := range first_str {

	for i:=0; i<=len(		
		for _, text := range Dna {
			best_pt := biolib.Profile_most_probable_kmer(text, k, profile)
			if best_pt == "" {best_pt = tex	t[:k]}
			profile = Add_profile(profile, best_pt)
		}

}

func Add_profile ([][]float64 profile, string best) [][]float64 {	

}
