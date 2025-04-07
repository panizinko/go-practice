package challenge

func ReverseString(str string) string {
	strSlice := []rune(str)

	for i := 0; i < len(strSlice)/2; i++ {
		tmp := strSlice[i]
		strSlice[i] = strSlice[len(strSlice)-1-i]
		strSlice[len(strSlice)-1-i] = tmp
	}
	return string(strSlice)
}
