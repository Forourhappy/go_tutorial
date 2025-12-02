package condition

import "fmt"

func If1() {
	/**
	제어문(조건문)
	if 문: 반드시 Boolean 검사 -> 1, 0 (사용 불가 : 자동 형 변환 불가)
	소괄호 사용 X
	*/

	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("15 이상")
	}

	if b >= 25 {
		fmt.Println("25 이상")
	}

	/**
	에러 발생
	컴파일러에서 자동으로 문장의 끝에 세미콜론을 붙이기 때문에 컴파일 에러
	if b >= 35
	{
	fmt.Println("35 이상")
	}
	*/
}
