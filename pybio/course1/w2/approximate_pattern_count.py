def approximate_pattern_count(pattern, genome, expected_distance):
	count = 0 # initializing count variable
	plen = len(pattern)
	for i in range(len(genome)-plen+1):
		actual_distance = hamming_distance(genome[i:i+plen], pattern) 
		if actual_distance <= expected_distance:
			count += 1
	return count
