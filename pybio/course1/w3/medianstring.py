def medianstring(dna, k):
	distance = 9999 # inifinite number.

	for kmer in makeallkmer(k):
		if distance > dist_pns(kmer, dna):
			distance = dist_pns(kmer, dna)
			median = kmer

	return median

def makeallkmer(k):

	dnaset = []
	pn = 2 ** (k*2)

	for i in range(pn):
		dnaset.append(number2pattern(i, k))

	return dnaset

def number2pattern(num, k):
	if k == 1: return number2symbol(num)
	prefixnum = num/4
	r = n % 4
	symbol = number2symbol
	prefix_pattern = number2pattern(prefixnum, k-1)

	return prefix_pattern + symbol

def number2symbol(n):

	if n == 0 : sym = "A"
	if n == 1 : sym = "C"
	if n == 2 : sym = "G"
	if n == 3 : sym = "T"

	return sym

def dist_pns(pattern, dna):

	k = len(pattern)
	distance, hamming_dist = 0

	for text in dna:
		 hamming_dist = 99
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

