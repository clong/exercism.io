def to_rna(dna):
    pairs = {'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
    dna_list = list(dna)
    for i, nucleotide in enumerate(dna_list):
        dna_list[i] = pairs[nucleotide]
    return ''.join(dna_list)
