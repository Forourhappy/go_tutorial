package variable

import "fmt"

func Variable3() {
	/**
	짧은 선언
	함수 안에서만 사용(전역 X), 선언 후 할당 예외 발생
	주로 제한된 범위의 함수 내에서 사용할 경우 코드 가독성을 넓일 수 있음 -> 추천
	*/

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := true

	// shortVar3 := false // 예외 발생

	fmt.Println(shortVar1)
	fmt.Println(shortVar2)
	fmt.Println(shortVar3)

}
