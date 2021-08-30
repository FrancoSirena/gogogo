package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res1 := strconv.Itoa(l1.Val)
	cur := l1.Next
	for cur != nil {
		res1 = strconv.Itoa(cur.Val) + res1
		cur = cur.Next
	}

	res2 := strconv.Itoa(l2.Val)
	cur = l2.Next
	for cur != nil {
		res2 = strconv.Itoa(cur.Val) + res2
		cur = cur.Next
	}

	fmt.Println(res1, res2)
	c1, _ := new(big.Int).SetString(res1, 0)
	c2, _ := new(big.Int).SetString(res2, 0)
	fmt.Println(c1, c2)
	res := c1.Add(c1, c2)
	fmt.Println(res)
	rr := strings.Split(res.String(), "")
	fmt.Println(rr)
	c, _ := strconv.Atoi(rr[len(rr)-1])
	result := &ListNode{c, nil}
	root := result
	for idx := len(rr) - 2; idx >= 0; idx-- {
		c, _ = strconv.Atoi(rr[idx])
		result.Next = &ListNode{c, nil}
		result = result.Next
	}
	return root
}

func main() {
	arr1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	arr2 := []int{5, 6, 4}
	link1 := &ListNode{}
	link2 := &ListNode{}
	root1 := link1
	root2 := link2
	for id, v := range arr1 {
		link1.Val = v
		if id < len(arr1)-1 {
			link1.Next = &ListNode{}
			link1 = link1.Next
		}
	}
	for id, v := range arr2 {
		link2.Val = v
		if id < len(arr2)-1 {
			link2.Next = &ListNode{}
			link2 = link2.Next
		}
	}
	addTwoNumbers(root1, root2)
}
