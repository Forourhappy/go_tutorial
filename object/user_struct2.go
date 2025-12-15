package object

import "fmt"

type cnt int

func UserStruct2() {
	// 기본 자료형 사용자 정의 타입

	a := cnt(5)
	fmt.Println(a)

	var b cnt = 15
	fmt.Println(b)

	// testConvertT(b) // 예외 발생!
	testConvertT(int(b))
	testConvertD(b)
}

func testConvertT(i int) {
	fmt.Println("(Default Type): ", i)
}

func testConvertD(i cnt) {
	fmt.Println("(Custom Type): ", i)
}
