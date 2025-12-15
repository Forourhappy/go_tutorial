package datatype

import "fmt"

func Numeric3() {
	/*
		실수(부동소수점)
		float32(7자리), float64(15자리)
	*/

	var num1 float32 = 0.14
	var num2 float64 = .75647
	var num3 float32 = 442.0378373
	var num4 float64 = 10.0

	var num5 float32 = 14e6
	var num6 float64 = .156875e+3
	var num7 float64 = 5.32521e-10

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)

	fmt.Println(num4)
	fmt.Printf("%.20f\n", num4-0.1)
	fmt.Printf("%.20f\n", float32(num4-0.1))
	fmt.Printf("%.20f\n", float64(num4-0.1))

	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(num7)
}
