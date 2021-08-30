package main

import "fmt"

func search(nums []int, target int) int {
	pivot := int(float64(len(nums) / 2))

	start := 0

	fmt.Println(pivot)
	if nums[pivot] < target {
		start = pivot
	}

	var res = -1
	for idx := start; idx < len(nums); idx++ {
		if nums[idx] == target {
			res = idx
			break
		}
	}

	return res
}

func main() {
	fmt.Print(search([]int{2, 5}, 2))
}
