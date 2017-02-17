def distance(s1, s2):
        distance = 0
        s1_list = list(s1)
        s2_list = list(s2)
        i = 0
        for nucleotide in s1_list:
            if s1_list[i] != s2_list[i]:
                distance += 1
            i += 1
        return distance
