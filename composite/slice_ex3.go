package composite

import "fmt"

func SliceEx3() {
	/*
		슬라이스 복사
		copy(복사 대상, 원본)
		make로 공간을 할당 후 복사해야 한다.
		복사된 슬라이스 값 변경해도 원본에는 영향 없음
	*/

	// 복사
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println(slice2)
	fmt.Println(slice3)

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)
	copy(b, a)

	b[0] = 7
	b[4] = 10
	fmt.Println(a)
	fmt.Println(b)

	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] // 주의! 부분적으로 슬라이스 추출은 참조 -> 원본 값 변경

	d[1] = 7
	fmt.Println(c)
	fmt.Println(d)

	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] // 용량 지정

	fmt.Println(len(f), cap(f))
	fmt.Println(f)
}
