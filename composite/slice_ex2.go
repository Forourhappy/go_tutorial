package composite

import (
	"fmt"
	"sort"
)

func SliceEx2() {
	/*
		슬라이스 추출 및 정렬
		slice[i:j] i -> j-1 까지 추출
		slice[i:] i -> 마지막까지 추출
		slice[:j] 처음 -> j-1 까지 추출
		slice[:] 처음 -> 마지막까지 추출
	*/

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice1[2:5])
	fmt.Println(slice1[2:])
	fmt.Println(slice1[:5])
	fmt.Println(slice1[:])

	// 정렬
	slice2 := []int{3, 6, 10, 9, 1, 5, 3, 6, 2, 7}
	slice3 := []string{"b", "d", "e", "a", "f", "c"}

	fmt.Println(sort.IntsAreSorted(slice2))
	sort.Ints(slice2)
	fmt.Println(slice2)

	fmt.Println()

	fmt.Println(sort.StringsAreSorted(slice3))
	sort.Strings(slice3)
	fmt.Println(slice3)

}
