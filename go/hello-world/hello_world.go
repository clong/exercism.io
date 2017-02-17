// Package greets a person by name if supplied to the program
package greeting

import (
	"os"
	"fmt"
)

const testVersion = 3

// Outputs "Hello, World!" if no name is suppled, otherwise outputs
// "Hello, $name!"
func HelloWorld(name string) string {
var greeting string
if len(name) == 0 {
	greeting = "Hello, World!"
} else {
	greeting = fmt.Sprintf("Hello, %s!", name)
}
	return greeting
}

func main() {
	name := os.Args[1]
	fmt.Println(HelloWorld(name))
}
