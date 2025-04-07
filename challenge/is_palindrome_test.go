package challenge

import "testing"

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty", "", true},
		{"Single", "a", true},
		{"Simple", "madam", true},
		{"WithSpaces", "race car", true},
		{"MixedCase", "Racecar", true},
		{"NoPalindrome", "hello", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsPalindrome(tc.input)
			if actual != tc.expected {
				t.Errorf("IsPalindrome(%q): expected %v, actual %v", tc.input, tc.expected, actual)
			}
		})
	}
}
