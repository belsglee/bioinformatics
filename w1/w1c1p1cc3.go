// This is code challenge as a part of coursera 'bioinformatics specialized course'
// created by Phillip, Pavel (UC San Diego)
// Week 1, Chapter 1, Part 1, Code Challenge 3

package main
import (
	"fmt"
	"io/ioutil"
)

var p = fmt.Println

func main() {

	var filename = "quiz4"
	f, _ := ioutil.ReadFile(filename)
	sample := string(f)

	tlen := len(sample)
	reversed := make([]int32, tlen, tlen)

	for i, nucle := range sample {
		if nucle == 65 {reversed[tlen-i-2] = 84}	// if nucleotide is "A"
		if nucle == 67 {reversed[tlen-i-2] = 71}
		if nucle == 71 {reversed[tlen-i-2] = 67}
		if nucle == 84 {reversed[tlen-i-2] = 65}
	}

	p(string(reversed))

	
}

