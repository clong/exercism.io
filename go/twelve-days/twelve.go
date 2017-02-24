// Prints the lyrics from "The Twelve Days of Christmas"
package twelve

import (
	"bytes"
	"fmt"
)

const testVersion = 1

// Returns the verse, provided the verse number
func Verse(verseNum int) string {
	var buffer bytes.Buffer
	days := []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	gifts := []string{"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens", "four Calling Birds", "five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}

	// If it's the first verse, just return it. No string building needed
	if verseNum == 1 {
		return fmt.Sprintf(
			"On the %s day of Christmas my true love gave to me, %s.",
			days[0],
			gifts[0],
		)
	}

	// We enter this code if it's higher than the first verse
	buffer.WriteString(
		fmt.Sprintf(
			"On the %s day of Christmas my true love gave to me, %s, ",
			days[verseNum-1],
			gifts[verseNum-1],
		),
	)
	for countdown := verseNum - 1; countdown > 1; countdown-- {
		buffer.WriteString(fmt.Sprintf("%s, ", gifts[countdown-1]))
	}

	// Once the verseNum countdown is back down to one, append the partridge
	buffer.WriteString(fmt.Sprintf("and %s.", gifts[0]))
	return buffer.String()
}

// Loop through each verse, add it to the buffer. Finally, return the full song
func Song() string {
	var buffer bytes.Buffer
	for i := 1; i < 13; i++ {
		buffer.WriteString(fmt.Sprintf("%s\n", Verse(i)))
	}
	return buffer.String()
}
