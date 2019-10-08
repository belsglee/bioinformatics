package main

import "fmt"

var p = fmt.Println


func main() {


	p(PatternToNumber("AC"))


}

func PatternToNumber (pattern string) int {
	
	plen := len(pattern)

	if plen == 0 {return 0}

	symbol := pattern[plen-1]
	prefix := pattern[0:plen-1]

	p(string(symbol), prefix)
	return 4 * PatternToNumber(prefix) + SymbolToNumber(string(symbol))

}

func SymbolToNumber (symbol string) int {

	num := 0

	if symbol == "A" {num = 0}
	if symbol == "C" {num = 1}
	if symbol == "G" {num = 2}
	if symbol == "T" {num = 3}
	
	return num

	}
