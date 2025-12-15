package function

import "fmt"

func multiply3(x, y int) int {
	return x * y
}

func FunctionEx2() {
	/*
		함수 고급
		합수를 변수에 할당
	*/

	// 슬라이스에 할당
	f := []func(int, int) int{multiply3, sum}

	a := f[0](10, 20)
	b := f[1](10, 20)

	fmt.Println(a, b)

	// 변수에 할당

	var f1 func(int, int) int = multiply3
	f2 := sum

	fmt.Println(f1(10, 20), f2(10, 20))

	// map에 할당
	m := map[string]func(int, int) int{
		"mul_func": multiply3,
		"sum_func": sum,
	}

	fmt.Println(m["mul_func"](10, 10))
	fmt.Println(m["sum_func"](10, 10))
}
