package datatype

import (
	"fmt"
)

func StringOperation1() {
	/*
		문자열 연산
		추출, 비교, 조합(결합)
	*/

	var str1 string = "Golang"
	var str2 string = "Wrold"
	fmt.Println(str1[0:2], str1[0])
	fmt.Println(str2[3:], str2[0])
	fmt.Println(str2[:4])
	fmt.Println(str2[1:3])
}
