package ch2

import "fmt"

const x int64 = 10
const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

// Ex2_3Const p30
func Ex2_3Const() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(idKey)
}
