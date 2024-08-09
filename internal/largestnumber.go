package internal

import (
	"fmt"
	"sort"
	"strconv"
)

func FindLargestNumber() {
	input := []int{56, 564, 568, 60}

	fmt.Println(LargestNumber(input))
}

func LargestNumber(input []int) string {
	strs := make([]string, len(input))
	for i, num := range input {
		strs[i] = strconv.Itoa(num)
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	var largestNumber string
	for _, str := range strs {
		largestNumber += str
	}

	return largestNumber
}
