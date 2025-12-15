package function

import "fmt"

func increaseCnt() func() int {
	n := 0 // 지역변수(캡쳐됨)

	return func() int {
		n += 1
		return n
	}
}

func Closure2() {
	cnt := increaseCnt()
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
}
