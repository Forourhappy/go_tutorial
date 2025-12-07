package object

import (
	"fmt"
	"reflect"
)

func InterfaceEx2() {
	/*
		타입 변환
		type Assertion
		실행(런타임)시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
		인터페이스.(타입) -> 형 변환
		interfaceVal.(type)
	*/

	var a any = 15
	b := a
	c := a.(int)
	// d := a.(float64)

	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))
	// fmt.Println(d)

	fmt.Println()

	if v, ok := a.(int); ok {
		fmt.Println(v, ok)
	}
}
