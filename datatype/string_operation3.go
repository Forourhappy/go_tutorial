package datatype

import (
	"fmt"
	"strings"
)

func StringOperation3() {
	// 결합: 일반연산
	str1 := "Hello" + "Golang" + "World" + "!!!!"
	str2 := "erme"

	// 새로운 string을 생성하여 결합
	fmt.Println(str1 + str2)
	// 결합: Join
	strSet := []string{}
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)
	// 새로운 string을 생성하지 않아서 효율적
	fmt.Println(strings.Join(strSet, "----"))
}
