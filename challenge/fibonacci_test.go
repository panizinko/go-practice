package challenge

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected []int
	}{
		{"Zero", 0, []int{}},
		{"One", 1, []int{0}},
		{"Two", 2, []int{0, 1}},
		{"Three", 3, []int{0, 1, 1}},
		{"Ten", 10, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Fibonacci(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Fibonacci(%d): expected %v, actual %v", tc.input, tc.expected, actual)
			}
		})
	}
}
