package condition

import "fmt"

func For1() {
	/**
	반복문 - for
	Go에서 유일하게 제공되는 반목문
	다양한 사용법 숙지
	*/

	for i := 0; i < 5; i++ {
		fmt.Println("ex1: ", i)
	}

	for {
		fmt.Println("ex2: Hello, Golang")
		fmt.Println("ex2: Infinite Loop!")
		break
	}

	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex3: ", index, name)
	}
}
