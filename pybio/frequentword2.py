def patterncount(text, pattern):

        count = 0
        plen = len(pattern)

        for i in range(len(text) - 1):
                if pattern == text[i:i+plen]:
                        count += 1

        return count


def frequentwords (text, k):
	previous_count, count = 0, 0
	patterns = []

	for i in range(len(text)-k):
		pattern = text[i:k]
		count = patterncount(text, pattern)

		if count == previous_count:
			patterns.append(pattern)

		if count > previous_count:
			patterns = []
			patterns.append(pattern)

		k += 1

	return patterns

print(frequentwords("CGCGCG", 2))
