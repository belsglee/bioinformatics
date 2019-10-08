# need to implemented

def clump(genome, k, L, t):

	for i in range(L):
		for j in range(frequent_words(genome[i

def frequent_words(text, k):

	patterns = []

	for i in range(len(text)-k+1):
		pattern = text[i:k]
		count = patterncount(text, pattern)

		if count >= min_count {
			if pattern not in patterns:
				patterns.append(pattern)
		
		k += 1

	return patterns
