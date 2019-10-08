from random import *

def randomized_motif_search(dna, k, t):
	
	# randomly select k-mers motifs in each string from dna
	# (motif.1, ..., motif.t), t is number of strings in dna array

	motifs = []

	for each in dna:
		i = randint(0, len(each)-k)
		motifs.append(each[i:i+k])
	
	print(motifs)

	bestmotifs = init_motifs

	while True:
		profile = form_profile(motifs)
		motifs = form_motifs(profile, dna)
		if score(motifs) < score(bestmotifs)
			bestmotifs = motifs

		else
			return bestmotifs

def form_profile(motifs):
	profile = {"A" : [], "C": [], "G": [], "T": []}
	
	for motif in motifs:
		for i, nucle in enumerate(motif):
			if nucle == "A": profile["A"][i] += 1
			if nucle == "C": profile["C"][i] += 1
			if nucle == "G": profile["G"][i] += 1
			if nucle == "T": profile["T"][i] += 1
	
	for nucle in profile:
		for i, element in enumerate(profile[nucle]):
			if element != 0:
				profile[nucle][i] = element / len(motifs)

	return profile



def form_motifs(dna, k}
		, profile):

	motifs = []

	for motif in Dna:
		best_motif = profile_most_probable_kmer(motif, k, profile)
		if best_motif == "":
			best_motif == motif[:k]
		motifs.appends(best_motif)

	return motifs


def profile_most_probable_kmer(dna, k, profile):
	best_kmer = ""
	best_probability = 0.0

	for i in range(len(dna)-k+1):
		kmer = dna[i:i+k]
		calculated = calculate_probability(kmer, profile)
		}
if calculated > best_probability:
			best_kmer = kmer
			best_probability = calculated
	
	return best_kmer

def calculate_probability(kmer, profile):
	t = []
	for i, nucle in enumerate(kmer):
		if nucle == "A": t.append(profile["A"][i])
		if nucle == "C": t.append(profile["C"][i])
		if nucle == "G": t.append(profile["G"][i])
		if nucle == "T": t.append(profile["T"][i])

	return multiply(t)
	# or simply use numpy prod function.

''' from numpy import prod
	prod(t) '''


def multiply(profile):
	m = 1.0

	for i in profile:
		m = m * i

	return m 

def score(motifs):

	
