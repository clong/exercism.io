// Generates acronyms given an input phrase
package acronym

import (
	"regexp"
	"strings"
	"text/scanner"
	"unicode"
)

const testVersion = 2

// Replace [:-,"] chars from the input phrase and replace them with spaces
func RemoveDelimiters(phrase string) string {
	re := regexp.MustCompile(`[\:\-\,\"]`)
	return re.ReplaceAllString(phrase, " ")
}

// Takes a slice of strings and joins them. Also removes double quotes
func RemoveDoubleQuotesAndJoin(acronym []string) string {
	return strings.Replace(strings.Join(acronym, ""), "\"", "", -1)
}

// Returns true if the supplied string contains all uppercase characters
func WordIsUppercase(word string) bool {
	if word == strings.ToUpper(word) {
		return true
	}
	return false
}

// Returns true if the supplied string contains all lowercase characters
func WordIsLowercase(word string) bool {
	if word == strings.ToLower(word) {
		return true
	}
	return false
}

// The rules for creating an acronym don't seem to be very clearcut. From what I
// can discern, the rules for generating an acronym are as follows:
// 1. Every character after a space gets an acronym letter regardless of
// capitalization (i.e. Ruby on Rails)
// 2. If the word contains all capitals, only the first letter of the word gets
// added to the acronym (i.e. PHP)
// 3. If a word contains a mix of caps and lowercase, all of the capital letters
// from the word get added to the acronym (i.e. HyperText)
func Abbreviate(phrase string) string {
	var acronym []string
	// Replace all punctuation with spaces
	phrase = RemoveDelimiters(phrase)
	// Break up the phrase by spaces into a slice of "words"
	phraseSlice := strings.Split(phrase, " ")

	for i := 0; i < len(phraseSlice); i++ {
		// Check if word is all caps or lowercase.
		// If it is, only the first letter gets used
		if WordIsUppercase(phraseSlice[i]) || WordIsLowercase(phraseSlice[i]) {
			for _, char := range phraseSlice[i] {
				acronym = append(acronym, strings.ToUpper(scanner.TokenString(char)))
				break
			}
		} else {
			// If the word isn't all caps, we'll extract all of the capital letters
			// from the word and add them to the acronym.
			for _, char := range phraseSlice[i] {
				if !unicode.IsLower(char) {
					acronym = append(acronym, scanner.TokenString(char))
				}
			}
		}
	}

	// At this point, our acronym is a slice containing runes. We need to remove
	// the quotes and turn that slice into a string.
	return RemoveDoubleQuotesAndJoin(acronym)
}
