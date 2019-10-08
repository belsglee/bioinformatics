def hamming_distance(p, q):
	distance = 0
	for i, nucleotide in enumerate(p):
		if nucleotide != q[i]:
			distance += 1
	return distance
