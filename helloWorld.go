package main // Has to match the declaration of "go mod init {dir}/{declaration}"

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// func main() {
// 	fmt.Println("Hello, World!")
// }
