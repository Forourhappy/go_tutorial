package function

import "fmt"

func multiply2(n ...int) int {
	tot := 1

	for _, val := range n {
		tot *= val
	}

	return tot
}

func sum2(n ...int) int {
	tot := 0

	for _, val := range n {
		tot += val
	}

	return tot
}

func prtWord(msg ...string) {
	for _, val := range msg {
		fmt.Println(val)
	}
}

func FunctionEx1() {
	/*
		함수 고급
		가변 인자 실습(매개 변수 개수가 동적으로 변할 때 - 정해져 있지 않음)
	*/

	x := multiply2(1, 2, 3, 4, 5)
	fmt.Println(x)

	y := sum2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(y)

	prtWord("Hello", "Golang", "World")

	a := []int{5, 6, 7, 8, 9, 10}

	m := multiply2(a...)
	fmt.Println(m)
}
