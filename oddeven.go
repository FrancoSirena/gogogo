package main

import "fmt"

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) print() {
	cur := this
	for cur != nil {
		fmt.Print(cur.Val)
		cur = cur.Next
		if cur != nil {
			fmt.Print("->")
		}
	}
	fmt.Println("")
}

func oddEvenList(head *ListNode) *ListNode {
	cur := head
	idx := 0

	var odd *ListNode = &ListNode{}
	var even *ListNode = &ListNode{}

	cur_odd := odd
	cur_even := even

	for cur != nil {
		if idx%2 == 0 {
			cur_odd.Next = &ListNode{}
			cur_odd = cur_odd.Next
			cur_odd.Val = cur.Val
		} else {
			cur_even.Next = &ListNode{}
			cur_even = cur_even.Next
			cur_even.Val = cur.Val
		}

		idx++
		cur = cur.Next
	}

	cur_even.Next = nil
	cur_odd.Next = nil

	cur_odd.Next = even.Next

	return odd.Next

}

func main() {
	var list []int = []int{1, 2, 3, 4, 5}
	head := &ListNode{}
	root := head
	for id, v := range list {
		fmt.Println(v)
		head.Val = v
		if id < len(list)-1 {
			head.Next = &ListNode{}
			head = head.Next
		}
	}

	root.print()
	res := oddEvenList(root)
	res.print()
}
