package composite

import "fmt"

func Pointer1() {
	/*
		포인터
		Go: 포인터 지원(c와 동일)
		변수의 지역성, 연속된 메모리 참조 ..., 힙, 스택...
		파이썬, 자바(JRE) -> 인터프리터, 컴파일러
		포인터 지원 X(파이썬, C#, Java 등)
		주소의 값은 직접 변경 불가능(잘못된 코딩으로 인한 버그 방지)
		*(애스터리스크)로 사용
		nil로 초기화 (nil == 0)
	*/

	var a *int
	var b *int = new(int)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println()

	i := 7

	a = &i
	b = &i

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*a)
	fmt.Println()
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(*b)
	fmt.Println()

	var c = &i
	d := &i

	*d = 10

	fmt.Println(c)
	fmt.Println(&c)
	fmt.Println(*c)
	fmt.Println()
	fmt.Println(d)
	fmt.Println(&d)
	fmt.Println(*d)
}
