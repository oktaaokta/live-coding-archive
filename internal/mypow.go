package internal

import "fmt"

func FindPow() {
	fmt.Println(Mypow(2, 10))
}

func Mypow(base, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	}

	result := 1
	currentBase := base

	for power > 0 {
		if power%2 == 1 {
			result *= currentBase
		}
		currentBase *= currentBase
		power /= 2
	}

	return result
}
