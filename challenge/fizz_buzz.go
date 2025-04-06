package challenge

import "fmt"

func FizzBuzz(n int) {
	rules := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	for i := 1; i <= n; i++ {
		result := ""

		for division, text := range rules {
			if i%division == 0 {
				result += text
			}
		}

		if result == "" {
			fmt.Println(i)
		} else {
			fmt.Println(result)
		}
	}
}
