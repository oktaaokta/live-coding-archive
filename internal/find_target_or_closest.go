package internal

import (
	"fmt"
)

func FindTarget() {
	input := []int{2, 4, 5, 7, 9}
	target := 6

	fmt.Println(findTargetOrClosest(input, target))
}

// finds target or closest to target (smallest before target in index)
// If it goes out of bounds, return -1.
// Else return target or smaller than target existing in array.
func findTargetOrClosest(input []int, target int) int {
	if len(input) == 0 {
		return -1
	}

	i, j := 0, len(input)-1
	for i < j {
		mid := (i + j) / 2
		if input[mid] == target {
			return target
		}

		if input[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}

	}

	if i > 0 {
		return input[i-1]
	}

	return -1
}
