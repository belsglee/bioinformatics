def distance_between_pattern_and_string(pattern, dna):
	plen = len(pattern)
	distance, hamming_distance = 0, 0
	for string in dna:
		hamming_distance = 9999 
		for kmer in list_pattern(string, plen):
			if hamming_distance > hamming_dist(pattern, kmer):
				hamming_distance = hamming_dist(pattern, kmer)
		distance += hamming_distance

	return distance
