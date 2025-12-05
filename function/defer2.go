package function

import "fmt"

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("hi")
	}()
}

func Defer2() {
	sayHello("Hello")
}
