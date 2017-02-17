// Package determines whether or not a given year is a leap year
package leap

const testVersion = 3

// Determines whether or not a supplied year is a leap year or not.
// Returns true if it's a leap year and false if not
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	} else {
		return false
	}
}
