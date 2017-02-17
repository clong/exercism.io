def is_pangram(phrase):
    alphabet = ['a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z']
    phrase = phrase.lower()
    phrase_list = list(phrase)
    print phrase_list
    output_list = []
    for letter in phrase_list:
        if letter not in output_list:
            if letter in alphabet:
                output_list.append(letter)
    output_list.sort()
    print output_list
    if output_list == alphabet:
        return True
    else:
        return False
