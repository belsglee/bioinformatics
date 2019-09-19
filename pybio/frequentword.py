def frequentwords (text, k):
	previous_count = 0
	patterns = []

	for i in range(len(text)-k):
		pattern = text[i:k]
		count = patterncount(text)

		if count == previous_count:
			patterns.append(pattern)

		if count > previous_count:
			patterns = []
			patterns.append(pattern)

print(frequentwords("CGCGCG", 2))
