text = """ATGCG
GCATG
CATGC
AGGCA
GGCAT
GGCAC"""

a = text.split('\n')

p_dict = {}

for i, p in enumerate(a):
    for p2 in a[i+1:]:
        p_pre = p[:-1]
        p_post = p[1:]
        p2_pre = p2[:-1]
        p2_post = p2[1:]

        if p_post == p2_pre:
            if p in p_dict.keys():
                p_dict[p].append(p2)
            else:
                p_dict[p] = [p2]
        elif p_pre == p2_post:
            if p2 in p_dict.keys():
                p_dict[p2].append(p)
            else:
                p_dict[p2] = [p]

for key, value in p_dict.items():
    print(key, '->', ', '.join(value))
