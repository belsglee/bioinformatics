def ReverseComplement(Pattern):
	return Complement(Pattern[::-1])

def Complement(Pattern):
	s = ""
	for nucle in Pattern:
		if nucle == "A": s += "T"
		if nucle == "C": s += "G"
		if nucle == "G": s += "C"
		if nucle == "T": s += "A"
	return s
