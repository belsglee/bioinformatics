def approximate_pattern_match(genome, pattern, allow_distance):
	positions = [] # initializing list of position
	plen = len(pattern)

	for i in range(len(genome)-plen+1):
		compared_genome = genome[i:i+plen]
		actual_distance = hamming_distance(pattern, compared_genome)
		
		if actual_distance <= allow_distance:
			positions.append(i)

	return positions
