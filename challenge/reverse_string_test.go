package challenge

import "testing"

func TestReverseString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty", "", ""},
		{"Single", "a", "a"},
		{"Hello", "hello", "olleh"},
		{"Palindrome", "kajak", "kajak"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ReverseString(tc.input)
			if actual != tc.expected {
				t.Errorf("ReverseString(%q): expected %q, actual %q", tc.input, tc.expected, actual)
			}
		})
	}
}
