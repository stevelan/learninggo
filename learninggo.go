package main

import (
	"fmt"
	"reflect"
	"runtime"

	"src.stevelan.au/learninggo/v2/ch1"
	"src.stevelan.au/learninggo/v2/ch2"
	"src.stevelan.au/learninggo/v2/ch3"
)

type example func()

func main() {
	call(ch1.Hello)
	call(ch2.Ex2_1ComplexNumbers)
	call(ch2.Ex2_2NumericTypeConversion)
	call(ch2.Ex2_3Const)
	call(ch3.Ex3_1SliceCapacity)
	call(ch3.CommonMistakeMakeAppend)
	call(ch3.Ex3_2NilSlice)
	call(ch3.Ex3_3DefaultValuedSlice)
	call(ch3.Ex3_4SlicingSlices)
}

func call(ex example) {
	fmt.Println("=====================================\n** " +
		runtime.FuncForPC(reflect.ValueOf(ex).Pointer()).Name() + "\n")
	ex()
	fmt.Println()
}
