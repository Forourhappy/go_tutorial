package condition

import "fmt"

func Switch3() {
	a := 30 / 15

	switch a {
	case 2, 4, 6:
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5:
		fmt.Println("a -> ", a, "는 홀수")
	}

	switch e := "go"; e {
	case "java":
		fmt.Println("Java!")
		fallthrough
	case "go":
		fmt.Println("Go!")
		fallthrough
	case "python":
		fmt.Println("Python!")
		fallthrough
	case "ruby":
		fmt.Println("Ruby!")
		fallthrough
	case "c":
		fmt.Println("C!")
	}
}
