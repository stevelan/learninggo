package ch3

import "fmt"

//Ex3_1SliceCapacity p40
func Ex3_1SliceCapacity() {
	var x []int
	fmt.Println("Slice, len, cap ")
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	// printSlice(x)
}

// func printSlice(x []Type) {
// 	fmt.Println(x, len(x), cap(x))
// }
