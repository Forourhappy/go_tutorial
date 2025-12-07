package object

import "fmt"

func printValue(s any) { // == interface{}
	fmt.Println(s)
}

func Interface5() {
	/*
		인터페이스 활용(빈 인터페이스)
		함수 내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다.(만능) -> 모든 타입 지정 가능
	*/

	dog := Dog{"poll", 10}
	cat := Cat{"marry", 12}

	printValue(dog)
	printValue(cat)
}
