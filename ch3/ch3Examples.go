package ch3

import (
	"fmt"
)

// Ex3_1SliceCapacity p40
func Ex3_1SliceCapacity() {
	var x []int
	fmt.Println("Slice, len, cap ")
	printSlice(x)

	for i := 0; i < 10; i++ {
		x = append(x, 10*(i+1))
		printSlice(x)
	}

}

func Ex3_2NilSlice() {
	var nilData []int
	printSlice(nilData)
	if nilData == nil {
		fmt.Println("Nil slice ", nilData)
	}
}

func Ex3_3DefaultValuedSlice() {
	data := []int{2, 4, 6, 8}
	printSlice(data)
}

func Ex3_4SlicingSlices() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]

	print("x:")
	printSlice(x)
	print("y:")
	printSlice(y)
	print("z:")
	printSlice(z)
	print("d:")
	printSlice(d)
	print("e:")
	printSlice(e)
}

// CommonMistakeMakeAppend p41
func CommonMistakeMakeAppend() {
	x := make([]int, 5)
	x = append(x, 10)

	printSlice(x)
}

func printSlice(x []int) {
	fmt.Println(x, len(x), cap(x))
}
