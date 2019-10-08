// This is code challenge as a part of coursera 'bioinformatics specialized course'
// created by Phillip, Pavel (UC San Diego)
// Week 1, Chapter 1, Part 1, Code Challenge 1

package main
import (
	"fmt"
	"io/ioutil"
)

var p = fmt.Println

func main() {

	var filename = "quiz2"
	f, _ := ioutil.ReadFile(filename)

	sample := string(f)

	p(PatternCount(sample, "CGCG"))
	
}

func PatternCount(text, pattern string) int {

	var count = 0
	var plen 	= len(pattern)

	for i := 0; i <= len(text) - plen; i++{
		if pattern == text[i:i+plen] {
			count++
		}
	}

	return count
}
	
