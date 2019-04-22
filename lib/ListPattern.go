package biolib

func List_pattern(text string, k int) []string {

	var set []string
	
	for i:=0; i<=len(text)-k; i++ {
		set = append(set, text[i:i+k])
	}

	return set

}
