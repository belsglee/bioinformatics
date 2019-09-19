package biolib

import (
	"fmt"
	"os"
	"bufio"
)

func main() {

	fmt.Println(Readfile("filename"))

}

func Readfile(name string) []string {

	f, _ := os.Open(name)

	defer f.Close()

	var ret = []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		ret = append(ret, scanner.Text())
	}

	return ret
}

/*func Motifenmeration(dna []string, k, d int) []string{

	var candidates = []string{}

	for _, pattern range dna {
		for i:=0; i < len(pattern); i++ {
//			for _, Neighbors(pattern, d)
		candidates = append(candidate, Freq_Words_W_Mismatches(pattern, k, d))
	}

	return remove_dupe(candidates)

}*/
