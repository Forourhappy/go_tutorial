package function

import "fmt"

func FunctionEx4() {
	/*
		익명 함수
		즉시 실행(선언과 동시에)
	*/

	func() {
		fmt.Println("Hello")
	}()

	msg := "Hello Golang!"
	func(m string) {
		fmt.Println(m)
	}(msg)

	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	r := func(x, y int) int {
		return x * y
	}(10, 100)
	fmt.Println(r)
}
