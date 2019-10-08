// This is code challenge as a part of coursera 'bioinformatics specialized course'
// created by Phillip, Pavel (UC San Diego)
// Week 1, Chapter 1, Part 1, Code Challenge 5

package main
import (
	"fmt"
	"io/ioutil"
)

var p = fmt.Println

func main() {

	var filename = "test5"
	f, _ := ioutil.ReadFile(filename)
	sample := string(f)
	var fuck []string

	for i:=0; i<len(sample)-498; i++ {
		for _, you := range(FrequentWords(sample[i:i+497], 11)){
			fuck = append(fuck, you)}
	}

	p(unique(fuck))

}

func FrequentWords(text string, k int) []string{

	var Pattern string
	var patterns []string
	var count int
	min_count := 20

	for i := 0; i <= len(text) - k; i++{
		Pattern = text[i:k]
		count = PatternCount(text, Pattern)

		if count >= min_count {
			patterns = append(patterns, Pattern)
		}
		k++
	}

	return(unique(patterns))

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
	var plen 	= len(pattern)

	for i := 0; i <= len(text) - plen; i++{
		if pattern == text[i:i+plen] {
//			p(pattern)
			count++
		}
	}

	return count
}
	
