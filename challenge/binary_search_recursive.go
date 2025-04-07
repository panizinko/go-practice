package challenge

func BinarySearchRecursive(sortedArr []int, seekNumber int, lowIndex int, highIndex int) int {
	if lowIndex > highIndex {
		return -1
	}

	middleIdx := (lowIndex + highIndex) / 2

	if seekNumber == sortedArr[middleIdx] {
		return middleIdx
	} else if seekNumber < sortedArr[middleIdx] {
		return BinarySearchRecursive(sortedArr, seekNumber, lowIndex, middleIdx-1)
	} else {
		return BinarySearchRecursive(sortedArr, seekNumber, middleIdx+1, highIndex)
	}
}
