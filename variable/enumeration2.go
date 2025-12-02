package variable

import "fmt"

func Enumeration2() {
	const (
		A = iota * 10
		B
		C
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
