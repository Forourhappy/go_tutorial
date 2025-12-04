package function

import "fmt"

func sum1(i int, f func(int, int)) {
	f(i, 10)
}

func add(a int, b int) {
	fmt.Println(a + b)
}

func multiValue(i int) {
	i = i * 10
}

func multiReference(i *int) {
	*i *= 10
}

func Function2() {
	/*
		함수(콜백) 전달, 참조 전달(call by reference), 값 전달(call by value)
	*/

	sum1(10, add)
	a := 100
	multiValue(a)
	fmt.Println(a)

	b := 100
	multiReference(&b)
	fmt.Println(b)
}
