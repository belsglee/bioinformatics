// This is code challenge as a part of coursera 'bioinformatics specialized course'
// created by Phillip, Pavel (UC San Diego)
// Week 1, Chapter 1, Part 1, Code Challenge 4
// finding position of pattern in text

// input : pattern and text
// ouput : posion

package main
import (
	"fmt"
	"io/ioutil"
)

var p = fmt.Println

func main() {

	var filename = "quiz5"
	f, _ := ioutil.ReadFile(filename)
	sample := string(f)
	
	Position("ATA", sample)

}

func Position(pattern, text string) {

	var plen = len(pattern)

	for i := 0; i <= len(text) - plen; i++ {
		if pattern == text[i:i+plen] {
			fmt.Print(i," ")
		}
	}

}

