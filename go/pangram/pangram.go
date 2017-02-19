// Package determines whether or not something is a pangram.
// A pangram is a sentence using every letter of the alphabet at least once.
package pangram

import "strings"

const testVersion = 1

// Returns true if the provided string is a pangram
func IsPangram(input string) bool {
  // Input string has to contain at least 26 chars to qualify
  if len(input) < 26 {
    return false
  }

  // Normalize the string by making it lowercase
  input = strings.ToLower(input)

  alphabet := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
   // Loop through each letter of the alphabet and count the instances of each
   // letter. If any of the counts are zero, the string is not a pangram.
    for i := range alphabet {
      if strings.Count(input, alphabet[i]) == 0 {
        return false
      }
    }
    return true
  }
