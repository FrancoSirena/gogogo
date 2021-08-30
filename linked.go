package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"time"
)

type Input struct {
	Calls    []string      `json:"calls"`
	Args     [][]int       `json:"args"`
	Expected []interface{} `json:"expected"`
}

type Node struct {
	val  int
	prev *Node
	next *Node
}

type MyLinkedList struct {
	root *Node
	tail *Node
	len  int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{nil, nil, 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.len {
		fmt.Println("GET", -1)
		return -1
	}

	if index == this.len-1 {
		fmt.Println("GET", this.tail.val)
		return this.tail.val
	}

	cur := this.root
	idx := 0
	for cur != nil {
		if idx == index {
			fmt.Println("GET", cur.val)
			return cur.val
		}
		idx++
		cur = cur.next
	}
	fmt.Println("GET", -1)
	return -1
}

func (this *MyLinkedList) _AddRoot(val int) {
	this.root = &Node{val, nil, nil}
	this.tail = this.root
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.len++
	if this.root == nil {
		this._AddRoot(val)
		this.print()
		return
	}

	nroot := &Node{val, nil, this.root}
	this.root.prev = nroot
	this.root = nroot
	this.print()
	return
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.len++
	if this.root == nil {
		this._AddRoot(val)
		this.print()
		return
	}

	if this.tail == this.root {
		this.tail = &Node{val, this.root, nil}
		this.root.next = this.tail
		this.print()
		return
	}

	ntail := &Node{val, this.tail, nil}
	this.tail.next = ntail
	this.tail = ntail
	this.print()
	return
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.len {
		return
	}

	if this.root == nil {
		this.len++
		this._AddRoot(val)
		this.print()
		return
	}

	if index == 0 {
		this.len++
		nroot := &Node{val, nil, this.root}
		this.root.prev = nroot
		this.root = nroot
		this.print()
		return
	}

	if index == this.len {
		this.len++
		if this.tail.prev == nil {
			ntail := &Node{val, this.root, nil}
			this.root.next = ntail
			this.tail = ntail
			this.print()
			return
		}
		ntail := &Node{val, this.tail, nil}
		this.tail.next = ntail
		this.tail = ntail
		this.print()
		return
	}

	this.len++
	cur := this.root
	idx := 0
	for cur != nil {
		if idx == index {
			ncur := &Node{val, cur.prev, cur}
			cur.prev.next = ncur
			cur.prev = ncur
			this.print()
			return
		}
		idx++
		cur = cur.next
	}
	return
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.len {
		return
	}

	if index == 0 && this.len-1 == 0 {
		this.len--
		this.root = nil
		this.tail = nil
		return
	}

	if index == this.len-1 {
		this.len--
		this.tail = this.tail.prev
		if this.tail != nil {
			this.tail.next = nil
		}
		this.print()
		return
	}

	if index == 0 {
		this.len--
		this.root = this.root.next
		if this.root != nil {
			this.root.prev = nil
		}
		this.print()
		return
	}
	this.len--
	cur := this.root
	idx := 0
	for cur != nil {
		if idx == index {
			cur.prev.next = cur.next
			cur.next.prev = cur.prev
			this.print()
			return
		}
		idx++
		cur = cur.next
	}
	return
}

func (this *MyLinkedList) print() {
	fmt.Println("Onwards", this.len)
	cur := this.root
	for cur != nil {
		fmt.Print(cur.val)
		cur = cur.next
		if cur != nil {
			fmt.Print("->")
		}
	}
	fmt.Println("")

	fmt.Println("Backwards")
	cur = this.tail
	for cur != nil {
		fmt.Print(cur.val)
		cur = cur.prev
		if cur != nil {
			fmt.Print("<-")
		}
	}
	fmt.Println("")
}

func main() {
	// myLinkedList := Constructor()
	// myLinkedList.AddAtHead(1)
	// myLinkedList.AddAtTail(3)
	// myLinkedList.AddAtIndex(1, 2)    // linked list becomes 1->2->3
	// fmt.Println(myLinkedList.Get(1)) // return 2
	// myLinkedList.DeleteAtIndex(1)    // now the linked list is 1->3
	// fmt.Println(myLinkedList.Get(1)) // return 3

	// myLinkedList1 := Constructor()
	// myLinkedList1.AddAtHead(7)
	// myLinkedList1.AddAtHead(2)
	// myLinkedList1.AddAtHead(1)
	// myLinkedList1.AddAtIndex(3, 0)
	// myLinkedList1.DeleteAtIndex(2)
	// myLinkedList1.AddAtHead(6)
	// myLinkedList1.AddAtTail(4)
	// fmt.Println(myLinkedList1.Get(4))
	// myLinkedList1.AddAtHead(4)
	// myLinkedList1.AddAtIndex(5, 0)
	// myLinkedList1.AddAtHead(6)

	// myHeadless := Constructor()
	// myHeadless.AddAtIndex(0, 10)
	// myHeadless.AddAtIndex(0, 20)
	// myHeadless.AddAtIndex(1, 30)

	// myCulo := Constructor()
	// myCulo.AddAtHead(0)
	// myCulo.AddAtIndex(1, 4)
	// myCulo.AddAtTail(8)

	// myCulo1 := Constructor()
	// myCulo1.AddAtHead(3)
	// myCulo1.AddAtHead(4)
	// myCulo1.AddAtTail(5)
	// myCulo1.AddAtTail(9)
	// myCulo1.AddAtTail(7)
	// myCulo1.AddAtIndex(1, 1)

	// instr := []string{"addAtHead", "addAtTail", "addAtTail", "addAtTail", "addAtTail", "addAtTail", "addAtTail", "deleteAtIndex", "addAtHead", "addAtHead", "get", "addAtTail", "addAtHead", "get", "addAtTail", "addAtIndex", "addAtTail", "addAtHead", "addAtHead", "addAtHead", "get", "addAtIndex", "addAtHead", "get", "addAtHead", "deleteAtIndex", "addAtHead", "addAtTail", "addAtTail", "addAtIndex", "addAtTail", "addAtHead", "get", "addAtTail", "deleteAtIndex", "addAtIndex", "deleteAtIndex", "addAtHead", "addAtTail", "addAtHead", "addAtHead", "addAtTail", "addAtTail", "get", "get", "addAtHead", "addAtTail", "addAtTail", "addAtTail", "addAtIndex", "get", "addAtHead", "addAtIndex", "addAtHead", "addAtTail", "addAtTail", "addAtIndex", "deleteAtIndex", "addAtIndex", "addAtHead", "addAtHead", "deleteAtIndex", "addAtTail", "deleteAtIndex", "addAtIndex", "addAtTail", "addAtHead", "get", "addAtIndex", "addAtTail", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "deleteAtIndex", "get", "get", "addAtHead", "get", "addAtTail", "addAtTail", "addAtIndex", "addAtIndex", "addAtHead", "addAtTail", "addAtTail", "get", "addAtIndex", "addAtHead", "deleteAtIndex", "addAtTail", "get", "addAtHead", "get", "addAtHead", "deleteAtIndex", "get", "addAtTail", "addAtTail"}

	// args := [][]int{{38}, {66}, {61}, {76}, {26}, {37}, {8}, {5}, {4}, {45}, {4}, {85}, {37}, {5}, {93}, {10, 23}, {21}, {52}, {15}, {47}, {12}, {6, 24}, {64}, {4}, {31}, {6}, {40}, {17}, {15}, {19, 2}, {11}, {86}, {17}, {55}, {15}, {14, 95}, {22}, {66}, {95}, {8}, {47}, {23}, {39}, {30}, {27}, {0}, {99}, {45}, {4}, {9, 11}, {6}, {81}, {18, 32}, {20}, {13}, {42}, {37, 91}, {36}, {10, 37}, {96}, {57}, {20}, {89}, {18}, {41, 5}, {23}, {75}, {7}, {25, 51}, {48}, {46}, {29}, {85}, {82}, {6}, {38}, {14}, {1}, {12}, {42}, {42}, {83}, {13}, {14, 20}, {17, 34}, {36}, {58}, {2}, {38}, {33, 59}, {37}, {15}, {64}, {56}, {0}, {40}, {92}, {63}, {35}, {62}, {32}}
	fmt.Println(time.Now())
	jsonFile, _ := os.Open("input_f.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result Input
	json.Unmarshal([]byte(byteValue), &result)

	var output []interface{}

	var valueOf = map[string]reflect.Value{
		"MyLinkedList": reflect.ValueOf(&MyLinkedList{}),
	}
	for idx, call := range result.Calls {
		norm_method := strings.ToUpper(call[0:1]) + call[1:]
		meth := valueOf["MyLinkedList"].MethodByName(norm_method)
		in := make([]reflect.Value, len(result.Args[idx]))
		for k, param := range result.Args[idx] {
			in[k] = reflect.ValueOf(param)
		}
		fmt.Println("")
		fmt.Println(norm_method, "start", result.Args[idx])
		res := meth.Call(in)
		var converted interface{}
		if len(res) == 0 {
			converted = nil
		} else {
			converted = res[0].Int()
		}
		output = append(output, converted)
		fmt.Println(norm_method, "finish")
		fmt.Println("")
	}

	fmt.Println(result.Expected...)
	fmt.Println("CURR")
	fmt.Println(output...)

	fmt.Println(time.Now())

}

// ["MyLinkedList","addAtIndex","addAtIndex","addAtIndex","get"]
// [[],[0,10],[0,20],[1,30],[0]]

// ["addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
// [[7],[2],[1],[3,0],[2],[6],[4],[4],[4],[5,0],[6]]
