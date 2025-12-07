package object

import "fmt"

func checkType(arg any) {
	switch arg.(type) {
	case bool:
		fmt.Println("bool")
	case int, int8, int16, int32, int64:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("unknown")
	}
}

func InterfaceEx3() {
	/*
		실제 타입 검사 switch
		빈 인터페이스는 어떠한 자료형도 전달 받을 수 있으므로, 타입 체크를 통해 형 변환 후 사용 가능
	*/

	checkType(true)
	checkType(1)
	checkType(123.23)
	checkType(nil)
	checkType("Hello Golang")
}
