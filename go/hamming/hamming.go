package hamming

import "strings"
import "errors"

const testVersion = 5

func Distance(strand1, strand2 string) (int, error) {
  hammingDistance := 0
  if len(strand1) != len(strand2) {
    return 0, errors.New("Error: strands are not equal in length")
  }
  // If strings are identical, the distance is zero
	if strings.Compare(strand1, strand2) == 0 {
    return 0, nil
  }
  // Iterate through both strings and compare the characters at each index.
  // If the characters are not equal, increment the hamming distance.
  for i := 0; i < len(strand1); i++ {
    if strand1[i] != strand2[i] {
      hammingDistance++
    }
  }
  return hammingDistance, nil
}
