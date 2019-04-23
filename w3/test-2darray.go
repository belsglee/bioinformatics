package main

import "fmt"

var p = fmt.Println

func main(){

	twod := [][]int{{1,2,3}, {4,5,6}}
	p(twod)	

	twod[0][2]++

	p(twod)

	


}
