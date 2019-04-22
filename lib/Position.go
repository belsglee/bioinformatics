// This is code challenge as a part of coursera 'bioinformatics specialized course'
// created by Phillip, Pavel (UC San Diego)
// Week 1, Chapter 1, Part 1, Code Challenge 4
// finding position of pattern in text

// input : pattern and text
// ouput : posion

package bioinformatics
import (
	"bioinformatics/lib"
)

func Position(pattern, text string) []int {

	var plen = len(pattern)
	var where []int

	for i := 0; i <= len(text) - plen; i++ {
		if pattern == text[i:i+plen] {
			where = append(where, i)
		}
	}

	return where

}

