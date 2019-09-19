package biolib

func Motifs (profile [][]float64, k int,  Dna []string) []string {

	var motifs []string
	var best_pt string
	var best_pb float64
	best_pb = 0.0

	for _, motif := range Dna {
		
		for i:=0; i<=len(motif)-k; i++{
			kmer := motif[i:i+k]
			if Calculate_probability(kmer, profile) > best_pb{
				best_pt = kmer
				best_pb = Calculate_probability(kmer, profile)
			}
		}
		if best_pt == "" {best_pt = motif[:k]}
		motifs = append(motifs, best_pt)
	}
	
	return motifs
}
