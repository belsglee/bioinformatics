#import sys
#k = int(sys.stdin.readline())
#text = sys.stdin.readline()

def main():
    command = input(
    """Enter your program name
    1. string reconstruction
    2. genome path
    3. overlap graph\n""")

    f = input("input file name: ")
    print(f)
#    f = open('dataset.txt', 'r')

    if command == 1:
        k = int(f.readline())
        text = f.readline().rstrip()
        compose(k, text)
        for kmer in compose(k, text):
            print (kmer)
    elif command == 2:
        patterns = f.readlines()
        print(patterns)
        print(genome_path(patterns))


# composition function
def compose(k, text):
    # define substring as array
    substring = []
    for i in range(len(text)-k+1):
        substring.append(text[i:i+k])     
    
    # optional; sorted(substring)
    return substring

def genome_path(patterns):

    path = patterns[0][:-1]
    for pattern in patterns:
        path += patterns[-1]

    return path

if __name__ == "__main__":
    main()
