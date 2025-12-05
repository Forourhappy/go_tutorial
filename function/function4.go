package function

import "fmt"

func multiply1(x, y int) (r1, r2 int) {
	r1 = x * 10
	r2 = y * 10
	return r1, r2
}

func Function4() {
	// 리턴 값 변수 사용
	a, b := multiply1(10, 5)
	fmt.Println(a, b)
}
