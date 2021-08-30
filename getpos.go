package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	if target == nums[len(nums)-1] {
		return len(nums) - 1
	}
	pivot := int(len(nums) / 2)
	var res int
	if nums[pivot] >= target {
		for idx := pivot; idx >= 0; idx-- {
			if nums[idx] < target {
				res = idx + 1
				break
			}
		}
	} else {
		for idx := pivot; idx < len(nums); idx++ {
			if nums[idx] >= target {
				res = idx
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Println(searchInsert([]int{1, 3, 5, 8, 9, 10, 20}, 11))
	fmt.Println(searchInsert([]int{1, 3}, 2))
	fmt.Println(searchInsert([]int{2, 7, 8, 9, 10}, 9))
}
