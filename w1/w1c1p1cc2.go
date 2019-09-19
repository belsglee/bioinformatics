// This is code challenge as a part of coursera 'bioinformatics specialized course'
// created by Phillip, Pavel (UC San Diego)
// Week 1, Chapter 1, Part 1, Code Challenge 2

package main

import (
	"fmt"
	"io/ioutil"
)

var p = fmt.Println

func main() {

	var filename = "quiz3_1"
	f, _ := ioutil.ReadFile(filename)
	sample := string(f)

	FrequentWords(sample, 2)

}

func FrequentWords(text string, k int) {

	var Pattern string
	var patterns []string

	old_count := 0

	for i := 0; i <= len(text)-k; i++ {
		Pattern = text[i:k]
		count := PatternCount(text, Pattern)

		if count == old_count {
			patterns = append(patterns, Pattern)

		}

		if count > old_count {
			patterns = nil
			patterns = append(patterns, Pattern)
			old_count = count
		}
		k++
	}
	p(unique(patterns))
}

func unique(Slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range Slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func PatternCount(text, pattern string) int {

	var count = 0
	var plen = len(pattern)

	for i := 0; i <= len(text)-plen; i++ {
		if pattern == text[i:i+plen] {
			count++
		}
	}

	return count
}
