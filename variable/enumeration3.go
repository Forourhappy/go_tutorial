package variable

import "fmt"

func Enumeration3() {
	const (
		_ = iota
		A
		B
		C
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		PLATINUM
	)

	fmt.Println("DEFAULT : ", DEFAULT)
	fmt.Println("SILVER : ", SILVER)
	fmt.Println("GOLD : ", GOLD)
	fmt.Println("PLATINUM : ", PLATINUM)
}
