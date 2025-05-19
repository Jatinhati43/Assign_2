package line

import (
	"strings"
)

// CountWords returns the number of words in a given line of text.
func CountWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}
