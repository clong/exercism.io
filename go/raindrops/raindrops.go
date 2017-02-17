package raindrops

import (
  "fmt"
  "bytes"
)

const testVersion = 2

// Converts an integer into a string based on the integers factors
func Convert(inputNumber int) string {
  var buffer bytes.Buffer
  if inputNumber % 3 == 0 {
    buffer.WriteString("Pling")
  }
  if inputNumber % 5 == 0 {
    buffer.WriteString("Plang")
  }
  if inputNumber % 7 == 0 {
    buffer.WriteString("Plong")
  }
  if buffer.Len() == 0 {
    return fmt.Sprintf("%d", inputNumber)
  }
  return buffer.String()
}

// BenchmarkConvert-4   	 1000000	      1781 ns/op
