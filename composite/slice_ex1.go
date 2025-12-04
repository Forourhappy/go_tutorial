package composite

import "fmt"

func SliceEx1() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	// 용량을 벗어나면 내부적으로 용량을 늘리는 과정을 거침 -> 성능에 악영향
	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...)      // 슬라이스를 삽입할 경우 ... 사용
	s3 = append(s2, s3[0:3]...) // 추출 후 병합

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("len: %d, cap: %d, value: %v\n", len(s4), cap(s4), s4)
	}
}
