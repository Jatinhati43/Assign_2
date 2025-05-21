// Project: Word Count 
// Description: This program reads a line of text from
//              the user input and counts the number of words
//              using a reusable function from the internal module.


package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Jatinhati43/Assign_2/internal/line"
)


// main is the entry point of the program
// It prompts the user to enter a line of text,
// reads the input, and prints the number of words
// using the CountWords function from the line package.
func main() {
	fmt.Print("Enter a line of text: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		count := line.CountWords(input)
		fmt.Printf("Word count: %d\n", count)
	}
}
