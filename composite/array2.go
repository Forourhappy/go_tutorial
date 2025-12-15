package composite

import "fmt"

func Array2() {
	// 배열 순회

	arr1 := [5]int{1, 10, 100, 1000, 10000}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	fmt.Println()

	arr2 := [5]int{1, 11, 111, 1111, 11111}

	// 인덱스 생략1
	for _, v := range arr2 {
		fmt.Println(v)
	}
	fmt.Println()

	// 값 생략
	for i := range arr2 {
		fmt.Println(i)
	}

}
