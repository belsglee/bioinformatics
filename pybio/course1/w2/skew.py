def skew (genome):
	candidate, min_num = 0, 0
	skewed = []

	for i, nucleotide in enumerate(genome):
		if nucleotide == "C": candidate -= 1
		if nucleotide == "G": candidate += 1
		if candidate < min_num:
			min_num = candidate
			skewed = []	#initialize list
		if candidate == min_num:
			skewed.append(i+1)
	return skewed
