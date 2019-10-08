def pattern_matching(pattern, genome):
	matching_list = []
	plen = len(pattern)
	for i in range(len(genome)-plen+1):
		if pattern == genome[i:i+plen]:
			matching_list.append(i)
	return matching_list

