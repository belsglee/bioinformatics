package biolib

import (
	"sort"
)

func Find_Neighbor(text string, k, d int) []string {

	Neighborhoods := []string{}

	for i := 0; i < len(text)-k; i++ {
		Neighborhoods = append(Neighborhoods, Neighbors(text[i:k+i], d)...)
	}

	return Neighborhoods
}

func Freq_Words_W_Mismatches(text string, k, d int) []string {

	Freq_Patterns := []string{}
	Neighborhoods := []string{}

	for i := 0; i < len(text)-k; i++ {
		Neighborhoods = append(Neighborhoods, Neighbors(text[i:k+i], d)...)
	}

	Index := []int{}
	Count := make([]int, len(Neighborhoods))

	for i, neighbor := range Neighborhoods {
		Index = append(Index, PatternToNumber(neighbor))
		Count[i] = 1
	}

	sort.Ints(Index)

	for i := 0; i < len(Neighborhoods)-1; i++ {
		if Index[i] == Index[i+1] {
			Count[i+1] = Count[i] + 1
		}
	}

	maxCount := Max(Count)

	for i := 0; i < len(Neighborhoods); i++ {
		if Count[i] == maxCount {
			Pattern := NumberToPattern(Index[i], k)
			Freq_Patterns = append(Freq_Patterns, Pattern)
		}
	}

	return Freq_Patterns
}

func Max(Count []int) int {

	var x int

	for _, c := range Count {
		if c > x {
			x = c
		}
	}
	return x
}

func NumberToPattern(n, k int) string {

	if k == 1 {
		return NumberToSymbol(n)
	}

	prefixn := int(n / 4)
	r := n % 4
	symbol := NumberToSymbol(r)
	prefixp := NumberToPattern(prefixn, k-1)

	return prefixp + symbol

}

func NumberToSymbol(n int) string {

	sym := ""

	if n == 0 {
		sym = "A"
	}
	if n == 1 {
		sym = "C"
	}
	if n == 2 {
		sym = "G"
	}
	if n == 3 {
		sym = "T"
	}

	return sym

}

func PatternToNumber(Pattern string) int {

	plen := len(Pattern)

	if plen == 0 {
		return 0
	}

	symbol := Pattern[plen-1]
	Prefix := Pattern[0 : plen-1]
	return 4*PatternToNumber(Prefix) + SymbolToNumber(string(symbol))

}

func SymbolToNumber(symbol string) int {

	num := 0

	if symbol == "A" {
		num = 0
	}
	if symbol == "C" {
		num = 1
	}
	if symbol == "G" {
		num = 2
	}
	if symbol == "T" {
		num = 3
	}

	return num
}

func Neighbors(pattern string, d int) []string {

	if d == 0 {
		return nil
	}
	if len(pattern) == 1 {
		return []string{"A", "C", "G", "T"}
	}

	Neighborhood := []string{}
	Suffix := pattern[1:]
	SuffixNeighbors := Neighbors(Suffix, d)

	for _, t := range SuffixNeighbors {

		if Hamming_distance(Suffix, t) == d {
			Neighborhood = append(Neighborhood, (string(pattern[0]) + t))
		}

		if Hamming_distance(Suffix, t) < d {
			Neighborhood = append(Neighborhood, ("A" + t))
			Neighborhood = append(Neighborhood, ("C" + t))
			Neighborhood = append(Neighborhood, ("G" + t))
			Neighborhood = append(Neighborhood, ("T" + t))
		}
	}
	return Neighborhood
}

func Hamming_distance(text, compare string) int {

	hamming := 0

	for i, nucle := range text {
		if byte(nucle) != compare[i] {
			hamming++
		}
	}

	return hamming

}
