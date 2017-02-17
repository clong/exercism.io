// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import "fmt"
import "time"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// A global var containing the time layout to use in time.Parse()
var TimeLayout = "15:04"

// Clock contains Hour and Minute Elements
type Clock struct {
	Hour, Min int
}

// Provided an hour and minute integer, this function returns a Clock
// struct converted to a legitimate time. Numbers outside the range of
// normal time may be provided and are automatically converted.
func New(hour, minute int) Clock {
  // If a negative hour value is provided, i.e. -10, take the modulus and
  // add it to 24. (24 + -14 = 10)
	if hour < 0 {
		hour = 24 + (hour % 24)
	}
  // If an hour larger than or equal to 24 is provided, take the modulus of
  // 24. Overflow is not factored in, so values of 48, 72, etc still equal zero
	if hour >= 24 {
		hour = hour % 24
	}
  // If a negative or large minute value is provided, generate a timestamp to
  // calculate the hourly overflow/underflow.
	if minute >= 60 || minute < 0 {
		initialTimestamp, err := time.Parse(TimeLayout, fmt.Sprintf("%02d:00", hour))
		if err != nil {
			fmt.Println(err)
		}
		timeAfterAddition := initialTimestamp.Add(time.Duration(minute) * time.Minute)
		hour = timeAfterAddition.Hour()
		minute = timeAfterAddition.Minute()
	}
	return Clock{Hour: hour, Min: minute}
}

// Returns a string representation of the clock struct
// %02d ensures the leading zeros will be shown for single digit ints
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Min)
}

// Adds the specified number of minutes to the time derived from a Clock struct
func (c Clock) Add(minutes int) Clock {
	initialTimestamp, err := time.Parse(
		TimeLayout,
		fmt.Sprintf("%02d:%02d", c.Hour, c.Min),
	)
	if err != nil {
		fmt.Println(err)
	}
	timeAfterAddition := initialTimestamp.Add(time.Duration(minutes) * time.Minute)
	return Clock{Hour: timeAfterAddition.Hour(), Min: timeAfterAddition.Minute()}
}
