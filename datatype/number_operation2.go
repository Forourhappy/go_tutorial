package datatype

import "fmt"

func NumberOperation2() {
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 * n2)
	fmt.Println(n1 / n2)
	fmt.Println(n1 % n2)
	fmt.Println(n1 << n2)
	fmt.Println(n1 >> n2)
	fmt.Println(n1 & n2)
	fmt.Println(n1 | n2)
	fmt.Println(^n1)

	var n3 int = 12
	var n4 float32 = 8.2

	fmt.Println(float32(n3) + n4)
	fmt.Println(n3 + int(n4))
}
