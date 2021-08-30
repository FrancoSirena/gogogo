package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str_x := strconv.Itoa(x)
	if len(str_x) == 1 {
		return true
	}
	pivot := int(len(str_x) / 2)
	right := len(str_x) - 1
	left := 0
	for left <= pivot {
		if str_x[left] != str_x[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome((-120)))
	fmt.Println(isPalindrome((121)))
	fmt.Println(isPalindrome((1342431)))
}
