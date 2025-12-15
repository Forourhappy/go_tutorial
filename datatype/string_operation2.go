package datatype

import "fmt"

func StringOperation2() {
	// 비교
	str1 := "Golang"
	str2 := "World"

	fmt.Println(str1 == str2)
	fmt.Println(str1 != str2)
	fmt.Println(str1 > str2)
	fmt.Println(str1 < str2) // Go 문자열 -> 아스키 코드에 대한 사전식 비교
}
