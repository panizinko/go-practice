package challenge

import "testing"

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Two", 2, true},
		{"Three", 3, true},
		{"Four", 4, false},
		{"Five", 5, true},
		{"Six", 6, false},
		{"Seven", 7, true},
		{"Eight", 8, false},
		{"Nine", 9, false},
		{"Ten", 10, false},
		{"Eleven", 11, true},
		{"OneHundred", 100, false},
		{"OneHundredOne", 101, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsPrime(tc.input)
			if actual != tc.expected {
				t.Errorf("IsPrime(%d): expected %v, actual %v", tc.input, tc.expected, actual)
			}
		})
	}
}
