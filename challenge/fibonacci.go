package challenge

func Fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	if n == 1 {
		return []int{0}
	}

	if n == 2 {
		return []int{0, 1}
	}

	fibonacciSlice := make([]int, n)
	fibonacciSlice[0] = 0
	fibonacciSlice[1] = 1

	for i := 2; i < n; i++ {
		fibonacciSlice[i] = fibonacciSlice[i-1] + fibonacciSlice[i-2]
	}

	return fibonacciSlice
}
