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
		line := scanner.Text()
		count := number.CountWords(line) // use correct package and function name
		fmt.Printf("Word count: %d\n", count)
	}
}
