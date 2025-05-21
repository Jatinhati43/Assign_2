package line

import "testing"

// TestCountWords tests the CountWords function with various input strings.
// It verifies that the function returns the correct number of words.
func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Empty", "", 0},
		{"Single", "Hello", 1},
		{"Multiple", "Hello World", 2},
		{"Whitespace", "  Hello   Go  ", 2},
		{"Newlines", "Go\nLang", 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountWords(tc.input); got != tc.expected {
				t.Errorf("CountWords(%q) = %d; want %d", tc.input, got, tc.expected)
			}
		})
	}
}
