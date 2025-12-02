package variable

import "fmt"

func getHeight() int {
	return 10
}

func Const1() {
	/**
	상수
	const 사용 초기화, 한번 선언 후 값 변경 금지, 고정된 값 관리용
	*/

	const a string = "TEST1"
	const b = "TEST2"
	const c int32 = 10 * 10
	// const d = getHeight() // 함수의 리턴 값으로 상수 선언 불가
	const e = 35.6
	const f = false
	/**
	에러 발생
	const g string
	g = "TEST#"
	*/

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(e)
	fmt.Println(f)
}
