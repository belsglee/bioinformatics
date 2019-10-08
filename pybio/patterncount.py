def patterncount(text, pattern):

	count = 0
	plen = len(pattern)

	for i in range(len(text) - 1):
		if pattern == text[i:i+plen]:
			count += 1

	return count

print(patterncount("CGCGABC", "CGCG"))
