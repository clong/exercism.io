// Bob takes input and gives out responses based on whether he interprets the
// input to be a question, yelling, empty, or anything else
package bob

import (
  "regexp"
  "strings"
)

const testVersion = 2 // same as targetTestVersion

// Input is a question if it ends in a question mark.
func checkQuestion(input string) bool {
  re := regexp.MustCompile(`.*\?$`)
  return re.MatchString(input)
}

// Yelling is defined by a string containing an exclamation mark or consisting
// of all caps
func checkYelling(input string) bool {
  re := regexp.MustCompile("[a-zA-Z0-9!]")
  re2 := regexp.MustCompile("[a-zA-z]")
  if re2.MatchString(input) {
    if input == strings.ToUpper(input) && re.MatchString(input) {
      return true
    }
    if strings.Count(input, "!") >= 2 {
      return true
    }
  }
  return false
}

// For you to have said something, it must contain an alphanumic character
func checkSayNothing(input string) bool {
  re := regexp.MustCompile("[a-zA-Z0-9]")
  return re.MatchString(input)
}

// This function determines what type of response is warranted by Bob.
func Hey(input string) string {
  input = strings.TrimSpace(input)
  if checkYelling(input) {
    return "Whoa, chill out!"
  } else if checkQuestion(input) {
    return "Sure."
  } else if !checkSayNothing(input) {
    return "Fine. Be that way!"
  } else {
    return "Whatever."
  }
}
