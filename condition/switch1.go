package condition

import "fmt"

func Switch1() {
	/**
	제어문(조건문): switch
	switch 뒤 표현식 생략 가능
	case 뒤 표현식 사용 가능
	자동 break 때문에 fallthrough 존재
	Type 분기 -> 값이 아닌 변수 Type으로 분기 가능
	*/

	a := -7

	switch {
	case a < 0:
		fmt.Println("음수")
	case a == 0:
		fmt.Println("0")
	default:
		fmt.Println("양수")
	}

	switch b := 20; {
	case b < 0:
		fmt.Println("음수")
	case b == 0:
		fmt.Println("0")
	}

	switch c := "go"; c {
	case "go":
		fmt.Println("go")
	case "golang":
		fmt.Println("golang")
	default:
		fmt.Println("default")
	}

	switch c := "go"; c + "lang" {
	case "golang":
	case "java":
	default:
		fmt.Println(c)
	}

	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i < j")
	case i == j:
		fmt.Println("i == j")
	default:
		fmt.Println("i > j")
	}
}
