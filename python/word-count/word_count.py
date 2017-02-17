import string

def word_count(phrase):
    needed_decode = False
    # Standardize by lowering
    phrase = phrase.lower()

    # Check if the string needs to be decoded
    try:
        if phrase.decode('ascii'):
            pass
    except:
        phrase = phrase.decode('utf_8')
        needed_decode = True

    # Blow up all chars into list
    phrase = list(phrase)

    # Replace all non-alpha/numeric chars with spaces
    for i, char in enumerate(phrase):
        if needed_decode:
            if char.isalpha() is False and char.isnumeric() is False:
                phrase[i] = ' '
        else:
            if char.isalpha() is False and unicode(char, 'utf-8').isnumeric() is False:
                phrase[i] = ' '

    phrase = ''.join(phrase)
    phrase_list = phrase.split(' ')
    print phrase_list
    word_count_dict = {}
    for word in phrase_list:
        # Unicode section
        if needed_decode:
            # Make sure the word is alphanumeric
            if word.isalpha() is True or word.isnumeric() is True:
                # If the word isn't in the dict yet, add it
                if word not in word_count_dict:
                    word_count_dict[word] = 1
                    # else increase count
                else:
                    word_count_dict[word] += 1
            else:
                continue
        # Non-unicode section
        else:
            # Make sure the word is alphanumeric
            if word.isalpha() is True or unicode(word, 'utf-8').isnumeric() is True:
                # If the word isn't in the dict yet, add it
                if word not in word_count_dict:
                    word_count_dict[word] = 1
                # else increase count
                else:
                    word_count_dict[word] += 1
            else:
                continue
    return word_count_dict
