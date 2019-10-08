def genome_path(patterns):
    string = patterns[0][:-1]
    for pattern in patterns:
        string += patterns[-1]
    return string
