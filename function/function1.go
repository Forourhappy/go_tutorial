package function

import (
	"fmt"
	"strconv"
)

func helloGolang() {
	fmt.Println("Hello Golang")
}

func sayOne(m string) {
	fmt.Println(m)
}

func sum(x int, y int) int {
	return x + y
}

func Function1() {
	/*
		   함수
		   선언: func 키워드로 선언
		   func 함수명(매개변수) (반환타입 or 반환 값 변수명): 반환 값 존재
		   func 함수명() (반환타입 or 반환 값 변수명): 매개변수 없음, 반환 값 존재
		   func 함수명(매개변수): 매개변수 없음, 반환 값 없음
			 타 언어와 달리 반환 값(return value) 다수 가능
	*/

	helloGolang()
	sayOne("Hello!")

	fmt.Println(sum(1, 2))
	// int -> string으로 변환
	fmt.Println(strconv.Itoa(sum(1, 2)))
}
