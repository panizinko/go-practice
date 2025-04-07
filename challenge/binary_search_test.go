package challenge

import "testing"

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		{"Empty", []int{}, 5, -1},
		{"SingleFound", []int{5}, 5, 0},
		{"SingleNotFound", []int{5}, 10, -1},
		{"First", []int{2, 5, 7, 8, 11, 12}, 2, 0},
		{"Last", []int{2, 5, 7, 8, 11, 12}, 12, 5},
		{"Middle", []int{2, 5, 7, 8, 11, 12}, 8, 3},
		{"NotFoundLow", []int{2, 5, 7, 8, 11, 12}, 0, -1},
		{"NotFoundHigh", []int{2, 5, 7, 8, 11, 12}, 13, -1},
		{"NotFoundMiddle", []int{2, 5, 7, 8, 11, 12}, 9, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := BinarySearch(tc.input, tc.target)
			if actual != tc.expected {
				t.Errorf("BinarySearch(%v, %d): expected %d, actual %d", tc.input, tc.target, tc.expected, actual)
			}
		})
	}
}
