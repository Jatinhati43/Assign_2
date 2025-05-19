package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Jatinhati43/Assign_2/internal/line"
)

func main() {
	fmt.Print("Enter a line of text: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()                       
		count := line.CountWords(input)             
		fmt.Printf("Word count: %d\n", count)
	}
}
