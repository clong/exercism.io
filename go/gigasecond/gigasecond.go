package gigasecond

import "time"

const testVersion = 4 // find the value in gigasecond_test.go

// The test case pre-parses the date into a time.Time object
// This function just adds a gigasecond to the input time and returns
// that value.
func AddGigasecond(inputTime time.Time) time.Time {
	      return inputTime.Add(time.Duration(1000000000) * time.Second)
}
