def median_string(dna, k):
	distance = 9999 # inifinite number.

	for i in range(4**k):
		pattern = number2pattern(i, k):
		if distance > dist_pns(pattern, dna):
			distance = dist_pns(pattern, dna)
			median = pattern

	return median

def number2pattern(num, k):
	if k == 1:
		return number2symbol(num)
	prefixnum = num / 4
	r = n % 4
	symbol = number2symbol(r)
	prefix_pattern = number2pattern(prefixnum, k-1)

	return prefix_pattern + symbol

def number2symbol(n):

	if n == 0:
		symbol = "A"
	if n == 1:
		symbol = "C"
	if n == 2:
		symbol = "G"
	if n == 3:
		symbol = "T"
	
	return symbol


def distance_btw_pattern_and_string(pattern, dna):

	k = len(pattern)
	distance, hamming_dist = 0, 0

	for text in dna:
		 hamming_dist = 9999
			for kmer in list_pattern(text, k):
				if hamming_dist > hamming_dist(pattern, kmer):
					hamming_dist = hamming_dist(pattern, kmer)
			distance += hamming_dist

	return distance

def hamming_dist (text, compare):

	hamming_distance = 0
	
	for i, nucleotide in text:	
		if nucleotide != compare[i]:
			hamming_distance += 1

	return hamming_distance 

def list_pattern(text, k):

	dnaset = []

	for i in range(text-k):
		dnaset.append(test[i:i+k])

	return dnaset

