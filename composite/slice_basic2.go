package composite

import "fmt"

func Slice2() {
	// 슬라이스 참조 타입 증명

	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1
	arr2[0] = 7

	fmt.Println(arr1)
	fmt.Println(arr2)

	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1
	slice2[0] = 7

	fmt.Println(slice1)
	fmt.Println(slice2)

	// 슬라이스 예외 상황
	slice3 := make([]int, 50, 100)
	fmt.Println(slice3[4])
	// fmt.Println(slice3[50]) //길이 index over 예외 (길이만큼 초기화)

	for i, v := range slice1 {
		fmt.Println(i, v)
	}
}
