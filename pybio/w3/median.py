def median_string(dna, k):
	distance = 999999 # inifinite number.
	median = []
	for i in range(0, 4**k):
		pattern = number2pattern(i, k)
		dis_pns = distance_btw_pattern_and_string(pattern, dna)
		if i % 100 == 0:
			input("check point: please enter")
		print(i, pattern, dis_pns, distance)
		if distance > dis_pns:
			distance = dis_pns
			median = [pattern]

		if distance == dis_pns:
			median.append(pattern)
	
	return median


def number2pattern(num, k):
	if k == 1:
		return number2symbol(num)
	prefixnum = num // 4
	r = num % 4
	symbol = number2symbol(r)
	prefix_pattern = number2pattern(prefixnum, k-1)
	
	return prefix_pattern + symbol

def number2symbol(n):
	if n == 0: 
		return "A"
	if n == 1:
		return "C"
	if n == 2:
		return "G"
	if n == 3:
		return "T"

def distance_btw_pattern_and_string(pattern, dna):
	k = len(pattern)
	distance = 0

	for text in dna:
		hamming_distance = 999999
		for i in range(len(text)-k+1):
			if hamming_distance > hamming_dist(pattern, text[i:i+k]):
				hamming_distance = hamming_dist(pattern, text[i:i+k])
		distance += hamming_distance
	
	return distance

def hamming_dist (text, compare):
	hamming_distance = 0
	for i, nucleotide in enumerate(text):
		if nucleotide != compare[i]:
			hamming_distance += 1
	return hamming_distance

a = "CTCGATGAGTAGGAAAGTAGTTTCACTGGGCGAACCACCCCGGCGCTAATCCTAGTGCCC"
b = "GCAATCCTACCCGAGGCCACATATCAGTAGGAACTAGAACCACCACGGGTGGCTAGTTTC"
c = "GGTGTTGAACCACGGGGTTAGTTTCATCTATTGTAGGAATCGGCTTCAAATCCTACACAG"

dna = [a, b, c]

print(median_string(dna, 7))
