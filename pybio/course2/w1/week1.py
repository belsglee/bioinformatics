#import sys
#k = int(sys.stdin.readline())
#text = sys.stdin.readline()

def main():
    command = input(
    """Enter your program name
    1. string reconstruction
    2. genome path
    3. overlap graph
    4. De Bruijn graph\n""")

    f = open(input("input file name: ") + '.txt', 'r')

    if command == '1':
        k = int(f.readline())
        text = f.readline().rstrip()
        compose(k, text)
        for kmer in compose(k, text):
            print (kmer)
    elif command == '2':
        patterns = [p.rstrip() for p in f.readlines()]
        f2 = open("genome_path_reseult.txt", 'w')
        f2.write(genome_path(patterns))
        f2.close()

    elif command == '3':
        patterns = [p.rstrip() for p in f.readlines()]
        p_dict = overlap_graph(patterns)
        f3 = open("overlap_graph_reseult.txt", 'w')
        for key, value in p_dict.items():
          f3.write(key, '->', ', '.join(value), '\n')
        f3.close()

    elif command == '4':
        k = int(f.readline())
        text = f.readline().rstrip()
        de_bruijn(k, text)

    f.close()


# composition function
def compose(k, text):
    # define substring as array
    substring = []
    for i in range(len(text)-k+1):
        substring.append(text[i:i+k])     
    
    # optional; sorted(substring)
    return substring

# genome_path
def genome_path(patterns):
    path = patterns[0][:-1]
    for pattern in patterns:
        print(pattern, path)
        path += pattern[-1]

    return path

def overlap_graph(patterns):

    p_dict = {}
    s_dict = {}

    for p in patterns:
        if p[:-1] not in p_dict.keys():
            p_dict[p[:-1]] = []
        p_dict[p[:-1]].append(p)

    for p in patterns:
        if p[1:] in p_dict.keys():
            suffix = p_dict[p[1:]]
            s_dict[p] = suffix 

    return s_dict

def de_bruijn(k, text):

    p_dict = {}

    for i in range(len(text)-k+1):
        p_dict
        text[i:i+k-1]

# version 2
'''
        p_post = patterns[0][1:]
        for p2 in patterns:
            p_pre = p2[:-1]
            if p_post == p_pre:
                if p not in p_dict.keys():
                    p_dict[p] = []
                p_dict[p].append(p2)
                if p2 in p_dict.keys():
                    patterns.remove(p)
                p_post = p2[1:]
            p = p2
'''

# previous version 
'''        for p2 in patterns[i+1:]:
            p_pre = p[:-1]
            p_post = p[1:]
            p2_pre = p2[:-1]
            p2_post = p2[1:]

            print(p, p2)

            if p_post == p2_pre:
                if p in p_dict.keys():
                    p_dict[p].append(p2)
                else:
                    p_dict[p] = [p2]
            elif p_pre == p2_post:
                if p2 in p_dict.keys():
                    p_dict[p2].append(p)
                else:
                    p_dict[p2] = [p]''' 

if __name__ == "__main__":
    main()
