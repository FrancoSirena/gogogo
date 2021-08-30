package main

import (
	"fmt"
)

// func main1() {
// 	x := 42
// 	y := "James Bond"
// 	z := true

// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// 	fmt.Println(x, y, z)
// }

// var x int = 42
// var y string = "James Bond"
// var z bool = false

// func main2() {
// 	fmt.Printf("%v\n%v\n%v\n", x, y, z)
// }

// func main3() {
// 	s := fmt.Sprintf("%v\n%v\n%v\n", x, y, z)
// 	fmt.Println(s)
// }

// func main4() {
// 	fmt.Println(x)
// 	fmt.Printf("%T\n", x)
// 	x = 42
// 	fmt.Println(x)
// }

type tt int

var x tt
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Printf("%T\n", y)
}
