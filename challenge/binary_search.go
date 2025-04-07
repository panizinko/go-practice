package challenge

func BinarySearch(sortedArr []int, seekNumber int) int {
	lowIndex := 0
	highIndex := len(sortedArr) - 1

	for lowIndex <= highIndex {
		middleIdx := (lowIndex + highIndex) / 2

		if seekNumber == sortedArr[middleIdx] {
			return middleIdx
		} else if seekNumber < sortedArr[middleIdx] {
			highIndex = middleIdx - 1
		} else {
			lowIndex = middleIdx + 1
		}
	}

	return -1
}
