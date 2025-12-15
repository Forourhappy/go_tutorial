package datatype

import "fmt"

func Numeric1() {
	/*
		데이터 타입: 숫자형
		정수, 실수, 복소수
		32비트, 64비트, unsigned(양수)
		정수: 8진수(0), 16진수(0x), 10진수
	*/

	var num1 int = 17
	var num2 int = -68
	var num3 int = 0631
	var num4 int = 0x32fa2c75
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
}
