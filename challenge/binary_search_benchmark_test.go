package challenge

import (
	"testing"
)

var sortedArr = []int{2, 5, 7, 8, 11, 12, 15, 20, 22, 25, 30, 35, 40, 42, 45, 50}
var target = 22

func BenchmarkBinarySearchIterative(b *testing.B) {
	for b.Loop() {
		BinarySearch(sortedArr, target)
	}
}

func BenchmarkBinarySearchRecursive(b *testing.B) {
	for b.Loop() {
		BinarySearchRecursive(sortedArr, target, 0, len(sortedArr)-1)
	}
}
