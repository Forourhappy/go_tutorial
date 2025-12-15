package datatype

import "fmt"

func Numeric2() {
	/*
		데티터 타입: 숫자형
		정수형 문자 출력
	*/

	var char1 byte = 72
	var char2 byte = 0110
	var char3 byte = 0x48

	var char4 rune = 50556
	var char5 rune = 0142574
	var char6 rune = 0xC57C

	fmt.Printf("%c %c %c \n", char1, char2, char3)
	fmt.Printf("%d %d %d \n", char1, char2, char3)
	fmt.Printf("%d %o %x \n", char1, char2, char3)
	fmt.Printf("%c %c %c \n", char4, char5, char6)
	fmt.Printf("%d %d %d \n", char4, char5, char6)
	fmt.Printf("%d %o %x \n", char4, char5, char6)

}
