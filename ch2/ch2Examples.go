package ch2

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Ex2_1ComplexNumbers p25
func Ex2_1ComplexNumbers() {
	x := complex(1, math.Sqrt(3))
	y := complex(10.2, 2)

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))
}

// Ex2_2NumericTypeConversion type conversion, p27
func Ex2_2NumericTypeConversion() {
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)
	fmt.Println(x, y, z, d)
}
