package function

import "fmt"

func a() {
	defer end(start("b"))
	fmt.Println("a")
}

func start(t string) string {
	fmt.Println("start: ", t)
	return t
}

func end(t string) {
	fmt.Println("end: ", t)
}

func Defer4() {
	a()
}
