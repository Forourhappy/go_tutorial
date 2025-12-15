package function

import "fmt"

func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n, "Hello")
	prtHello(n - 1)
}

func FunctionEx3() {
	/*
		재귀 함수
		프로그램이 보기 쉽고, 코드 간결, 오류 수정이 용이: 장점
		코드 이해하기 어렵고, 메모리를 많이 사용, 무한 루프 가능함
	*/

	x := fact(5)
	fmt.Println(x)

	prtHello(5)
}
