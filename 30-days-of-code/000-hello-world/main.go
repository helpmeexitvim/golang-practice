package main

import (
	"bufio"
	"fmt"
	"os"
)

// reads a string from StdIn and prints out "Hello World." on the first line, and the input string on the second line
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Printf("Hello, World.\n%s", text)
}
