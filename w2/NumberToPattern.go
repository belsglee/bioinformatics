package main

import "fmt"

var p = fmt.Println


func main() {

	p(NumberToPattern(6899, 8))

}

func NumberToPattern (Number, k int) string {
	
	if k == 1 {
		return NumberToSymbol(Number)}
	
	prefixnum := int(Number / 4)
	r := Number % 4
	symbol := NumberToSymbol(r)
	prefixPattern := NumberToPattern(prefixnum, k-1)

	p(symbol, prefixPattern)

	return prefixPattern + symbol

}

func NumberToSymbol (num int) string {

	sym := ""

	if num == 0 {sym = "A"}
	if num == 1 {sym = "C"}
	if num == 2 {sym = "G"}
	if num == 3 {sym = "T"}
	
	return sym

}
